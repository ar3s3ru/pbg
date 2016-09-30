package pbgServer

import (
    "github.com/valyala/fasthttp"
    "github.com/buaazp/fasthttprouter"
)

type (
    // Rappresenta i possibili metodi HTTP utilizzabili nella funzione Handle().
    //
    // Represents HTTP methods that can be used into the Handle() function.
    HTTPMethod string

    // Funzione callback che viene chiamata dal router del server ogni qualvolta che arriva
    // una richiesta HTTP da gestire in un particolare percorso.
    //
    // Callback function called by server's router every time an HTTP request arrives and needs to be handled.
    Handler func(IServerContext, *fasthttp.RequestCtx, fasthttprouter.Params)
)

const (
    GET     HTTPMethod = "GET"
    POST    HTTPMethod = "POST"
    PUT     HTTPMethod = "PUT"
    OPTIONS HTTPMethod = "OPTIONS"
    DELETE  HTTPMethod = "DELETE"
)

// Registra un nuovo handler per un determinato percorso e con un determinato metodo HTTP.
//
// Registers a new handler for a specified path, with a specified HTTP method.
func (srv *pbgServer) Handle(method HTTPMethod, path string, handler Handler) PBGServer {
    if (handler != nil) {
        srv.httpRouter.Handle(string(method), path, func(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
            // New handle server-specific
            handler(srv, ctx, ps)
        })
    }
    // Method chaining
    return srv
}