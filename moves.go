package main

import (
    "github.com/ar3s3ru/PokemonBattleGo/pbgServer"
    "github.com/valyala/fasthttp"
    "github.com/buaazp/fasthttprouter"
    "strconv"
)

const (
    CfgMovesListFile = "MOVES_LIST_FILE"
    CfgMovesIdFile   = "MOVES_ID_FILE"
)

func handleMoves(sctx pbgServer.IServerContext,
                  ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) (int, interface{}, error) {
    return fasthttp.StatusOK, sctx.GetDataMechanism().GetMoves(), nil
}

func handleMoveId(sctx pbgServer.IServerContext,
                   ctx *fasthttp.RequestCtx, pm fasthttprouter.Params) (int, interface{}, error) {
    // Contents
    if idx, err := strconv.Atoi(pm.ByName("id")); err != nil {
        return fasthttp.StatusBadRequest, nil, err
    } else if pkm, err := sctx.GetDataMechanism().GetMoveById(idx); err != nil {
        return fasthttp.StatusNotFound, nil, err
    } else {
        return fasthttp.StatusOK, pkm, nil
    }
}