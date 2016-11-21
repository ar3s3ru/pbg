package pbg

import (
	"fmt"
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"os"
)

type (
	// Rappresenta un server HTTP del framework pbg
	// Include tutti i metodi per registrare RequestHandlers e
	// creare un REST API
	//
	// GET    -> Leggi risorse
	// PUT    -> Aggiorna risorse
	// POST   -> Crea risorse
	// DELETE -> Cancella risorse
	ServerHTTP interface {
		GET(path string, handle fasthttp.RequestHandler)
		PUT(path string, handle fasthttp.RequestHandler)
		POST(path string, handle fasthttp.RequestHandler)
		DELETE(path string, handle fasthttp.RequestHandler)

		API_GET(path string, handle fasthttp.RequestHandler)
		API_PUT(path string, handle fasthttp.RequestHandler)
		API_POST(path string, handle fasthttp.RequestHandler)
		API_DELETE(path string, handle fasthttp.RequestHandler)
	}

	// Rappresenta tutti gli Adapter messi a disposizione dal Server
	// Forniscono accesso a particolari entità protette del software mediante
	// RequestCtx
	//
	// Per richiederli, usare nei RequestHandler:
	//
	//     func(ctx *fasthttp.RequestCtx) {
	//         // Type assertion qui, ctx.UserValue() restituisce interface{}
	//         ent, ok := ctx.UserValue(ENTITY_KEY).(Entity)
	//         if !ok {
	//             // Errore!
	//             ...
	//         }
	//
	//         // Usa il valore ent
	//         ...
	//     }
	//
	ServerAdapters interface {
		// Fornisce interfaccia al DB delle mosse mediante RequestCtx
		// L'interfaccia è disponibile con la chiave pbg.MoveDBInterfaceKey
		WithMoveDBAccess(fasthttp.RequestHandler) fasthttp.RequestHandler
		// Fornisce interfaccia al DB dei Pokèmon mediante RequestCtx
		// L'interfaccia è disponibile con la chiave pbg.PokèmonDBInterfaceKey
		WithPokèmonDBAccess(fasthttp.RequestHandler) fasthttp.RequestHandler
		// Fornisce interfaccia al DB degli allenatori mediante RequestCtx
		// L'interfaccia è disponibile con la chiave pbg.TrainerDBInterfaceKey
		WithTrainerDBAccess(fasthttp.RequestHandler) fasthttp.RequestHandler
		// Fornisce un SessionInterface mediante RequestCtx
		// L'interfaccia è disponibile con la chiave pbg.SessionDBInterfaceKey
		WithSessionDBAccess(fasthttp.RequestHandler) fasthttp.RequestHandler
	}

	// Un Server del framework pbg ha la possibilità di utilizzare un Logger,
	// le capacità di usare il protocollo HTTP (quindi registrare RequestHandlers)
	// e degli Adapter per estendere a dovere i RequestHandler utente con l'accesso a risorse protette
	Server interface {
		fasthttp.Logger
		ServerHTTP
		ServerAdapters

		// Avvia il server
		Start()
	}

	server struct {
		*fasthttp.Server

		port   int
		router *fasthttprouter.Router
		logger *log.Logger

		apiEndpoint  string       // Endpoint dell'API server
		apiResponser APIResponser // Traduttore per le richieste API

		pokèmonDB PokèmonDBComponent // Riferimento al componente software del Models DB
		moveDB    MoveDBComponent    //
		trainerDB TrainerDBComponent //
		sessionDB SessionComponent   // Riferimento al componente software del Sessions DB
	}
)

var (
	defaultLogger = log.New(os.Stderr, "PBG Server: ", log.LstdFlags|log.Lshortfile)

	checkValueWithLogger = func(logger *log.Logger, value interface{}, err error) {
		if logger != nil && value == nil {
			logger.Fatal(err)
		}
	}

	checkErrorWithLogger = func(logger *log.Logger, err error) {
		if logger != nil && err != nil {
			logger.Fatal(err)
		}
	}

	checkValue = func(value interface{}, err error) { checkValueWithLogger(defaultLogger, value, err) }
	checkError = func(err error) { checkErrorWithLogger(defaultLogger, err) }
)

