package main

import (
	"errors"

	"github.com/ar3s3ru/PokemonBattleGo/mem"
	"github.com/ar3s3ru/PokemonBattleGo/pbg"
	"github.com/valyala/fasthttp"
)

const MePath = "/me"

var (
	authorizationHeader         = []byte("Authorization")
	ErrSessionInterfaceNotFound = errors.New("SessionInterface not found, sure you adapted with pbg.WithSessionAccess ?")
)

func authorizedHandler(srv pbg.Server) pbg.Adapter {
	return func(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
		return srv.WithSessionDBAccess(func(ctx *fasthttp.RequestCtx) {
			si, ok := ctx.UserValue(pbg.SessionDBInterfaceKey).(pbg.SessionInterface)
			if !ok || si == nil {
				pbg.WriteAPIError(ctx, ErrSessionInterfaceNotFound, fasthttp.StatusInternalServerError)
				return
			}

			token, err := mem.BasicAuthorization(ctx.Request.Header.PeekBytes(authorizationHeader))
			if err != nil {
				pbg.WriteAPIError(ctx, err, fasthttp.StatusInternalServerError)
				return
			}

			session, err := si.GetSession(string(token))
			if err != nil {
				pbg.WriteAPIError(ctx, err, fasthttp.StatusNotFound)
				return
			}

			ctx.SetUserValue(pbg.SessionKey, session)
			handler(ctx)
		})
	}
}

func handleMePath(ctx *fasthttp.RequestCtx) {
	session, ok := ctx.UserValue(pbg.SessionKey).(pbg.Session)
	if !ok {
		pbg.WriteAPIError(ctx, ErrInHandlerConversion, fasthttp.StatusInternalServerError)
		return
	}

	pbg.WriteAPISuccess(ctx, session.Reference(), fasthttp.StatusOK)
}
