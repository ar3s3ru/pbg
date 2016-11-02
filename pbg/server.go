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

        APIHandle(HTTPMethod, string, fasthttp.RequestHandler)
        APIAuthHandle(HTTPMethod, string, fasthttp.RequestHandler)
    }

    ServerBuilder interface {
        WithAPIResponser(APIResponser)                     ServerBuilder
        WithConfiguration(Configuration)                   ServerBuilder
        WithDataMechanism(DataMechanism)                   ServerBuilder
        WithSessionMechanism(SessionMechanism)             ServerBuilder
        WithAuthorizationMechanism(AuthorizationMechanism) ServerBuilder
        //WithLogger()                              ServerBuilder

        Build() Server
    }

    server struct {
        apiResponser           APIResponser
        dataMechanism          DataMechanism
        sessionMechanism       SessionMechanism
        authorizationMechanism AuthorizationMechanism

        config                 Configuration
        router                 *fasthttprouter.Router
    }

    serverBuilder func(
        Configuration, DataMechanism, SessionMechanism, AuthorizationMechanism, APIResponser,
    ) Server
)

var (
    // Base APIResponder is JSON
    jsonResponder = NewJSONResponser()
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
    return serverBuilder(func(cfg Configuration,
                dm DataMechanism, sm SessionMechanism, am AuthorizationMechanism, ar APIResponser) Server {
        // Parameters needed, checking if they have legit values
        check(cfg, ErrInvalidConfiguration)
        check(dm,  ErrUnspecifiedDataM)
        check(sm,  ErrUnspecifiedSessM)
        check(am,  ErrUnspecifiedAuthM)
        check(ar,  ErrInvalidAPIResponser)

        return &server{
            config: cfg,
            router: fasthttprouter.New(),
            // Module mechanisms
            dataMechanism: dm, sessionMechanism: sm, authorizationMechanism: am,
            apiResponser:  ar,
        }
    })
}

func (sb serverBuilder) WithConfiguration(cfg Configuration) ServerBuilder {
    return serverBuilder(func(_ Configuration,
                dm DataMechanism, sm SessionMechanism, am AuthorizationMechanism, ar APIResponser) Server {
        return sb(cfg, dm, sm, am, ar)
    })
}

func (sb serverBuilder) WithDataMechanism(dm DataMechanism) ServerBuilder {
    return serverBuilder(func(cfg Configuration,
                _ DataMechanism, sm SessionMechanism, am AuthorizationMechanism, ar APIResponser) Server {
        return sb(cfg, dm, sm, am, ar)
    })
}

func (sb serverBuilder) WithSessionMechanism(sm SessionMechanism) ServerBuilder {
    return serverBuilder(func(cfg Configuration,
                dm DataMechanism, _ SessionMechanism, am AuthorizationMechanism, ar APIResponser) Server {
        return sb(cfg, dm, sm, am, ar)
    })
}

func (sb serverBuilder) WithAuthorizationMechanism(am AuthorizationMechanism) ServerBuilder {
    return serverBuilder(func(cfg Configuration,
                dm DataMechanism, sm SessionMechanism, _ AuthorizationMechanism, ar APIResponser) Server {
        return sb(cfg, dm, sm, am, ar)
    })
}

func (sb serverBuilder) WithAPIResponser(ar APIResponser) ServerBuilder {
    return serverBuilder(func(cfg Configuration,
                dm DataMechanism, sm SessionMechanism, am AuthorizationMechanism, _ APIResponser) Server {
        return sb(cfg, dm, sm, am, ar)
    })
}

func (sb serverBuilder) Build() Server {
    return sb(baseConfig, nil, nil, nil, jsonResponder)
}

// Server
func (srv *server) Start() {
    fasthttp.ListenAndServe(
        fmt.Sprintf("%s:%d", addressing(srv.config), srv.config.HTTPPort()),
        srv.router.Handler,
    )
}
