package main

import (
    "github.com/valyala/fasthttp"
    "github.com/ar3s3ru/PokemonBattleGo/pbg"
    "strconv"
)

const (
    pokèmonPath   = "/pokemon"
    pokèmonIdPath = pokèmonPath + "/:id"
)

func handlePokèmonList(ctx *fasthttp.RequestCtx) {
    dataMechanism, ok := ctx.UserValue(pbg.DataAccessKey).(pbg.DataMechanism)
    if !ok {
        // Error here
        return
    }

    pbg.WriteAPISuccess(ctx,
        dataMechanism.ListPokèmon(),
        fasthttp.StatusOK,
    )
}

func handlePokèmonId(ctx *fasthttp.RequestCtx) {
    dataMechanism, ok := ctx.UserValue(pbg.DataAccessKey).(pbg.DataMechanism)
    if !ok {
        // Error here
        return
    }

    strArg, ok := ctx.UserValue("id").(string)
    if !ok {
        // Error her
        return
    }

    id, err := strconv.Atoi(strArg)
    if err != nil {
        pbg.WriteAPIError(ctx, err, fasthttp.StatusInternalServerError)
        return
    }

    pokèmon, err := dataMechanism.GetPokèmon(id)
    if err != nil {
        pbg.WriteAPIError(ctx, err, fasthttp.StatusNotFound)
        return
    }

    pbg.WriteAPISuccess(ctx, pokèmon, fasthttp.StatusOK)
}