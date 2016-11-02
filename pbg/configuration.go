package pbg

type (
    Configuration interface {
        HTTPPort()    int
        LocalHost()   bool
        APIEndpoint() string
    }

    BaseConfiguration struct {
        Port     int    `json:"port"`
        Local    bool   `json:"local"`
        Endpoint string `json:"endpoint"`
    }
)

func (cfg BaseConfiguration) HTTPPort() int {
    return cfg.Port
}

func (cfg BaseConfiguration) LocalHost() bool {
    return cfg.Local
}

func (cfg BaseConfiguration) APIEndpoint() string {
    return cfg.Endpoint
}
