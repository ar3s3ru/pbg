package main

import (
    "github.com/ar3s3ru/PokemonBattleGo/pbgServer"
    "github.com/valyala/fasthttp"
    "github.com/buaazp/fasthttprouter"
    "strconv"
    "html/template"
)

const (
    CfgMovesListFile = "MOVES_LIST_FILE"
    CfgMovesIdFile   = "MOVES_ID_FILE"
)

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