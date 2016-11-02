package pbg

import "github.com/valyala/fasthttp"

type Adapter func(fasthttp.RequestHandler) fasthttp.RequestHandler

func Adapt(rh fasthttp.RequestHandler, adapters ...Adapter) fasthttp.RequestHandler {
    for _, adapter := range adapters {
        rh = adapter(rh)
    }

    return rh
}
