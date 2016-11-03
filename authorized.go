package main

import (
    "github.com/valyala/fasthttp"
    "github.com/ar3s3ru/PokemonBattleGo/pbg"
)

const (
    mePath = "/me"
)

func handleMePath(ctx *fasthttp.RequestCtx) {
    session, ok := ctx.UserValue(pbg.SessionKey).(pbg.Session)
    if !ok {
        pbg.WriteAPIError(ctx, ErrInHandlerConversion, fasthttp.StatusInternalServerError)
        return
    }

    pbg.WriteAPISuccess(ctx, session.GetUserReference(), fasthttp.StatusOK)
}
