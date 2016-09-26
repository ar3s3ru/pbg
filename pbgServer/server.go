package pbgServer

import (
    "github.com/valyala/fasthttp"
    "github.com/buaazp/fasthttprouter"
    "log"
    "fmt"
)

type(
    pbgServer struct {
        // Mechanisms
        dataMechanism IDataMechanism
        authMechanism IAuthMechanism
        sessMechanism ISessionsMechanism
        // Private fields
        configuration Configuration
        httpRouter    *fasthttprouter.Router
    }

    pbgBuilder struct {
        config     Configuration
        dmCallback DMConstructor
        amCallback AMConstructor
        smCallback SMConstructor
    }

    IServerContext interface {
        GetDataMechanism() IDataMechanism
        GetAuthMechanism() IAuthMechanism
        GetSessMechanism() ISessionsMechanism
    }

    PBGServer interface {
        IServerContext
        GetConfiguration() Configuration

        CheckInitialization() error
        StartServer()
        Handle(HTTPMethod, string, Handler) PBGServer
    }

    PBGBuilder interface {
        UseConfiguration(Configuration)  PBGBuilder
        UseDataMechanism(DMConstructor)  PBGBuilder
        UseAuthMechanism(AMConstructor)  PBGBuilder
        UseSessMechanism(SMConstructor)  PBGBuilder

        Build() PBGServer
    }

    // Costruttori callback: vogliamo costruire (in maniera
    // decisa dall'utente) i meccanismi di accesso alla memoria
    // in base alla particolare configurazione passata al builder
    DMConstructor func(cfg Configuration) IDataMechanism
    AMConstructor func(cfg Configuration) IAuthMechanism
    SMConstructor func(cfg Configuration) ISessionsMechanism
)

// Builder methods ------------------------------------------------------------- //
// ----------------------------------------------------------------------------- //
func Builder() PBGBuilder {
    return &pbgBuilder{}
}

func (builder *pbgBuilder) UseConfiguration(cfg Configuration) PBGBuilder {
    if cfg == nil {
        panic("Using nil as configuration is not allowed!")
    } else {
        builder.config = cfg
    }

    return builder
}

func (builder *pbgBuilder) UseDataMechanism(callback DMConstructor) PBGBuilder {
    if callback == nil {
        panic("Using nil as DataMechanism constructor callback is not allowed")
    } else {
        builder.dmCallback = callback
    }

    return builder
}

func (builder *pbgBuilder) UseAuthMechanism(callback AMConstructor) PBGBuilder {
    if callback == nil {
        panic("Using nil as AuthMechanism constructor callback is not allowed")
    } else {
        builder.amCallback = callback
    }

    return builder
}

func (builder *pbgBuilder) UseSessMechanism(callback SMConstructor) PBGBuilder {
    if callback == nil {
        panic("Using nil as SessMechanism constructor callback is not allowed")
    } else {
        builder.smCallback = callback
    }

    return builder
}

func (builder *pbgBuilder) Build() PBGServer {
    if builder.config == nil {
        panic("Configuration not set")
    }

    dataMechanism := builder.dmCallback(builder.config)
    authMechanism := builder.amCallback(builder.config)
    sessMechanism := builder.smCallback(builder.config)

    //switch {
    //case dataMechanism == nil:
    //    panic("DataMechanism not created in dedicated callback")
    //case authMechanism == nil:
    //    panic("AuthMechanism not created in dedicated callback")
    //case sessMechanism == nil:
    //    panic("SessMechanism not created in dedicated callback")
    //}

    return &pbgServer{
        dataMechanism: dataMechanism,
        authMechanism: authMechanism,
        sessMechanism: sessMechanism,
        configuration: builder.config,
        httpRouter:    fasthttprouter.New(),
    }
}
// ----------------------------------------------------------------------------- //
// ----------------------------------------------------------------------------- //

// Server methods -------------------------------------------------------------- //
// ----------------------------------------------------------------------------- //
func (srv *pbgServer) StartServer() {
    if port := srv.configuration.GetHTTPPort(); port != -1 {
        log.Fatal(fasthttp.ListenAndServe(
            fmt.Sprintf(":%d", port),
            srv.httpRouter.Handler,
        ))
    } else {
        panic(ErrHTTPPortNotSet)
    }
}

func (srv *pbgServer) GetConfiguration() Configuration {
    return srv.configuration
}

func (srv *pbgServer) GetDataMechanism() IDataMechanism {
    return srv.dataMechanism
}

func (srv *pbgServer) GetAuthMechanism() IAuthMechanism {
    return srv.authMechanism
}

func (srv *pbgServer) GetSessMechanism() ISessionsMechanism {
    return srv.sessMechanism
}

func (srv *pbgServer) CheckInitialization() error {
    // TODO: finish this
    return nil
}
// ----------------------------------------------------------------------------- //
// ----------------------------------------------------------------------------- //
