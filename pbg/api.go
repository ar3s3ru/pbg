package pbg

import "github.com/valyala/fasthttp"

type (
    ApiResponser interface {
        // Adapter
        Writer(fasthttp.RequestHandler) fasthttp.RequestHandler
        // Utilities
        WriteSuccess(*fasthttp.RequestCtx)
        WriteError(*fasthttp.RequestCtx)
    }
)
