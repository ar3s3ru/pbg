package pbg

const (
    // Context key per i dati delle risposte negli handler API
    APIDataKey = "apiData"
    // Context key per gli errori delle risposte negli handler API
    APIErrorKey = "apiError"

    // Context key per accede al logger negli handler
    // Disponibile con l'adapter Server.WithLogger()
    LoggerKey = "logger"
    // Context key per accedere alla sessione autenticata negli handler
    // Usare un decoratore per controllare il token d'autorizzazione nelle richieste HTTP
    SessionKey = "session"

    // Context key per accedere all'interfaccia col DB dei modelli del server
    // Disponibile col decoratore Server.WithDataAccess()
    DataInterfaceKey    = "dataInterface"
    // Context key per accedere all'interfaccia col DB delle sessioni del server
    // Disponibile col decoratore Server.WithSessionAccess()
    SessionInterfaceKey = "sessionInterface"
)
