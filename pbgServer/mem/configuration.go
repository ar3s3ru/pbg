package mem

import (
    "github.com/ar3s3ru/PokemonBattleGo/pbgServer"
    "fmt"
)

type (
    memConfiguration struct {
        httpPort  int
        lstAndSrv string
    }

    memBuilder struct {
        httpPort int
    }

    CfgBuilder interface {
        UseHTTPPort(int) CfgBuilder
        Build() pbgServer.IConfiguration
    }
)

func ConfigBuilder() CfgBuilder {
    return &memBuilder{}
}

func (builder *memBuilder) UseHTTPPort(port int) CfgBuilder {
    builder.httpPort = port
    return builder
}

func (builder *memBuilder) Build() pbgServer.IConfiguration {
    return &memConfiguration{
        httpPort:  builder.httpPort,
        lstAndSrv: fmt.Sprintf(":%d", builder.httpPort),
    }
}

func (cfg *memConfiguration) GetHTTPPort() int {
    return cfg.httpPort
}

func (cfg *memConfiguration) GetListenAndServe() string {
    return cfg.lstAndSrv
}
