package mem

type (
    MemConfiguration struct {
        httpPort  int
        lstAndSrv string
    }
)

func NewConfig(httpPort int) *MemConfiguration {
    return &MemConfiguration{
        httpPort:  httpPort,
        lstAndSrv: ":" + string(httpPort),
    }
}

func (cfg *MemConfiguration) GetHTTPPort() int {
    return cfg.httpPort
}

func (cfg *MemConfiguration) GetListenAndServe() string {
    return cfg.lstAndSrv
}
