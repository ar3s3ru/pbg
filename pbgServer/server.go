package pbgServer

import (
    "github.com/valyala/fasthttp"
    "github.com/buaazp/fasthttprouter"
    "log"
    "fmt"
    "errors"
)

type(
    pbgServer struct {
        // Mechanisms ----------------------
        dataMechanism IDataMechanism
        authMechanism IAuthMechanism
        sessMechanism ISessionMechanism
        // HTTP stuff ----------------------
        apiResponse   APIResponse
        // Private fields ------------------
        configuration Configuration
        httpRouter    *fasthttprouter.Router
    }

    pbgBuilder struct {
        config     Configuration
        // Callbacks -----------
        dmCallback DMConstructor
        amCallback AMConstructor
        smCallback SMConstructor
        apiRespons APIResponse
    }

    // Interfaccia che espone le funzionalità riguardo la gestione dei dati internamente al server.
    // Permette dunque l'accesso ad eventuale database e meccanismo di autenticazione all'interno
    // delle callback registrate dal server per gestire le varie richieste HTTP.
    //
    // Interface that expose functionalities regarding server internal data managing.
    // In other words, it allows database and authentication mechanisms access inside the callbacks
    // registered to handle all the HTTP requests.
    IServerContext interface {
        GetConfiguration() Configuration
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
        // HTTP stuff
        IHTTPServer
        // Start server, of course...
        StartServer()
    }

    // Builder per la costruzione di oggetti PBGServer.
    // Si occupa di registrare un oggetto di configurazione e delle callback per la creazione dei
    // vari meccanismi che si occupano dell'accesso ai dati.
    //
    // Builder used for PBGServer objects' construction.
    // It registers configuration object and callbacks for the creation of the various mechanisms that
    // handles data access throughout the server itself.
    PBGBuilder interface {
        // Method chaining!
        UseDataMechanism(DMConstructor) PBGBuilder
        UseAuthMechanism(AMConstructor) PBGBuilder
        UseSessMechanism(SMConstructor) PBGBuilder
        // API response handler
        UseAPIResponse(APIResponse)     PBGBuilder

        Build() PBGServer
    }

    // Costruttore callback per l'oggetto che gestisce l'accesso ai dati del server.
    //
    // Callback constructor for the object that handles server's data access.
    DMConstructor func(Configuration)                    IDataMechanism
    // Costruttore callback per l'oggetto che gestisce la creazione e la gestione delle sessioni utente.
    //
    // Callback constructor for the object that handles user sessions' creation and management.
    SMConstructor func(Configuration, IDataMechanism)    ISessionMechanism
    // Costruttore callback per l'oggetto che gestisce la logica di autorizzazione utilizzato nel server.
    //
    // Callback constructor for the object that handles authorization logic used throughout the server.
    AMConstructor func(Configuration, ISessionMechanism) IAuthMechanism
)

var (
    ErrMethodNotAllowed = errors.New(fasthttp.StatusMessage(fasthttp.StatusMethodNotAllowed))
    ErrNotFound         = errors.New(fasthttp.StatusMessage(fasthttp.StatusNotFound))
)

/**
    Builder methods -------------------------------------------------------------
  */

// Crea e restituisce un oggetto PBGBuilder, per la conseguente creazione di oggetti PBGServer.
//
// Creates and returns a PBGBuilder object, used for the consequent creation of PBGServer objects.
func Builder(config Configuration) PBGBuilder {
    if config == nil {
        panic("Configuration not set")
    }

    return &pbgBuilder{
        config:  config,
    }
}

// Registra una callback per la creazione di un oggetto DataMechanism.
//
// Registers a callback fo// Registra una callback per la creazione di un oggetto DataMechanism.
//
// Registers a callback for the creation of a DataMechanism object.r the creation of a DataMechanism object.
func (builder *pbgBuilder) UseDataMechanism(callback DMConstructor) PBGBuilder {
    if callback == nil {
        panic("Using nil as DataMechanism constructor callback is not allowed")
    } else {
        builder.dmCallback = callback
    }

    return builder
}

