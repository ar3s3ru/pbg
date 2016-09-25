package mem

import (
    "github.com/ar3s3ru/PokemonBattleGo/pbgServer"
    "fmt"
)

type (
    memConfiguration struct {
        httpPort     int
        lstAndSrv    string
        pokèmonFile  string
        trainersFile string
    }

    memBuilder struct {
        httpPort     int
        pokèmonFile  string
        trainersFile string
    }

    MemConfiguration interface {
        pbgServer.IConfiguration

        GetPokèmonFile()  string
        GetTrainersFile() string
    }

    CfgBuilder interface {
        UseHTTPPort(int)        CfgBuilder
        UsePokèmonFile(string)  CfgBuilder
        UseTrainersFile(string) CfgBuilder

        Build() pbgServer.IConfiguration
    }
)

func NewConfigBuilder() CfgBuilder {
    return &memBuilder{}
}

func (builder *memBuilder) UseHTTPPort(port int) CfgBuilder {
    builder.httpPort = port
    return builder
}

func (builder *memBuilder) UsePokèmonFile(file string) CfgBuilder {
    builder.pokèmonFile = file
    return builder
}

func (builder *memBuilder) UseTrainersFile(file string) CfgBuilder {
    builder.trainersFile = file
    return builder
}

func (builder *memBuilder) Build() pbgServer.IConfiguration {
    return &memConfiguration{
        httpPort:     builder.httpPort,
        lstAndSrv:    fmt.Sprintf(":%d", builder.httpPort),
        pokèmonFile:  builder.pokèmonFile,
        trainersFile: builder.trainersFile,
    }
}

func (cfg *memConfiguration) GetHTTPPort() int {
    return cfg.httpPort
}

func (cfg *memConfiguration) GetListenAndServe() string {
    return cfg.lstAndSrv
}

func (cfg *memConfiguration) GetPokèmonFile() string {
    return cfg.pokèmonFile
}

func (cfg *memConfiguration) GetTrainersFile() string {
    return cfg.trainersFile
}
