package pbgServer

import (
    "github.com/valyala/fasthttp"
    "github.com/buaazp/fasthttprouter"
)

type (
    IServerContext interface {
        GetDataMechanism() IDataMechanism
        GetAuthMechanism() IAuthMechanism
        GetSessMechanism() ISessionsMechanism

        Handle(HTTPMethod, string, Handler)
    }

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

func (srv *pbgServer) GetDataMechanism() IDataMechanism {
    return srv.dataMechanism
}

func (srv *pbgServer) GetAuthMechanism() IAuthMechanism {
    return srv.authMechanism
}

func (srv *pbgServer) GetSessMechanism() ISessionsMechanism {
    return srv.sessMechanism
}

func (srv *pbgServer) Handle(method HTTPMethod, path string, handler Handler) {
    if (handler != nil) {
        srv.httpRouter.Handle(string(method), path, func(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
            // New handle server-specific
            handler(srv, ctx, ps)
        })
    }
}