// Registra una callback per la creazione di un oggetto AuthMechanism.
//
// Registers a callback for the creation of a AuthMechanism object.
func (builder *pbgBuilder) UseAuthMechanism(callback AMConstructor) PBGBuilder {
    if callback == nil {
        panic("Using nil as AuthMechanism constructor callback is not allowed")
    } else {
        builder.amCallback = callback
    }

    return builder
}

// Registra una callback per la creazione di un oggetto SessionMechanism.
//
// Registers a callback for the creation of a SessionMechanism object.
func (builder *pbgBuilder) UseSessMechanism(callback SMConstructor) PBGBuilder {
    if callback == nil {
        panic("Using nil as SessMechanism constructor callback is not allowed")
    } else {
        builder.smCallback = callback
    }

    return builder
}

func (builder *pbgBuilder) UseAPIResponse(response APIResponse) PBGBuilder {
    builder.apiRespons = response
    return builder
}

// Costruisce l'oggetto PBGServer dal builder.
// Prima di chiamare questo metodo, assicurarsi che la configurazione passata e le callback siano valide;
// in particolare, le callback devono creare 'veri' oggetti: non possono restituire <nil> come valore di ritorno.
//
// Builds a PBGServer object from the builder.
// Be sure that configuration and callbacks registered are valid, before calling this method;
// particularly, the callbacks must create 'real' objects: they can't return <nil> as return value.
func (builder *pbgBuilder) Build() PBGServer {
    var (
        dataMechanism IDataMechanism
        sessMechanism ISessionMechanism
        authMechanism IAuthMechanism
    )

    if dataMechanism = builder.dmCallback(builder.config); dataMechanism == nil {
        panic("DataMechanism not created in dedicated callback")
    } else if sessMechanism = builder.smCallback(builder.config, dataMechanism); sessMechanism == nil {
        panic("SessMechanism not created in dedicated callback")
    } else if authMechanism = builder.amCallback(builder.config, sessMechanism); authMechanism == nil {
        panic("AuthMechanism not created in dedicated callback")
    } else if builder.apiRespons == nil {
        panic("Invalid APIResponse callback used")
    }

    router := fasthttprouter.New()
    // Method not allowed handling
    router.MethodNotAllowed = func (ctx *fasthttp.RequestCtx) {
        builder.apiRespons(fasthttp.StatusMethodNotAllowed, nil, ErrMethodNotAllowed, ctx)
    }
    // Not found handling
    router.NotFound = func (ctx *fasthttp.RequestCtx) {
        builder.apiRespons(fasthttp.StatusNotFound, nil, ErrNotFound, ctx)
    }

    return &pbgServer{
        dataMechanism: dataMechanism,
        authMechanism: authMechanism,
        sessMechanism: sessMechanism,
        configuration: builder.config,
        httpRouter:    router,
        apiResponse:   builder.apiRespons,
    }
}

/**
   Server methods --------------------------------------------------------------
*/

// Inizia l'ascolto di richieste HTTP sulla porta specificata nella configurazione usata.
//
// Starts HTTP requests listening on the port specified into the configuration used.
func (srv *pbgServer) StartServer() {
    if port := srv.configuration.GetHTTPPort(); port != -1 {
        log.Printf("Starting server on :%d\n", port)
        log.Fatal(fasthttp.ListenAndServe(
            fmt.Sprintf(":%d", port),
            srv.httpRouter.Handler,
        ))
    } else {
        panic(ErrHTTPPortNotSet)
    }
}

// Ritorna la configurazione usata correntemente dal server.
//
// Returns the actual configuration used by the server.
func (srv *pbgServer) GetConfiguration() Configuration {
    return srv.configuration
}

// Ritorna il meccanismo di accesso ai dati usato dal server.
//
// Returns data access mechanism used by the server.
func (srv *pbgServer) GetDataMechanism() IDataMechanism {
    return srv.dataMechanism
}

// Ritorna il meccanismo di autenticazione usato dal server.
//
// Returns the authorization mechanism used by the server.
func (srv *pbgServer) GetAuthMechanism() IAuthMechanism {
    return srv.authMechanism
}

// Ritorna il meccanismo di gestione di sessioni utente usato dal server.
//
// Returns the user session management mechanism used by the server.
func (srv *pbgServer) GetSessMechanism() ISessionMechanism {
    return srv.sessMechanism
}
