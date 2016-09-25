package pbgServer

import (
    "github.com/valyala/fasthttp"
    "github.com/buaazp/fasthttprouter"
)

type (
    HTTPMethod string
    Handler func(IServerContext, *fasthttp.RequestCtx, fasthttprouter.Params)
)

const (
    GET    HTTPMethod = "GET"
    POST   HTTPMethod = "POST"
    PUT    HTTPMethod = "PUT"
    OPTION HTTPMethod = "OPTION"
    DELETE HTTPMethod = "DELETE"
)

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