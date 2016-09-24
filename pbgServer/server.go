package pbgServer

import (
    "github.com/valyala/fasthttp"
    "github.com/buaazp/fasthttprouter"
    "log"
)

type(
    pbgServer struct {
        // Mechanisms
        dataMechanism IDataMechanism
        authMechanism IAuthMechanism
        sessMechanism ISessionsMechanism
        // Private fields
        configuration IConfiguration
        httpRouter    *fasthttprouter.Router
    }

    pbgBuilder struct {
        config     IConfiguration
        dmCallback DMConstructor
        amCallback AMConstructor
        smCallback SMConstructor
    }

    PBGServer interface {
        IServerContext
        GetConfiguration() IConfiguration

        CheckInitialization() error
        StartServer()
    }

    PBGBuilder interface {
        UseConfiguration(IConfiguration) PBGBuilder
        UseDataMechanism(DMConstructor)  PBGBuilder
        UseAuthMechanism(AMConstructor)  PBGBuilder
        UseSessMechanism(SMConstructor)  PBGBuilder

        Build() PBGServer
    }

    // Costruttori callback: vogliamo costruire (in maniera
    // decisa dall'utente) i meccanismi di accesso alla memoria
    // in base alla particolare configurazione passata al builder
    DMConstructor func(cfg IConfiguration) IDataMechanism
    AMConstructor func(cfg IConfiguration) IAuthMechanism
    SMConstructor func(cfg IConfiguration) ISessionsMechanism
)

// Builder functions ----------------------------------------------------------- //
// ----------------------------------------------------------------------------- //
func Builder() PBGBuilder {
    return &pbgBuilder{}
}

func (builder *pbgBuilder) UseConfiguration(cfg IConfiguration) PBGBuilder {
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
    } else {
        dataMechanism := builder.dmCallback(builder.config)
        authMechanism := builder.amCallback(builder.config)
        sessMechanism := builder.smCallback(builder.config)

        return &pbgServer{
            dataMechanism: dataMechanism,
            authMechanism: authMechanism,
            sessMechanism: sessMechanism,
            configuration: builder.config,
            httpRouter:    fasthttprouter.New(),
        }
    }
}
// ----------------------------------------------------------------------------- //
// ----------------------------------------------------------------------------- //

// Server functions ------------------------------------------------------------ //
// ----------------------------------------------------------------------------- //
func (srv *pbgServer) StartServer() {
    log.Fatal(fasthttp.ListenAndServe(
        srv.configuration.GetListenAndServe(),
        srv.httpRouter.Handler,
    ))
}

func (srv *pbgServer) GetConfiguration() IConfiguration {
    return srv.configuration
}

func (srv *pbgServer) CheckInitialization() error {
    // TODO: finish this
    return nil
}
// ----------------------------------------------------------------------------- //
// ----------------------------------------------------------------------------- //
