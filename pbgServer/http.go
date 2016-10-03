package pbgServer

import (
    "github.com/valyala/fasthttp"
    "github.com/buaazp/fasthttprouter"
    "log"
    "bytes"
    "encoding/base64"
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

    // Funzione callback per la gestione di richieste HTTP autenticate.
    //
    // Callback function for authenticated HTTP requests management.
    AHandler func(Session, IServerContext, *fasthttp.RequestCtx, fasthttprouter.Params)
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

// Registra un nuovo handler per un determinato percorso e con un determinato metodo HTTP.
//
// Registers a new handler for a specified path, with a specified HTTP method.
func (srv *pbgServer) Handle(method HTTPMethod, path string, handler Handler) PBGServer {
    if (handler != nil) {
        srv.httpRouter.Handle(string(method), path, func(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
            // New handle server-specific
            handler(srv, ctx, ps)
            // Logging HTTP request details
            log.Printf("[%s - %d] (%s) %s\n", ctx.Method(), ctx.Response.StatusCode(), ctx.RemoteIP(), ctx.RequestURI())
        })
    }
    // Method chaining
    return srv
}

func (srv *pbgServer) AuthHandle(method HTTPMethod, path string, handler AHandler) PBGServer {
    if handler != nil {
        srv.Handle(method, path, func(sctx IServerContext, ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
            statMsg := fasthttp.StatusMessage(fasthttp.StatusUnauthorized)

            if bToken := basicAuth(ctx); bToken != nil {
                if sess, err := sctx.GetSessMechanism().GetSession(string(bToken));
                    err == ErrSessionExpired || err == ErrSessionNotFound {
                    // Token not valid anymore
                    sctx.GetSessMechanism().RemoveSession(string(bToken))
                    statMsg = err.Error()
                } else if err == nil {
                    handler(sess, sctx, ctx, ps)
                    return
                }
            }

            // Request Basic Authentication otherwise
            ctx.Response.Header.Set("WWW-Authenticate", "Basic realm=Restricted")
            ctx.Error(statMsg, fasthttp.StatusUnauthorized)
        })
    }

    return srv
}