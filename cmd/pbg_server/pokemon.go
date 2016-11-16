package main

import (
	"strconv"

	"github.com/ar3s3ru/pbg"
	"github.com/valyala/fasthttp"
)

const (
	PokèmonPath   = "/pokemon"
	PokèmonIdPath = PokèmonPath + "/:id"
)

func handlePokèmonList(ctx *fasthttp.RequestCtx) {
	pokèmonDB, ok := ctx.UserValue(pbg.PokèmonDBInterfaceKey).(pbg.PokèmonDBInterface)
	if !ok {
		// Error here
		pbg.WriteAPIError(ctx, ErrInHandlerConversion, fasthttp.StatusInternalServerError)
		return
	}

	pbg.WriteAPISuccess(ctx, pokèmonDB.GetPokèmons(), fasthttp.StatusOK)
}

func handlePokèmonId(ctx *fasthttp.RequestCtx) {
	pokèmonDB, ok := ctx.UserValue(pbg.PokèmonDBInterfaceKey).(pbg.PokèmonDBInterface)
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

	pokèmon, err := pokèmonDB.GetPokèmon(id)
	if err != nil {
		pbg.WriteAPIError(ctx, err, fasthttp.StatusNotFound)
		return
	}

	pbg.WriteAPISuccess(ctx, pokèmon, fasthttp.StatusOK)
}
