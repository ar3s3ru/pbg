package main

import (
    "github.com/ar3s3ru/PokemonBattleGo/pbgServer"
    "github.com/valyala/fasthttp"
    "github.com/buaazp/fasthttprouter"
    "strconv"
    "html/template"
)

const (
    CfgPokemonListFile = "POKEMON_LIST_FILE"
    CfgPokemonIdFile   = "POKEMON_ID_FILE"
)

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
