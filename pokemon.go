package main

import (
    "github.com/ar3s3ru/PokemonBattleGo/pbgServer"
    "github.com/valyala/fasthttp"
    "github.com/buaazp/fasthttprouter"
    "strconv"
)

const (
    CfgPokemonListFile = "POKEMON_LIST_FILE"
    CfgPokemonIdFile   = "POKEMON_ID_FILE"

    APIPokèmonList  = APIEndpoint + "/pokemon"
    APIPokèmonEntry = APIEndpoint + "/pokemon/:id"
)

func handlePokèmons(sctx pbgServer.IServerContext,
                     ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) (int, interface{}, error) {
    return fasthttp.StatusOK, sctx.GetDataMechanism().GetPokèmons(), nil
}

func handlePokèmonId(sctx pbgServer.IServerContext,
                      ctx *fasthttp.RequestCtx, pm fasthttprouter.Params) (int, interface{}, error) {
    if idx, err := strconv.Atoi(pm.ByName("id")); err != nil {
        return fasthttp.StatusBadRequest, nil, err
    } else if pkm, err := sctx.GetDataMechanism().GetPokèmonById(idx); err != nil {
        return fasthttp.StatusNotFound, nil, err
    } else {
        return fasthttp.StatusOK, pkm, nil
    }
}
