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
        sessMechanism ISessionMechanism
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

    // Interfaccia che espone le funzionalità riguardo la gestione dei dati internamente al server.
    // Permette dunque l'accesso ad eventuale database e meccanismo di autenticazione all'interno
    // delle callback registrate dal server per gestire le varie richieste HTTP.
    //
    // Interface that expose functionalities regarding server internal data managing.
    // In other words, it allows database and authentication mechanisms access inside the callbacks
    // registered to handle all the HTTP requests.
    IServerContext interface {
        GetDataMechanism() IDataMechanism
        GetAuthMechanism() IAuthMechanism
        GetSessMechanism() ISessionMechanism
    }

    // Rappresenta l'oggetto Server col quale l'utente interagisce.
    // Ha una relazione di composizione con IServerContext; inoltre, espone la configurazione
    // correntemente usata, la registrazione di handler per la gestione di richieste HTTP, e la messa in ascolto
    // di tali richieste sulla porta specificata nella particolare configurazione del server.
    //
    // Represents the Server object the user interacts with.
    // It has a composition relationship with IServerContext; moreover, it exposes the actual used configuration,
    // handlers registration for HTTP requests, and a ListenAndServe() method to start listening for HTTP request onto
    // the particular port specified into the configuration actually used.
    PBGServer interface {
        IServerContext
        GetConfiguration() Configuration

        Handle(HTTPMethod, string, Handler) PBGServer
        StartServer()
    }

    // TODO: finish documentation
    PBGBuilder interface {
        UseConfiguration(Configuration) PBGBuilder
        UseDataMechanism(DMConstructor) PBGBuilder
        UseAuthMechanism(AMConstructor) PBGBuilder
        UseSessMechanism(SMConstructor) PBGBuilder

        Build() PBGServer
    }

    // Costruttori callback: vogliamo costruire (in maniera
    // decisa dall'utente) i meccanismi di accesso alla memoria
    // in base alla particolare configurazione passata al builder
    DMConstructor func(Configuration)                    IDataMechanism
    SMConstructor func(Configuration, IDataMechanism)    ISessionMechanism
    AMConstructor func(Configuration, ISessionMechanism) IAuthMechanism
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

    var (
        dataMechanism IDataMechanism
        sessMechanism ISessionMechanism
        authMechanism IAuthMechanism
    )

    if dataMechanism = builder.dmCallback(builder.config); dataMechanism == nil {
        panic("DataMechanism not created in dedicated callback")
    }

    if sessMechanism = builder.smCallback(builder.config, dataMechanism); sessMechanism == nil {
        panic("SessMechanism not created in dedicated callback")
    }

    if authMechanism = builder.amCallback(builder.config, sessMechanism); authMechanism == nil {
        panic("AuthMechanism not created in dedicated callback")
    }

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

func (srv *pbgServer) GetSessMechanism() ISessionMechanism {
    return srv.sessMechanism
}

func (srv *pbgServer) CheckInitialization() error {
    // TODO: finish this
    return nil
}
// ----------------------------------------------------------------------------- //
// ----------------------------------------------------------------------------- //
