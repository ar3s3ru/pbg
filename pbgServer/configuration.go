package pbgServer

import "errors"

type (
    // Interfaccia che rappresenta la configurazione del server.
    //
    // Interface that represent server configuration.
    Configuration interface {
        GetValue(string)              interface{}   // Getter
        SetValue(string, interface{}) Configuration // Setter

        // Needed from server for StartServer()
        SetHTTPPort(int) Configuration // Setter
        GetHTTPPort()    int           // Getter
    }

    configuration struct {
        httpPort int
        values   map[string]interface{}
    }
)

var (
    ErrHTTPPortNotSet = errors.New("Config HTTP port has not been set, use cfg.SetHTTPPort(port) first")
)

// Restituisce un nuovo oggetto Configuration pronto per essere usato.
//
// Returns a new Configuration object, ready to be used.
func NewConfig() Configuration {
    return &configuration{
        httpPort: 80,
        values:   make(map[string]interface{}),
    }
}

// Recupera un valore dalla configurazione attuale.
//
// Gets back a value from the current configuration.
func (cfg *configuration) GetValue(key string) interface{} {
    return cfg.values[key]
}

// Inserisce un valore nella configurazione attuale.
//
// Inserts a value into the current configuration.
func (cfg *configuration) SetValue(key string, value interface{}) Configuration {
    cfg.values[key] = value
    return cfg
}

// Setta la porta HTTP usata nella configurazione corrente.
//
// Sets the HTTP port used in the current configuration.
func (cfg *configuration) SetHTTPPort(port int) Configuration {
    cfg.httpPort = port
    return cfg
}

// Recupera la porta HTTP correntemente utilizzata.
//
// Gets the HTTP port currently used.
func (cfg *configuration) GetHTTPPort() int {
    return cfg.httpPort
}
