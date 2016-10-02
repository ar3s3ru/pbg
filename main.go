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
    CfgLayoutFile  = "LAYOUT_FILE"

    CfgMovesListFile = "MOVES_LIST_FILE"
    CfgMovesIdFile   = "MOVES_ID_FILE"

    CfgPokemonListFile = "POKEMON_LIST_FILE"
    CfgPokemonIdFile   = "POKEMON_ID_FILE"
)

func handleHello(sctx pbgServer.IServerContext, ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
    fmt.Fprintf(ctx, "Called \"/hello\" with %v\n", sctx)
}

func handlePokèmons(sctx pbgServer.IServerContext, ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
    if lay, htm :=
        sctx.GetConfiguration().GetValue(CfgLayoutFile),
        sctx.GetConfiguration().GetValue(CfgPokemonListFile); lay == nil || htm == nil {
        ctx.Error("HTML templates not configured properly", fasthttp.StatusInternalServerError)
    } else if t, err := template.ParseFiles(lay.(string), htm.(string)); err != nil {
        ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
    } else if err := t.Execute(ctx, sctx.GetDataMechanism().GetPokèmons()); err != nil {
        ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
    } else {
        ctx.SetContentType("text/html; charset=utf-8")
    }
}

func handlePokèmonId(sctx pbgServer.IServerContext, ctx *fasthttp.RequestCtx, pm fasthttprouter.Params) {
    // Contents
    if idx, err := strconv.Atoi(pm.ByName("id")); err != nil {
        ctx.Error("Invalid id used, must be an integer (ex. /pokemon/1, /pokemon/2, ...)", fasthttp.StatusBadRequest)
    } else if lay, htm :=
        sctx.GetConfiguration().GetValue(CfgLayoutFile),
        sctx.GetConfiguration().GetValue(CfgPokemonIdFile); lay == nil || htm == nil {
        ctx.Error("HTML templates not configured properly", fasthttp.StatusInternalServerError)
    } else if pkm, err := sctx.GetDataMechanism().GetPokèmonById(idx); err != nil {
        ctx.Error(err.Error(), fasthttp.StatusNotFound)
    } else if t, err := template.ParseFiles(lay.(string), htm.(string)); err != nil {
        ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
    } else if err := t.Execute(ctx, pkm); err != nil {
        ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
    } else {
        ctx.SetContentType("text/html; charset=utf-8")
    }
}

func handleMoves(sctx pbgServer.IServerContext, ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
    if lay, htm :=
        sctx.GetConfiguration().GetValue(CfgLayoutFile),
        sctx.GetConfiguration().GetValue(CfgMovesListFile); lay == nil || htm == nil {
        ctx.Error("HTML templates not configured properly", fasthttp.StatusInternalServerError)
    } else if t, err := template.ParseFiles(lay.(string), htm.(string)); err != nil {
        ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
    } else if err := t.Execute(ctx, sctx.GetDataMechanism().GetMoves()); err != nil {
        ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
    } else {
        ctx.SetContentType("text/html; charset=utf-8")
    }
}

func handleMoveId(sctx pbgServer.IServerContext, ctx *fasthttp.RequestCtx, pm fasthttprouter.Params) {
    // Contents
    if idx, err := strconv.Atoi(pm.ByName("id")); err != nil {
        ctx.Error("Invalid id used, must be an integer (ex. /move/1, /move/2, ...)", fasthttp.StatusBadRequest)
    } else if lay, htm :=
        sctx.GetConfiguration().GetValue(CfgLayoutFile),
        sctx.GetConfiguration().GetValue(CfgMovesIdFile); lay == nil || htm == nil {
        ctx.Error("HTML templates not configured properly", fasthttp.StatusInternalServerError)
    } else if pkm, err := sctx.GetDataMechanism().GetMoveById(idx); err != nil {
        ctx.Error(err.Error(), fasthttp.StatusNotFound)
    } else if t, err := template.ParseFiles(lay.(string), htm.(string)); err != nil {
        ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
    } else if err := t.Execute(ctx, pkm); err != nil {
        ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
    } else {
        ctx.SetContentType("text/html; charset=utf-8")
    }
}

func handleStatic(_ pbgServer.IServerContext, ctx *fasthttp.RequestCtx, pm fasthttprouter.Params) {
    pth := path.Join("static", pm.ByName("resource"))
    fasthttp.ServeFile(ctx, pth)
}

func dmCallback(cfg pbgServer.Configuration) pbgServer.IDataMechanism {
    if pokèmonFile := cfg.GetValue(CfgPokèmonFile); pokèmonFile == nil {
        panic("PokèmonFile not configured")
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
    return pbgServer.Builder(
        pbgServer.NewConfig().SetHTTPPort(8080).SetValue(
            CfgPokèmonFile, "pokedb.json",
        ).SetValue(
            CfgLayoutFile, "templates/layout.html",
        ).SetValue(
            CfgPokemonListFile, "templates/pokemons.html",
        ).SetValue(
            CfgPokemonIdFile, "templates/pokemon_id.html",
        ).SetValue(
            CfgMovesListFile, "templates/moves.html",
        ).SetValue(
            CfgMovesIdFile, "templates/move_id.html",
        ),
    ).UseDataMechanism(dmCallback).UseAuthMechanism(amCallback).UseSessMechanism(smCallback).Build()
}

func main() {
    // Get server instance
    srv := getServer()
    // Handle HTTP request
    srv.Handle(
        pbgServer.GET, "/pokemons", handlePokèmons,
    ).Handle(
        pbgServer.GET, "/pokemon/:id", handlePokèmonId,
    ).Handle(
        pbgServer.GET, "/moves", handleMoves,
    ).Handle(
        pbgServer.GET, "/move/:id", handleMoveId,
    ).Handle(
        pbgServer.GET, "/hello", handleHello,
    ).Handle(
        pbgServer.GET, "/static/*resource", handleStatic,
    ).StartServer() // Start HTTP server
}
