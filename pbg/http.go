package pbg

import "github.com/valyala/fasthttp"

type HTTPMethod string

const (
    GET    HTTPMethod = "GET"
    POST   HTTPMethod = "POST"
    PUT    HTTPMethod = "PUT"
    DELETE HTTPMethod = "DELETE"
    HEAD   HTTPMethod = "HEAD"

    authorizationHeader = "Authorization"
    authenticateHeader  = "WWW-Authenticate"
    authenticateValue   = "Basic realm=Restricted"
)

func adaptCheckAuth(am AuthorizationMechanism) Adapter {
    return func(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
        return func(ctx *fasthttp.RequestCtx) {
            statusCode, session, err := am.CheckAuthorization(ctx.Request.Header.Peek(authorizationHeader))
            if err != nil {
                ctx.Error(err.Error(), statusCode)
                ctx.Response.Header.Set(authenticateHeader, authenticateValue)
                return
            }

            ctx.SetUserValue(SessionKey, session)
            handler(ctx)
        }
    }
}

func adaptCheckAuthAPI(am AuthorizationMechanism) Adapter {
    return func(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
        return func(ctx *fasthttp.RequestCtx) {
            statusCode, session, err := am.CheckAuthorization(ctx.Request.Header.Peek(authorizationHeader))
            if err != nil {
                ctx.SetUserValue(APIErrorKey, err)
                ctx.SetStatusCode(statusCode)
                ctx.Response.Header.Set(authenticateHeader, authenticateValue)
                return
            }

            ctx.SetUserValue(SessionKey, session)
            handler(ctx)
        }
    }
}

func (srv *server) Handle(method HTTPMethod, path string, handler fasthttp.RequestHandler) {
    srv.router.Handle(string(method), path, handler)
}

func (srv *server) AuthHandle(method HTTPMethod, path string, handler fasthttp.RequestHandler) {
    srv.Handle(method, path,
        Adapt(handler, adaptCheckAuth(srv.authorizationMechanism)),
    )
}

func (srv *server) APIHandle(method HTTPMethod, path string, handler fasthttp.RequestHandler) {
    // TODO: string optimization with slices of []byte
    srv.Handle(method, srv.config.APIEndpoint() + path,
        Adapt(handler, srv.apiWriter),
    )
}

func (srv *server) APIAuthHandle(method HTTPMethod, path string, handler fasthttp.RequestHandler) {
    srv.APIHandle(method, path,
        Adapt(handler, adaptCheckAuthAPI(srv.authorizationMechanism)),
    )
}
