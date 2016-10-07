package main

import (
    "github.com/ar3s3ru/PokemonBattleGo/pbgServer"
    "github.com/valyala/fasthttp"
    "github.com/buaazp/fasthttprouter"
    "errors"
    "encoding/json"
)

type (
    teamEntry struct {
        Pkmn  int    `json:"pokemon"`
        Level int    `json:"level"`
        Moves [4]int `json:"moves"`
        IVs   [6]int `json:"ivs"`
        EVs   [6]int `json:"evs"`
    }

    setupBody [6]*teamEntry
)

const (
    APIMe    = APIEndpoint + "/me"
    APISetUp = APIMe + "/setup"
)

var (
    ErrAlreadySetUp = errors.New("Trainer is already set up, use /me/update instead")
)

func handleMe(sess pbgServer.Session, _ pbgServer.IServerContext,
               ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) (int, interface{}, error) {
    return fasthttp.StatusOK, sess.GetUserReference(), nil
}

func handleSettingUp(sess pbgServer.Session, sctx pbgServer.IServerContext,
                      ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) (int, interface{}, error) {
    if sess.GetUserReference().IsSet() {
        return fasthttp.StatusBadRequest, nil, ErrAlreadySetUp
    }

    body := setupBody{}
    if err := json.Unmarshal(ctx.PostBody(), &body); err != nil {
        return fasthttp.StatusBadRequest, nil, err
    }

    team := sess.GetUserReference().GetTeam()
    for i, v := range body {
        if v == nil {
            // No value, we don't want to use it
            if i != 0 {
                continue
            } else {
                return fasthttp.StatusBadRequest, nil, pbgServer.ErrInvalidFirstPokemon
            }
        }

        if pk, err := sctx.GetDataMechanism().AddPokèmonTeam(v.Pkmn, v.Moves, v.Level, v.IVs, v.EVs); err != nil {
            return fasthttp.StatusInternalServerError, nil, err
        } else {
            team[i] = pk
        }
    }

    if err := sess.GetUserReference().SetTrainer(team, pbgServer.TrainerC); err != nil {
        return fasthttp.StatusInternalServerError, nil, err
    }

    return fasthttp.StatusCreated, sess.GetUserReference(), nil
}
