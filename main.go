package main

import (
    "fmt"
    "github.com/valyala/fasthttp"
    "github.com/buaazp/fasthttprouter"
    "github.com/ar3s3ru/PokemonBattleGo/pbgServer"
    "github.com/ar3s3ru/PokemonBattleGo/pbgServer/mem"
    "strconv"
)

const (
    CfgPokèmonFile = "POKÈMON_FILE"
)

func handleHello(sctx pbgServer.IServerContext, ctx *fasthttp.RequestCtx, pm fasthttprouter.Params) {
    fmt.Fprintf(ctx, "Called \"/hello\" with %v\n", sctx)
}

func handlePokèmon(sctx pbgServer.IServerContext, ctx *fasthttp.RequestCtx, pm fasthttprouter.Params) {
    if id := pm.ByName("id"); id == "" {
        fmt.Fprintln(ctx, "Invalid id used")
    } else if idx, err := strconv.Atoi(id); err != nil {
        fmt.Fprintf(ctx, "Error: %v", err)
    } else if pkm, err := sctx.GetDataMechanism().GetPokèmonById(idx); err != nil {
        fmt.Fprintf(ctx, "Error occurred: %s", err)
    } else {
        fmt.Fprintf(ctx, "Pokèmon: %s\n", pkm.GetName())
        fmt.Fprintf(ctx, "    Type: %s\n", pkm.GetType().String())
        fmt.Fprintf(ctx, "    Number: %d\n", pkm.GetPokèdex())
        fmt.Fprintf(ctx, "    Base Stats: %v\n", pkm.GetBaseStats())
    }
}

func dmCallback(cfg pbgServer.Configuration) pbgServer.IDataMechanism {
    if pokèmonFile, err := cfg.GetValue(CfgPokèmonFile); err != nil {
        panic(err)
    } else {
        return mem.NewDataBuilder().UsePokèmonFile(pokèmonFile.(string)).Build()
    }
}

func smCallback(cfg pbgServer.Configuration, dm pbgServer.IDataMechanism) pbgServer.ISessionMechanism {
    fmt.Printf("Using dataMechanism: %v\n", dm)
    return mem.AuthBuilder().UseDataMechanism(dm).Build()
}

func amCallback(cfg pbgServer.Configuration, sm pbgServer.ISessionMechanism) pbgServer.IAuthMechanism {
    fmt.Printf("Using sessMechanism: %v\n", sm)
    return sm.(mem.IAuthority)
}

func getServer() pbgServer.PBGServer {
    return pbgServer.Builder().UseConfiguration(
        pbgServer.NewConfig().SetHTTPPort(8080).SetValue(CfgPokèmonFile, "pokedexTest.json"),
    ).UseDataMechanism(dmCallback).UseAuthMechanism(amCallback).UseSessMechanism(smCallback).Build()
}

func main() {
    // Get server instance
    srv := getServer()
    // Handle HTTP request
    srv.Handle(
        pbgServer.GET, "/pokemon/:id", handlePokèmon,
    ).StartServer() // Start HTTP server
}
