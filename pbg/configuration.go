package pbg

type (
    Configuration interface {
        HTTPPort()    int
        LocalHost()   bool
        ApiEndpoint() string
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

func (cfg BaseConfiguration) ApiEndpoint() string {
    return cfg.Endpoint
}
