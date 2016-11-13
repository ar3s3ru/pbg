package main

import (
    "strconv"

    "github.com/valyala/fasthttp"
    "github.com/ar3s3ru/PokemonBattleGo/pbg"
)

const (
    PokèmonPath   = "/pokemon"
    PokèmonIdPath = PokèmonPath + "/:id"
)

func handlePokèmonList(ctx *fasthttp.RequestCtx) {
    _, ok := ctx.UserValue(pbg.DataInterfaceKey).(pbg.DataInterface)
    if !ok {
        // Error here
        pbg.WriteAPIError(ctx, ErrInHandlerConversion, fasthttp.StatusInternalServerError)
        return
    }

    //pbg.WriteAPISuccess(ctx,
    //    di.ListPokèmon(),
    //    fasthttp.StatusOK,
    //)
}

func handlePokèmonId(ctx *fasthttp.RequestCtx) {
    di, ok := ctx.UserValue(pbg.DataInterfaceKey).(pbg.DataInterface)
    if !ok {
        // Error here
        pbg.WriteAPIError(ctx, ErrInHandlerConversion, fasthttp.StatusInternalServerError)
        return
    }

    strArg, ok := ctx.UserValue("id").(string)
    if !ok {
        // Error here
        pbg.WriteAPIError(ctx, ErrInHandlerConversion, fasthttp.StatusInternalServerError)
        return
    }

    id, err := strconv.Atoi(strArg)
    if err != nil {
        pbg.WriteAPIError(ctx, err, fasthttp.StatusInternalServerError)
        return
    }

    pokèmon, err := di.GetPokèmon(id)
    if err != nil {
        pbg.WriteAPIError(ctx, err, fasthttp.StatusNotFound)
        return
    }

    pbg.WriteAPISuccess(ctx, pokèmon, fasthttp.StatusOK)
}