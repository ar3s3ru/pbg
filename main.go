package main

import (
    "fmt"
    "github.com/valyala/fasthttp"
    "github.com/buaazp/fasthttprouter"
    "github.com/ar3s3ru/PokemonBattleGo/pbgServer"
    "github.com/ar3s3ru/PokemonBattleGo/pbgServer/mem"
)

const (
    CfgPokèmonFile = "POKÈMON_FILE"
)

func handleHello(sctx pbgServer.IServerContext, ctx *fasthttp.RequestCtx, pm fasthttprouter.Params) {
    fmt.Fprintf(ctx, "Called \"/hello\" with %v\n", sctx)
}

func dmCallback(cfg pbgServer.Configuration) pbgServer.IDataMechanism {
    if pokèmonFile, err := cfg.GetValue(CfgPokèmonFile); err != nil {
        return mem.NewDataBuilder().UsePokèmonFile(pokèmonFile.(string)).Build()
    } else {
        fmt.Printf("err val: %v", err)
        panic(err)
    }
}

func amCallback(cfg pbgServer.Configuration) pbgServer.IAuthMechanism {
    fmt.Printf("Building Auth Mechanism using %v\n", cfg)
    return nil
}

func smCallback(cfg pbgServer.Configuration) pbgServer.ISessionsMechanism {
    fmt.Printf("Building Sess Mechanism using %v\n", cfg)
    return nil
}

func getServer() pbgServer.PBGServer {
    return pbgServer.Builder().UseConfiguration(
        pbgServer.NewConfig().SetHTTPPort(8080).SetValue(CfgPokèmonFile, "lol"),
    ).UseDataMechanism(dmCallback).UseAuthMechanism(amCallback).UseSessMechanism(smCallback).Build()
}

func main() {
    // Get server instance
    srv := getServer()
    // Handle HTTP request
    srv.Handle(pbgServer.GET, "/hello", handleHello).StartServer() // Start HTTP server
}
