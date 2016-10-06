package main

import (
    "github.com/ar3s3ru/PokemonBattleGo/pbgServer"
    "github.com/valyala/fasthttp"
    "github.com/buaazp/fasthttprouter"
    "encoding/json"
)

type logRegPostBody struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

const (
    APIRegister  = APIEndpoint + "/register"
    APILogin     = APIEndpoint + "/login"
    APIMe        = APIEndpoint + "/me"
)

func getPostBody(postBody []byte) (user string, pass string, err error) {
    req := logRegPostBody{}
    if err = json.Unmarshal(postBody, &req); err != nil {
        return
    }

    user, pass, err = req.Username, req.Password, nil
    return
}

func handleRegister(sctx pbgServer.IServerContext,
                     ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) (int, interface{}, error) {
    if un, pw, err := getPostBody(ctx.PostBody()); err != nil {
        return fasthttp.StatusInternalServerError, nil, err
    } else if _, _, err := sctx.GetAuthMechanism().Register(un, pw); err != nil {
        return fasthttp.StatusBadRequest, nil, err
    } else {
        return fasthttp.StatusCreated, "Registered", nil
    }
}

func handleLogin(sctx pbgServer.IServerContext,
                 ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) (int, interface{}, error) {
    if un, pw, err := getPostBody(ctx.PostBody()); err != nil {
        return fasthttp.StatusInternalServerError, nil, err
    } else if sess, err := sctx.GetAuthMechanism().DoLogin(un, pw); err != nil {
        return fasthttp.StatusBadRequest, nil, err
    } else {
        return fasthttp.StatusCreated, sess, nil
    }
}

func handleMe(sess pbgServer.Session, _ pbgServer.IServerContext,
               ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) (int, interface{}, error) {
    return fasthttp.StatusOK, sess.GetUserReference(), nil
}
