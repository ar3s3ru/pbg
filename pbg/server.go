package pbg

import (
    "fmt"

    "github.com/buaazp/fasthttprouter"
    "github.com/valyala/fasthttp"
)

type (
    Server interface {
        Start()

        Handle(method, path string, handler fasthttp.RequestHandler)
        AuthHandle(method, path string, handler fasthttp.RequestHandler)

        ApiHandle(method, path string, handler fasthttp.RequestHandler)
        ApiAuthHandle(method, path string, handler fasthttp.RequestHandler)
    }

    ServerBuilder interface {
        WithConfiguration(Configuration)          ServerBuilder
        WithDataMechanism(DataMechanism)          ServerBuilder
        WithSessionMechanism(DataMechanism)       ServerBuilder
        WithAuthorizationMechanism(DataMechanism) ServerBuilder

        WithApiResponser(ApiResponser) ServerBuilder
        //WithLogger()                   ServerBuilder

        Build() Server
    }

    server struct {
        apiResponser ApiResponser

        dataMechanism          DataMechanism
        sessionMechanism       SessionMechanism
        authorizationMechanism AuthorizationMechanism

        config Configuration
        router *fasthttprouter.Router
    }

    serverBuilder func(
        Configuration, DataMechanism, SessionMechanism, AuthorizationMechanism, ApiResponser,
    ) server
)

var (
    baseConfig = &BaseConfiguration{Port: 8080, Local: true}
)

// Builder
func NewServerBuilder() ServerBuilder {
    return func(cfg Configuration,
                dm DataMechanism, sm SessionMechanism, am AuthorizationMechanism, ar ApiResponser) server {
        return &server{
            config: cfg,
            router: fasthttprouter.New(),
            // Module mechanisms
            dataMechanism: dm, sessionMechanism: sm, authorizationMechanism: am,
            apiResponser:  ar,
        }
    }
}

func (sb serverBuilder) WithConfiguration(cfg Configuration) ServerBuilder {
    return func(_ Configuration,
                dm DataMechanism, sm SessionMechanism, am AuthorizationMechanism, ar ApiResponser) server {
        return sb(cfg, dm, sm, am, ar)
    }
}

func (sb serverBuilder) WithDataMechanism(dm DataMechanism) ServerBuilder {
    return func(cfg Configuration,
                _ DataMechanism, sm SessionMechanism, am AuthorizationMechanism, ar ApiResponser) server {
        return sb(cfg, dm, sm, am, ar)
    }
}

func (sb serverBuilder) WithSessionMechanism(sm SessionMechanism) ServerBuilder {
    return func(cfg Configuration,
                dm DataMechanism, _ SessionMechanism, am AuthorizationMechanism, ar ApiResponser) server {
        return sb(cfg, dm, sm, am, ar)
    }
}

func (sb serverBuilder) WithAuthorizationMechanism(am AuthorizationMechanism) ServerBuilder {
    return func(cfg Configuration,
                dm DataMechanism, sm SessionMechanism, _ AuthorizationMechanism, ar ApiResponser) server {
        return sb(cfg, dm, sm, am, ar)
    }
}

func (sb serverBuilder) WithApiResponser(ar ApiResponser) {
    return func(cfg Configuration,
                dm DataMechanism, sm SessionMechanism, am AuthorizationMechanism, _ ApiResponser) server {
        return sb(cfg, dm, sm, am, ar)
    }
}

func (sb serverBuilder) Build() Server {
    return sb(baseConfig, nil, nil, nil, nil)
}

// Server
func (srv server) Start() {
    var local string
    if srv.config.LocalHost() {
        local = "localhost"
    }

    fasthttp.ListenAndServe(
        fmt.Sprintf("%s:%d", local, srv.config.HTTPPort()),
        srv.router.Handler,
    )
}

func (srv server) Handle(method, path string, handler fasthttp.RequestHandler) {
    srv.router.Handle(method, path, handler)
}

func (srv server) AuthHandle(method, path string, handler fasthttp.RequestHandler) {
    srv.Handle(method, path, Adapt(handler, srv.authorizationMechanism.CheckAuthorization))
}

func (srv server) ApiHandle(method, path string, handler fasthttp.RequestHandler) {
    srv.Handle(method, srv.config.ApiEndpoint() + path, Adapt(handler, srv.apiResponser.Writer))
}

func (srv server) ApiAuthHandle(method, path string, handler fasthttp.RequestHandler) {
    srv.Handle(method, srv.config.ApiEndpoint() + path,
        Adapt(handler, srv.apiResponser.Writer, srv.authorizationMechanism.CheckAuthorization),
    )
}