// Factory methods per oggetti pbg.Server
// Usare i Server functional options per utilizzare determinate proprietà
// sull'oggetto costruito
func NewServer(options ...ServerOption) Server {
	srv := &server{
		port:         80,
		router:       fasthttprouter.New(),
		apiResponser: NewJSONResponser(),
		apiEndpoint:  "/api",
	}

	for _, option := range options {
		checkError(option(srv))
	}

	if srv.logger == nil {
		srv.logger = defaultLogger
	}

	// Check needed parameters
	checkValue(srv.apiResponser, ErrInvalidAPIResponser)
	checkValue(srv.moveDB, ErrInvalidMoveDBComponent)
	checkValue(srv.pokèmonDB, ErrInvalidPokèmonDBComponent)
	checkValue(srv.sessionDB, ErrInvalidSessionComponent)

	// If server is not set up, use a standard server
	if srv.Server == nil {
		srv.Server = &fasthttp.Server{
			Logger: srv.logger,
		}
	}

	// Server handler from fasthttprouter
	srv.Server.Handler = srv.router.Handler

	return srv
}

func (srv *server) GET(path string, handle fasthttp.RequestHandler) {
	srv.router.GET(path, handle)
}

func (srv *server) PUT(path string, handle fasthttp.RequestHandler) {
	srv.router.PUT(path, handle)
}

func (srv *server) POST(path string, handle fasthttp.RequestHandler) {
	srv.router.POST(path, handle)
}

func (srv *server) DELETE(path string, handle fasthttp.RequestHandler) {
	srv.router.DELETE(path, handle)
}

func (srv *server) API_GET(path string, handle fasthttp.RequestHandler) {
	srv.router.GET(srv.apiEndpoint+path, Adapt(handle, srv.apiWriter))
}

func (srv *server) API_PUT(path string, handle fasthttp.RequestHandler) {
	srv.router.PUT(srv.apiEndpoint+path, Adapt(handle, srv.apiWriter))
}

func (srv *server) API_POST(path string, handle fasthttp.RequestHandler) {
	srv.router.POST(srv.apiEndpoint+path, Adapt(handle, srv.apiWriter))
}

func (srv *server) API_DELETE(path string, handle fasthttp.RequestHandler) {
	srv.router.DELETE(srv.apiEndpoint+path, Adapt(handle, srv.apiWriter))
}

func (srv *server) Start() {
	address := fmt.Sprintf(":%d", srv.port)
	srv.logger.Fatal(
		srv.ListenAndServe(address),
	)
}

func (srv *server) Printf(format string, v ...interface{}) {
	srv.logger.Printf(format, v...)
}

/**************
    Server
   Adapters
**************/

func (srv *server) WithMoveDBAccess(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		dataInterface := srv.moveDB.Supply()
		defer srv.moveDB.Retrieve(dataInterface)

		ctx.SetUserValue(MoveDBInterfaceKey, dataInterface)
		handler(ctx)
	}
}

func (srv *server) WithPokèmonDBAccess(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		dataInterface := srv.pokèmonDB.Supply()
		defer srv.pokèmonDB.Retrieve(dataInterface)

		ctx.SetUserValue(PokèmonDBInterfaceKey, dataInterface)
		handler(ctx)
	}
}

func (srv *server) WithTrainerDBAccess(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		dataInterface := srv.trainerDB.Supply()
		defer srv.trainerDB.Retrieve(dataInterface)

		ctx.SetUserValue(TrainerDBInterfaceKey, dataInterface)
		handler(ctx)
	}
}

func (srv *server) WithSessionDBAccess(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		sessionInterface := srv.sessionDB.Supply()
		defer srv.sessionDB.Retrieve(sessionInterface)

		ctx.SetUserValue(SessionDBInterfaceKey, sessionInterface)
		handler(ctx)
	}
}
