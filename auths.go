package main

import (
    "github.com/ar3s3ru/PokemonBattleGo/pbgServer"
    "github.com/valyala/fasthttp"
    "github.com/buaazp/fasthttprouter"
    "encoding/json"
    "fmt"
)

func getPostBody(postBody []byte) (user string, pass string, err error) {
    req := struct {
        Username string `json:"username"`
        Password string `json:"password"`
    } {}

    if err = json.Unmarshal(postBody, &req); err != nil {
        return
    }

    user, pass, err = req.Username, req.Password, nil
    return
}

func handleRegister(sctx pbgServer.IServerContext, ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
    if un, pw, err := getPostBody(ctx.PostBody()); err != nil {
        ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
    } else if u, id, err := sctx.GetAuthMechanism().Register(un, pw); err != nil {
        ctx.Error(err.Error(), fasthttp.StatusBadRequest)
    } else {
        ctx.SetStatusCode(fasthttp.StatusCreated)
        ctx.URI().SetPath(id.Hex())
        fmt.Fprintf(ctx, "Id: %s\nRegistered: %v\n", id.Hex(), u)
    }
}

func handleLogin(sctx pbgServer.IServerContext, ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
    if un, pw, err := getPostBody(ctx.PostBody()); err != nil {
        ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
    } else if sess, err := sctx.GetAuthMechanism().DoLogin(un, pw); err != nil {
        ctx.Error(err.Error(), fasthttp.StatusBadRequest)
    } else {
        ctx.SetStatusCode(fasthttp.StatusCreated)
        fmt.Fprintf(ctx, "Session created with token %s\n", sess.GetToken())
    }
}
