package main

import (
    "fmt"
    "encoding/json"

    "golang.org/x/crypto/bcrypt"

    "github.com/valyala/fasthttp"
    "github.com/ar3s3ru/PokemonBattleGo/pbg"
    "log"
)

type PostBody struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

const (
    LoginPath       = "/login"
    RegistratonPath = "/signup"
)

func EncodePostBody(postBody *PostBody) ([]byte, error) {
    return json.Marshal(postBody)
}

func DecodePostBody(bPostBody []byte) (user string, pass string, err error) {
    goPostBody := PostBody{}
    if err = json.Unmarshal(bPostBody, &goPostBody); err != nil {
        return
    }

    user, pass = goPostBody.Username, goPostBody.Password
    return
}

func handleRegistration(ctx *fasthttp.RequestCtx) {
    user, pass, err := DecodePostBody(ctx.PostBody())
    if err != nil {
        pbg.WriteAPIError(ctx, err, fasthttp.StatusInternalServerError)
        return
    }

    di, ok := ctx.UserValue(pbg.DataInterfaceKey).(pbg.DataInterface)
    if !ok {
        // Error here
        pbg.WriteAPIError(ctx, ErrInHandlerConversion, fasthttp.StatusInternalServerError)
        return
    }

    logger := ctx.UserValue(pbg.LoggerKey).(*log.Logger)
    logger.Println("Adding trainer")

    switch id, err := di.AddTrainer(user, pass); err {
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
    user, pass, err := DecodePostBody(ctx.PostBody())
    if err != nil {
        pbg.WriteAPIError(ctx, err, fasthttp.StatusInternalServerError)
        return
    }

    dataMechanism, ok := ctx.UserValue(pbg.DataInterfaceKey).(pbg.DataInterface)
    if !ok {
        // Error here
        pbg.WriteAPIError(ctx, ErrInHandlerConversion, fasthttp.StatusInternalServerError)
        return
    }

    sessionMechanism, ok := ctx.UserValue(pbg.SessionInterfaceKey).(pbg.SessionInterface)
    if !ok {
        // Error here
        pbg.WriteAPIError(ctx, ErrInHandlerConversion, fasthttp.StatusInternalServerError)
        return
    }

    trainer, err := dataMechanism.GetTrainerByName(user)
    if err == pbg.ErrTrainerNotFound {
        pbg.WriteAPIError(ctx, err, fasthttp.StatusNotFound)
        return
    } else if err != nil {
        pbg.WriteAPIError(ctx, err, fasthttp.StatusBadRequest)
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(trainer.PasswordHash()), []byte(pass)); err != nil {
        pbg.WriteAPIError(ctx, pbg.ErrInvalidPasswordUsed, fasthttp.StatusBadRequest)
        return
    }

    if session, err := sessionMechanism.AddSession(trainer); err != nil {
        pbg.WriteAPIError(ctx, err, fasthttp.StatusInternalServerError)
    } else {
        pbg.WriteAPISuccess(ctx, session, fasthttp.StatusCreated)
    }
}
