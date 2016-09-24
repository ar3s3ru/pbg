package pbgServer

import "fmt"

type IConfiguration interface {
    GetListenAndServe() string
    GetHTTPPort()       int
}

type ObjectConfiguration struct {
    HTTPPort int
    // Private struct fields
    listenAndServePath string
}

func ObjectConfigurationBuilder(httpPort int) *ObjectConfiguration {
    return &ObjectConfiguration{
        HTTPPort: httpPort,
        listenAndServePath: fmt.Sprintf(":%d", httpPort),
    }
}

func (cfg *ObjectConfiguration) GetListenAndServe() string {
    return cfg.listenAndServePath
}

func (cfg *ObjectConfiguration) GetHTTPPort() int {
    return cfg.HTTPPort
}
