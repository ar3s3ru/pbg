package main

import (
    "fmt"
    "time"
    "encoding/json"

    "golang.org/x/crypto/bcrypt"
    "github.com/valyala/fasthttp"

    "github.com/ar3s3ru/PokemonBattleGo/pbg"
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

func HandleRegistration(ctx *fasthttp.RequestCtx) {
    user, pass, err := DecodePostBody(ctx.PostBody())
    if err != nil {
        pbg.WriteAPIError(ctx, err, fasthttp.StatusInternalServerError)
        return
    }

    trainerDB, ok := ctx.UserValue(pbg.TrainerDBInterfaceKey).(pbg.TrainerDBInterface)
    if !ok {
        // Error here
        pbg.WriteAPIError(ctx, ErrInHandlerConversion, fasthttp.StatusInternalServerError)
        return
    }

    switch id, err := trainerDB.AddTrainer(user, pass); err {
    case nil:
        pbg.WriteAPISuccess(ctx,
            fmt.Sprintf("Created at %s", id.Hex()),
            fasthttp.StatusCreated,
        )
    case pbg.ErrTrainerAlreadyExists:
        pbg.WriteAPIError(ctx, err, fasthttp.StatusBadRequest)
    default:
        pbg.WriteAPIError(ctx, err, fasthttp.StatusInternalServerError)
    }
}

func HandleLogin(ctx *fasthttp.RequestCtx) {
    user, pass, err := DecodePostBody(ctx.PostBody())
    if err != nil {
        pbg.WriteAPIError(ctx, err, fasthttp.StatusInternalServerError)
        return
    }

    trainerDB, ok := ctx.UserValue(pbg.TrainerDBInterfaceKey).(pbg.TrainerDBInterface)
    if !ok {
        // Error here
        pbg.WriteAPIError(ctx, ErrInHandlerConversion, fasthttp.StatusInternalServerError)
        return
    }

    sessionDB, ok := ctx.UserValue(pbg.SessionDBInterfaceKey).(pbg.SessionInterface)
    if !ok {
        // Error here
        pbg.WriteAPIError(ctx, ErrInHandlerConversion, fasthttp.StatusInternalServerError)
        return
    }

    trainer, err := trainerDB.GetTrainerByName(user)
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

    if session, err := sessionDB.AddSession(
        trainer, time.Now().Add(time.Minute * 5),
    ); err != nil {
        pbg.WriteAPIError(ctx, err, fasthttp.StatusInternalServerError)
    } else {
        pbg.WriteAPISuccess(ctx, session, fasthttp.StatusCreated)
    }
}
