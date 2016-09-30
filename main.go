package main

import (
    "fmt"
    "github.com/valyala/fasthttp"
    "github.com/buaazp/fasthttprouter"
    "github.com/ar3s3ru/PokemonBattleGo/pbgServer"
    "github.com/ar3s3ru/PokemonBattleGo/pbgServer/mem"
    "strconv"
    "html/template"
    "path"
)

const (
    CfgPokèmonFile = "POKÈMON_FILE"
)

func handleHello(sctx pbgServer.IServerContext, ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
    fmt.Fprintf(ctx, "Called \"/hello\" with %v\n", sctx)
}

func handlePokèmon(sctx pbgServer.IServerContext, ctx *fasthttp.RequestCtx, pm fasthttprouter.Params) {
    if idx, err := strconv.Atoi(pm.ByName("id")); err != nil {
        ctx.Error("Invalid id used, must be an integer (ex. /pokemon/1, /pokemon/2, ...)", fasthttp.StatusBadRequest)
    } else if pkm, err := sctx.GetDataMechanism().GetPokèmonById(idx); err != nil {
        ctx.Error(err.Error(), fasthttp.StatusNotFound)
    } else if t, err := template.ParseFiles(
        path.Join("templates", "layout.html"),
        path.Join("templates", "pokemon_id.html"),
    ); err != nil {
        ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
    } else if err := t.Execute(ctx, pkm); err != nil {
        ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
    } else {
        ctx.SetContentType("text/html")
    }
}

func handleStatic(_ pbgServer.IServerContext, ctx *fasthttp.RequestCtx, pm fasthttprouter.Params) {
    pth := path.Join("static", pm.ByName("resource"))
    fasthttp.ServeFile(ctx, pth)
}

func dmCallback(cfg pbgServer.Configuration) pbgServer.IDataMechanism {
    if pokèmonFile, err := cfg.GetValue(CfgPokèmonFile); err != nil {
        panic(err)
    } else {
        return mem.NewDataBuilder().UsePokèmonFile(pokèmonFile.(string)).Build()
    }
}

func smCallback(_ pbgServer.Configuration, dm pbgServer.IDataMechanism) pbgServer.ISessionMechanism {
    return mem.AuthBuilder().UseDataMechanism(dm).Build()
}

func amCallback(_ pbgServer.Configuration, sm pbgServer.ISessionMechanism) pbgServer.IAuthMechanism {
    return sm.(mem.IAuthority)
}

func getServer() pbgServer.PBGServer {
    return pbgServer.Builder().UseConfiguration(
        pbgServer.NewConfig().SetHTTPPort(8080).SetValue(CfgPokèmonFile, "pokedb.json"),
    ).UseDataMechanism(dmCallback).UseAuthMechanism(amCallback).UseSessMechanism(smCallback).Build()
}

func main() {
    // Get server instance
    srv := getServer()
    // Handle HTTP request
    srv.Handle(
        pbgServer.GET, "/pokemon/:id", handlePokèmon,
    ).Handle(
        pbgServer.GET, "/hello", handleHello,
    ).Handle(
        pbgServer.GET, "/static/:resource", handleStatic,
    ).StartServer() // Start HTTP server
}
