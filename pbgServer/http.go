package pbgServer

import (
    "github.com/valyala/fasthttp"
    "github.com/buaazp/fasthttprouter"
    "bytes"
    "encoding/base64"
    "errors"
)

type (
    // Interfaccia che specifica il comportamento di un server HTTP
    IHTTPServer interface {
        GetAPIResponse() APIResponse

        Handle(HTTPMethod, string, fasthttprouter.Handle) PBGServer
        ServHandle(HTTPMethod, string, Handler)           PBGServer
        APIHandle(HTTPMethod, string, APIHandler)         PBGServer
        APIAuthHandle(HTTPMethod, string, APIAuthHandler) PBGServer
    }

    // Rappresenta i possibili metodi HTTP utilizzabili nella funzione Handle().
    //
    // Represents HTTP methods that can be used into the Handle() function.
    HTTPMethod string

    // Callback che si occupa di fare marshalling delle risposte alle richieste HTTP effettuate
    // all'API REST del server.
    APIResponse func (int, interface{}, error, *fasthttp.RequestCtx)

    // Handle per richieste HTTP che riguardano il contesto del server, ma che non appartengono all'API REST
    // del server stesso.
    Handler func (IServerContext, *fasthttp.RequestCtx, fasthttprouter.Params)

    // Funzione callback che viene chiamata dal router del server ogni qualvolta che arriva
    // una richiesta HTTP da gestire in un particolare percorso.
    //
    // Callback function called by server's router every time an HTTP request arrives and needs to be handled.
    APIHandler func (IServerContext, *fasthttp.RequestCtx, fasthttprouter.Params) (int, interface{}, error)

    // Funzione callback per la gestione di richieste HTTP autenticate.
    //
    // Callback function for authenticated HTTP requests management.
    APIAuthHandler func (Session, IServerContext, *fasthttp.RequestCtx, fasthttprouter.Params) (int, interface{}, error)
)

const (
    GET     HTTPMethod = "GET"
    POST    HTTPMethod = "POST"
    PUT     HTTPMethod = "PUT"
    OPTIONS HTTPMethod = "OPTIONS"
    DELETE  HTTPMethod = "DELETE"
)

var (
    basicAuthPrefix = []byte("Basic ")

    authComma = []byte(":")
    authPass  = []byte("x")

    // Errors
    ErrInvalidAPIHandler     = errors.New("Invalid API handler used, <nil> value")
    ErrInvalidAPIAuthHandler = errors.New("Invalid API authHandler used, <nil> value")
    ErrUnauthorized          = errors.New(fasthttp.StatusMessage(fasthttp.StatusUnauthorized))
)

func basicAuth(ctx *fasthttp.RequestCtx) []byte {
    // Get the Basic Authentication credentials
    auth := ctx.Request.Header.Peek("Authorization")
    if bytes.HasPrefix(auth, basicAuthPrefix) {
        // Check credentials
        payload, err := base64.StdEncoding.DecodeString(string(auth[len(basicAuthPrefix):]))
        if err == nil {
            pair := bytes.SplitN(payload, authComma, 2)
            if len(pair) == 2 && bytes.Equal(pair[1], authPass) {
                return pair[0]
            }
        }
    }

    // Something went wrong
    return nil
}

func (srv *pbgServer) GetAPIResponse() APIResponse {
    return srv.apiResponse
}

// Registra un nuovo handler per un determinato percorso e con un determinato metodo HTTP.
//
// Registers a new handler for a specified path, with a specified HTTP method.
func (srv *pbgServer) Handle(method HTTPMethod, path string, handler fasthttprouter.Handle) PBGServer {
    // Handle
    srv.httpRouter.Handle(string(method), path, handler)
    // Method chaining
    return srv
}

func (srv *pbgServer) ServHandle(method HTTPMethod, path string, handler Handler) PBGServer {
    if handler == nil {
        panic(ErrInvalidAPIHandler)
    }

    srv.httpRouter.Handle(string(method), path, func (ctx *fasthttp.RequestCtx, pm fasthttprouter.Params) {
        handler(srv, ctx, pm)
    })

    return srv
}

func (srv *pbgServer) APIHandle(method HTTPMethod, path string, handler APIHandler) PBGServer {
    if handler == nil {
        panic(ErrInvalidAPIHandler)
    }

    srv.httpRouter.Handle(string(method), path, func (ctx *fasthttp.RequestCtx, pm fasthttprouter.Params) {
        code, data, err := handler(srv, ctx, pm)
        srv.GetAPIResponse()(code, data, err, ctx)
    })
    // Method chaining
    return srv
}

func (srv *pbgServer) APIAuthHandle(method HTTPMethod, path string, handler APIAuthHandler) PBGServer {
    if handler == nil {
        panic(ErrInvalidAPIAuthHandler)
    }

    srv.APIHandle(method, path,
        func (sctx IServerContext, ctx *fasthttp.RequestCtx, pm fasthttprouter.Params) (int, interface{}, error) {
            err := ErrUnauthorized

            if bToken := basicAuth(ctx); bToken != nil {
                if sess, err := sctx.GetSessMechanism().GetSession(string(bToken));
                   err == ErrSessionExpired || err == ErrSessionNotFound {
                    // Token not valid anymore
                    sctx.GetSessMechanism().RemoveSession(string(bToken))
                } else if err == nil {
                    return handler(sess, sctx, ctx, pm)
                }
            }

            // Request Basic Authentication otherwise
            ctx.Response.Header.Set("WWW-Authenticate", "Basic realm=Restricted")
            return fasthttp.StatusUnauthorized, nil, err
        },
    )

    return srv
}