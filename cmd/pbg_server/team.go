package main

import (
	"encoding/json"
	"errors"

	"github.com/valyala/fasthttp"

	"github.com/ar3s3ru/pbg"
	"github.com/ar3s3ru/pbg/mem"
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
	SetupPath = MePath + "/setup"
)

var (
	ErrTrainerAlreadySetUp = errors.New("Trainer is already setted up, use /me/update instead")
	ErrInvalidPOSTBody     = errors.New("There's been some error with your POST body, please check it out")
)

func handleSettingTeamUp(ctx *fasthttp.RequestCtx) {
	session, ok := ctx.UserValue(pbg.SessionKey).(pbg.Session)
	if !ok {
		pbg.WriteAPIError(ctx, ErrInHandlerConversion, fasthttp.StatusInternalServerError)
		return
	}

	pokèmonDB, ok := ctx.UserValue(pbg.PokèmonDBInterfaceKey).(pbg.PokèmonDBInterface)
	if !ok {
		pbg.WriteAPIError(ctx, ErrInHandlerConversion, fasthttp.StatusInternalServerError)
		return
	}

	moveDB, ok := ctx.UserValue(pbg.MoveDBInterfaceKey).(pbg.MoveDBInterface)
	if !ok {
		pbg.WriteAPIError(ctx, ErrInHandlerConversion, fasthttp.StatusInternalServerError)
		return
	}

	user := session.Reference()
	if user.Set() {
		pbg.WriteAPIError(ctx, ErrTrainerAlreadySetUp, fasthttp.StatusConflict)
		return
	}

	body := setupBody{}
	if err := json.Unmarshal(ctx.PostBody(), &body); err != nil {
		pbg.WriteAPIError(ctx, ErrInvalidPOSTBody, fasthttp.StatusBadRequest)
		return
	}

	team := [6]pbg.PokèmonTeam{nil, nil, nil, nil, nil, nil}
	for i, pokèmonBody := range body {
		if pokèmonBody == nil {
			continue
		}

		ctx.Logger().Printf("Pokèmon body req is: %v\n", pokèmonBody)
		pokèmon, err := pokèmonDB.GetPokèmon(pokèmonBody.Pkmn)
		if err != nil {
			pbg.WriteAPIError(ctx, err, fasthttp.StatusBadRequest)
			return
		}

		moves := [4]pbg.Move{nil, nil, nil, nil}
		for i, moveId := range pokèmonBody.Moves {
			move, err := moveDB.GetMove(moveId)
			if err != nil {
				pbg.WriteAPIError(ctx, err, fasthttp.StatusBadRequest)
				return
			}

			moves[i] = move
		}

		if pokèmonTeam, err := mem.NewPokèmonTeam(
			mem.WithPokèmonReference(pokèmon),
			mem.WithPokèmonMoves(moves[0], moves[1], moves[2], moves[3]),
			mem.WithPokèmonLevel(pokèmonBody.Level),
			mem.WithPokèmonIVs(pokèmonBody.IVs),
			mem.WithPokèmonEVs(pokèmonBody.EVs),
		); err == mem.ErrInvalidReferenceValue {
			pbg.WriteAPIError(ctx, err, fasthttp.StatusInternalServerError)
		} else if err != nil {
			pbg.WriteAPIError(ctx, err, fasthttp.StatusBadRequest)
		} else {
			team[i] = pokèmonTeam
		}
	}

	if err := user.SetTrainer(team, pbg.TrainerC); err != nil {
		pbg.WriteAPIError(ctx, err, fasthttp.StatusBadRequest)
	} else {
		pbg.WriteAPISuccess(ctx, user, fasthttp.StatusOK)
	}
}
