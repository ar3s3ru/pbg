package pbg

import "github.com/valyala/fasthttp"

type (
    Adapter func(fasthttp.RequestHandler) fasthttp.RequestHandler
    ServerAdapters interface {
        WithDataAccess(fasthttp.RequestHandler)    fasthttp.RequestHandler
        WithSessionAccess(fasthttp.RequestHandler) fasthttp.RequestHandler
    }
)

func Adapt(rh fasthttp.RequestHandler, adapters ...Adapter) fasthttp.RequestHandler {
    for _, adapter := range adapters {
        rh = adapter(rh)
    }

    return rh
}

func (srv *server) WithDataAccess(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
    return func(ctx *fasthttp.RequestCtx) {
        ctx.SetUserValue(DataAccessKey, srv.dataMechanism)
        handler(ctx)
    }
}

func (srv *server) WithSessionAccess(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
    return func(ctx *fasthttp.RequestCtx) {
        ctx.SetUserValue(SessionAccessKey, srv.sessionMechanism)
        handler(ctx)
    }
}
