package pbg

import "github.com/valyala/fasthttp"

type HTTPMethod string

const (
    GET    HTTPMethod = "GET"
    POST   HTTPMethod = "POST"
    PUT    HTTPMethod = "PUT"
    DELETE HTTPMethod = "DELETE"
    HEAD   HTTPMethod = "HEAD"
)

func (srv server) Handle(method HTTPMethod, path string, handler fasthttp.RequestHandler) {
    srv.router.Handle(string(method), path, handler)
}

func (srv server) AuthHandle(method HTTPMethod, path string, handler fasthttp.RequestHandler) {
    srv.Handle(string(method), path,
        Adapt(handler, srv.authorizationMechanism.CheckAuthorization),
    )
}

func (srv server) ApiHandle(method HTTPMethod, path string, handler fasthttp.RequestHandler) {
    // TODO: string optimization with slices of []byte
    srv.Handle(method, srv.config.ApiEndpoint() + path,
        Adapt(handler, srv.apiWriter),
    )
}

func (srv server) ApiAuthHandle(method HTTPMethod, path string, handler fasthttp.RequestHandler) {
    srv.ApiHandle(method, path,
        Adapt(handler, srv.authorizationMechanism.CheckAuthorization),
    )
}
