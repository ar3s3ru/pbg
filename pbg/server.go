package pbg

import (
    "fmt"

    "github.com/buaazp/fasthttprouter"
    "github.com/valyala/fasthttp"
)

type (
    Server interface {
        Start()

        Handle(HTTPMethod, string, fasthttp.RequestHandler)
        AuthHandle(HTTPMethod, string, fasthttp.RequestHandler)

        ApiHandle(HTTPMethod, string, fasthttp.RequestHandler)
        ApiAuthHandle(HTTPMethod, string, fasthttp.RequestHandler)
    }

    ServerBuilder interface {
        WithConfiguration(Configuration)          ServerBuilder
        WithDataMechanism(DataMechanism)          ServerBuilder
        WithSessionMechanism(DataMechanism)       ServerBuilder
        WithAuthorizationMechanism(DataMechanism) ServerBuilder
        WithApiResponser(ApiResponser)            ServerBuilder
        //WithLogger()                   ServerBuilder

        Build() Server
    }

    server struct {
        apiResponser           ApiResponser
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
    // Base configuration
    baseConfig = &BaseConfiguration{Port: 8080, Local: true}
    // Check function (to show how many ways there are to declare fuctions in Go)
    check = func(data interface{}, err error) {
        if data == nil {
            panic(err)
        }
    }
    // Returns hostname string from Configuration
    addressing = func(cfg Configuration) string {
        if cfg.LocalHost() {
            return "localhost"
        } else {
            return ""
        }
    }
)

// Builder
func NewServerBuilder() ServerBuilder {
    return func(cfg Configuration,
                dm DataMechanism, sm SessionMechanism, am AuthorizationMechanism, ar ApiResponser) server {
        // Parameters needed, checking if they have legit values
        check(cfg, ErrUnspecifiedConfig)
        check(dm,  ErrUnspecifiedDataM)
        check(sm,  ErrUnspecifiedSessM)
        check(am,  ErrUnspecifiedAuthM)

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
    fasthttp.ListenAndServe(
        fmt.Sprintf("%s:%d", addressing(srv.config), srv.config.HTTPPort()),
        srv.router.Handler,
    )
}
