package main

import (
    "fmt"
    "encoding/json"

    "golang.org/x/crypto/bcrypt"

    "github.com/valyala/fasthttp"
    "github.com/ar3s3ru/PokemonBattleGo/pbg"
)

type postBody struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

const (
    loginPath       = "/login"
    registratonPath = "/signup"
)

func decodePostBody(bPostBody []byte) (user string, pass string, err error) {
    goPostBody := postBody{}
    if err = json.Unmarshal(bPostBody, &goPostBody); err != nil {
        return
    }

    user, pass = goPostBody.Username, goPostBody.Password
    return
}

func handleRegistration(ctx *fasthttp.RequestCtx) {
    user, pass, err := decodePostBody(ctx.PostBody())
    if err != nil {
        pbg.WriteAPIError(ctx, err, fasthttp.StatusInternalServerError)
        return
    }

    dataMechanism, ok := ctx.UserValue(pbg.DataAccessKey).(pbg.DataMechanism)
    if !ok {
        // Error here
        pbg.WriteAPIError(ctx, ErrInHandlerConversion, fasthttp.StatusInternalServerError)
        return
    }

    switch id, err := dataMechanism.AddTrainer(user, pass); err {
    case nil:
        pbg.WriteAPISuccess(ctx,
            fmt.Sprintf("Created at %s", id.Hex()),
            fasthttp.StatusCreated,
        )
    case pbg.ErrTrainerAlreadyExists:
        pbg.WriteAPIError(ctx, err, fasthttp.StatusBadRequest)
    case pbg.ErrPasswordSalting:
        pbg.WriteAPIError(ctx, err, fasthttp.StatusInternalServerError)
    default:
        pbg.WriteAPIError(ctx, err, fasthttp.StatusInternalServerError)
    }
}

func handleLogin(ctx *fasthttp.RequestCtx) {
    user, pass, err := decodePostBody(ctx.PostBody())
    if err != nil {
        pbg.WriteAPIError(ctx, err, fasthttp.StatusInternalServerError)
        return
    }

    dataMechanism, ok := ctx.UserValue(pbg.DataAccessKey).(pbg.DataMechanism)
    if !ok {
        // Error here
        pbg.WriteAPIError(ctx, ErrInHandlerConversion, fasthttp.StatusInternalServerError)
        return
    }

    sessionMechanism, ok := ctx.UserValue(pbg.SessionAccessKey).(pbg.SessionMechanism)
    if !ok {
        // Error here
        pbg.WriteAPIError(ctx, ErrInHandlerConversion, fasthttp.StatusInternalServerError)
        return
    }

    trainer, err := dataMechanism.GetTrainerByName(user)
    if err != nil {
        pbg.WriteAPIError(ctx, err, fasthttp.StatusBadRequest)
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(trainer.GetPasswordHash()), []byte(pass)); err != nil {
        pbg.WriteAPIError(ctx, pbg.ErrInvalidPasswordUsed, fasthttp.StatusBadRequest)
        return
    }

    if session, err := sessionMechanism.AddSession(trainer); err != nil {
        pbg.WriteAPIError(ctx, err, fasthttp.StatusInternalServerError)
    } else {
        pbg.WriteAPISuccess(ctx, session, fasthttp.StatusCreated)
    }
}
