package pbg

import (
    "fmt"
    "log"

    "github.com/buaazp/fasthttprouter"
    "github.com/valyala/fasthttp"
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
        GET(path string,    handle fasthttp.RequestHandler)
        PUT(path string,    handle fasthttp.RequestHandler)
        POST(path string,   handle fasthttp.RequestHandler)
        DELETE(path string, handle fasthttp.RequestHandler)

        API_GET(path string,    handle fasthttp.RequestHandler)
        API_PUT(path string,    handle fasthttp.RequestHandler)
        API_POST(path string,   handle fasthttp.RequestHandler)
        API_DELETE(path string, handle fasthttp.RequestHandler)
    }

    // Rappresenta tutti gli Adapter messi a disposizione dal Server
    // Forniscono accesso a particolari entità protette del software mediante
    // RequestCtx
    //
    // Per richiederli, usare nei RequestHandler:
    //     ```
    //         func(ctx *fasthttp.RequestCtx) {
    //             // Type assertion qui, ctx.UserValue() restituisce interface{}
    //             ent, ok := ctx.UserValue(ENTITY_KEY).(Entity)
    //             if !ok {
    //                 // Errore!
    //                 ...
    //             }
    //
    //             // Usa il valore ent
    //             ...
    //         }
    //     ```
    ServerAdapters interface {
        // Fornisce l'accesso al logger mediante RequestCtx
        // Il logger è disponibile con la chiave pbg.LoggerKey
        WithLogger(fasthttp.RequestHandler)        fasthttp.RequestHandler
        // Fornisce un DataInterface mediante RequestCtx
        // L'interfaccia è disponibile con la chiave pbg.DataInterfaceKey
        WithDataAccess(fasthttp.RequestHandler)    fasthttp.RequestHandler
        // Fornisce un SessionInterface mediante RequestCtx
        // L'interfaccia è disponibile con la chiave pbg.SessionInterfaceKey
        WithSessionAccess(fasthttp.RequestHandler) fasthttp.RequestHandler
    }

    // Un Server del framework pbg ha la possibilità di utilizzare un Logger,
    // le capacità di usare il protocollo HTTP (quindi registrare RequestHandlers)
    // e degli Adapter per estendere a dovere i RequestHandler utente con l'accesso a risorse protette
    Server interface {
        Logger
        ServerHTTP
        ServerAdapters

        // Avvia il server
        Start()
    }

    server struct {
        port   int                     // Porta HTTP
        logger *log.Logger             // Logger
        router *fasthttprouter.Router  // HTTP router

        apiEndpoint      string            // Endpoint dell'API server
        apiResponser     APIResponser      // Traduttore per le richieste API
        dataComponent    DataComponent     // Riferimento al componente software del Models DB
        sessionComponent SessionComponent  // Riferimento al componente software del Sessions DB
    }
)

var (
    // Check functions (to show how many ways there are to declare fuctions in Go)
    check = func(data interface{}, err error) {
        if data == nil {
            panic(err)
        }
    }

    checkErr = func(err error) {
        if err != nil {
            panic(err)
        }
    }
)

// Factory methods per oggetti pbg.Server
// Usare i Server functional options per utilizzare determinate proprietà
// sull'oggetto costruito
func NewServer(options ...ServerOption) Server {
    srv := &server{
        port:         8080,
        router:       fasthttprouter.New(),
        apiResponser: NewJSONResponser(),
        apiEndpoint:  "/api",
    }

    for _, option := range options {
        // TODO: cleanup and exiting
        checkErr(option(srv))
    }

    // Check needed parameters
    check(srv.apiResponser,     ErrInvalidAPIResponser)
    check(srv.dataComponent,    ErrInvalidDataComponent)
    check(srv.sessionComponent, ErrInvalidSessionComponent)

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
    srv.router.GET(srv.apiEndpoint + path, Adapt(handle, srv.apiWriter))
}

func (srv *server) API_PUT(path string, handle fasthttp.RequestHandler) {
    srv.router.PUT(srv.apiEndpoint + path, Adapt(handle, srv.apiWriter))
}

func (srv *server) API_POST(path string, handle fasthttp.RequestHandler) {
    srv.router.POST(srv.apiEndpoint + path, Adapt(handle, srv.apiWriter))
}

func (srv *server) API_DELETE(path string, handle fasthttp.RequestHandler) {
    srv.router.DELETE(srv.apiEndpoint + path, Adapt(handle, srv.apiWriter))
}

func (srv *server) Start() {
    address := fmt.Sprintf(":%d", srv.port)

    srv.Log("Serving on", address)
    srv.Log(
        fasthttp.ListenAndServe(address, srv.router.Handler),
    )
}

func (srv *server) Log(v ...interface{}) {
    if srv.logger != nil {
        // Stampa sul logger solo se è diverso da nil
        // TODO: in effetti è sempre diverso da nil, forse si può eliminare il check
        srv.logger.Println(v...)
    }
}

/**************
     Server
    Adapters
 **************/

func (srv *server) WithLogger(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
    return func(ctx *fasthttp.RequestCtx) {
        ctx.SetUserValue(LoggerKey, srv.logger)
        handler(ctx)
    }
}

func (srv *server) WithDataAccess(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
    return func(ctx *fasthttp.RequestCtx) {
        dataInterface := srv.dataComponent.Supply()
        defer srv.dataComponent.Retrieve(dataInterface)

        ctx.SetUserValue(DataInterfaceKey, dataInterface)
        handler(ctx)
    }
}

func (srv *server) WithSessionAccess(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
    return func(ctx *fasthttp.RequestCtx) {
        sessionInterface := srv.sessionComponent.Supply()
        defer srv.sessionComponent.Retrieve(sessionInterface)

        ctx.SetUserValue(SessionInterfaceKey, sessionInterface)
        handler(ctx)
    }
}
