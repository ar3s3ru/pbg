package pbg

import (
	"log"
	"strings"
	"github.com/valyala/fasthttp"
)

type (
	// Opzione funzionale legata a oggetti di tipo *server (non esportato)
	serverOption func(*server) error
	// Opzione funzionale per modificare le proprietà degli oggetti Server
	// durante la costruzione nell'apposito Factory method
	ServerOption func(Server) error
)

// Adatta opzioni funzionali del tipo server non esportato ad opzioni funzionali
// dell'interfaccia Server esportata
func adaptServerOption(option serverOption) ServerOption {
	return func(srv Server) error {
		s, ok := srv.(*server)
		if !ok {
			return ErrInvalidServerType
		}

		return option(s)
	}
}

// Valida il valore della porta HTTP passata come argomento
// Torna true nel caso in cui il valore di port denota una porta HTTP valida,
// false altrimenti
func validPortInput(port int) bool {
	return port >= 1 && port <= 65536
}

// Valida il valore dell'API endpoint
// Ricordiamo che un API endpoint deve essere lungo almeno 2 caratteri
// e deve iniziare con "/"
//
// es.
//     "/a"   -> OK
//     "/API" -> OK
//     ""     -> NON OK
//     "/"    -> NON OK
//
func validAPIEndpoint(endpoint string) bool {
	return len(endpoint) >= 2 && strings.HasPrefix(endpoint, "/")
}

// Specifica la porta HTTP che il Server da costruire deve utilizzare
func WithHTTPPort(port int) ServerOption {
	return adaptServerOption(func(srv *server) error {
		if !validPortInput(port) {
			return ErrInvalidHTTPPort
		}

		srv.port = port
		return nil
	})
}

// Specifica l'API Endpoint che il Server da costruire deve utilizzare
func WithAPIEndpont(endpoint string) ServerOption {
	return adaptServerOption(func(srv *server) error {
		if !validAPIEndpoint(endpoint) {
			return ErrInvalidAPIEndpoint
		}

		srv.apiEndpoint = endpoint
		return nil
	})
}

// Specifica l'APIResponser che il Server da costruire deve utilizzare
func WithAPIResponser(responser APIResponser) ServerOption {
	return adaptServerOption(func(srv *server) error {
		if responser == nil {
			return ErrInvalidAPIResponser
		}

		srv.apiResponser = responser
		return nil
	})
}

// Specifica che il Server da costruire deve usare il Logger passato
// come argomento
func WithLogger(logger *log.Logger) ServerOption {
	return adaptServerOption(func(srv *server) error {
		if logger == nil {
			return ErrInvalidLogger
		}

		srv.logger = logger
		return nil
	})
}

// Specifica il MoveDBComponent che il Server deve usare
func WithMoveDBComponent(mdb MoveDBComponent) ServerOption {
	return adaptServerOption(func(srv *server) error {
		if mdb == nil {
			return ErrInvalidMoveDBComponent
		}

		srv.moveDB = mdb
		return nil
	})
}

// Specifica il MoveDBComponent che il Server deve usare
func WithPokèmonDBComponent(pdb PokèmonDBComponent) ServerOption {
	return adaptServerOption(func(srv *server) error {
		if pdb == nil {
			return ErrInvalidMoveDBComponent
		}

		srv.pokèmonDB = pdb
		return nil
	})
}

// Specifica il MoveDBComponent che il Server deve usare
func WithTrainerDBComponent(tdb TrainerDBComponent) ServerOption {
	return adaptServerOption(func(srv *server) error {
		if tdb == nil {
			return ErrInvalidMoveDBComponent
		}

		srv.trainerDB = tdb
		return nil
	})
}

// Specifica il SessionComponent che il Server deve usare
func WithSessionDBComponent(sc SessionComponent) ServerOption {
	return adaptServerOption(func(srv *server) error {
		if sc == nil {
			return ErrInvalidSessionComponent
		}

		srv.sessionDB = sc
		return nil
	})
}

func WithFastHTTPServer(fsrv *fasthttp.Server) ServerOption {
	return adaptServerOption(func(srv *server) error {
		if fsrv == nil {
			return ErrInvalidFastHTTPServer
		}

		srv.Server = fsrv
		return nil
	})
}
