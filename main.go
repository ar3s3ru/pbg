package main

import (
    "fmt"
    "github.com/valyala/fasthttp"
    "github.com/buaazp/fasthttprouter"
    "github.com/ar3s3ru/PokemonBattleGo/pbgServer"
    "github.com/ar3s3ru/PokemonBattleGo/pbgServer/mem"
)

func handleHello(sctx pbgServer.IServerContext, ctx *fasthttp.RequestCtx, pm fasthttprouter.Params) {
    fmt.Fprintf(ctx, "Called \"/hello\" with %v\n", sctx)
}

func getServer() pbgServer.PBGServer {
    return pbgServer.Builder().UseConfiguration(
        // Configuration from mem package
        mem.NewConfigBuilder().UseHTTPPort(8080).UsePokèmonFile("lol").Build(),
    ).UseDataMechanism(func(cfg pbgServer.IConfiguration) pbgServer.IDataMechanism {
        memCfg := cfg.(mem.MemConfiguration)
        return mem.NewDataBuilder().UsePokèmonFile(memCfg.GetPokèmonFile()).Build()
    }).UseAuthMechanism(func (cfg pbgServer.IConfiguration) pbgServer.IAuthMechanism {
        fmt.Printf("Building Auth Mechanism using %v\n", cfg)
        return nil
    }).UseSessMechanism(func (cfg pbgServer.IConfiguration) pbgServer.ISessionsMechanism {
        fmt.Printf("Building Sess Mechanism using %v\n", cfg)
        return nil
    }).Build()
}

func main() {
    // Get server instance
    srv := getServer()
    // Handle HTTP request
    srv.Handle(
        pbgServer.GET,
        "/hello",
        handleHello,
    ).StartServer()    // Start HTTP server
}
