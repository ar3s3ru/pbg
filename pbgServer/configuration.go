package pbgServer

import "errors"

type (
    Configuration interface {
        GetValue(string)              (interface{}, error)
        SetValue(string, interface{}) Configuration

        // Needed from server for StartServer()
        SetHTTPPort(int) Configuration
        GetHTTPPort()    int
    }

    configuration struct {
        httpPort int
        values   map[string]interface{}
    }
)

var (
    ErrConfigValueNotFound = errors.New("Config value requested has not been found")
    ErrHTTPPortNotSet      = errors.New("Config HTTP port has not been set, use cfg.SetHTTPPort(port) first")
)

func NewConfig() Configuration {
    return &configuration{
        httpPort: -1,
        values:   make(map[string]interface{}),
    }
}

func (cfg *configuration) GetValue(key string) (interface{}, error) {
    if val := cfg.values[key]; val != nil {
        return val, nil
    } else {
        return nil, ErrConfigValueNotFound
    }
}

func (cfg *configuration) SetValue(key string, value interface{}) Configuration {
    cfg.values[key] = value
    return cfg
}

func (cfg *configuration) SetHTTPPort(port int) Configuration {
    cfg.httpPort = port
    return cfg
}

func (cfg *configuration) GetHTTPPort() int {
    return cfg.httpPort
}
