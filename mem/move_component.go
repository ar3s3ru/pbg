package mem

import (
    "github.com/ar3s3ru/PokemonBattleGo/pbg"
    "log"
)

type (
    moveDBComponent struct {
        moves    []pbg.Move
        logger   *log.Logger
    }
)

func NewMoveDBComponent(options ...pbg.MoveDBComponentOption) pbg.MoveDBComponent {
    movedb := &moveDBComponent{}

    for _, option := range options {
        if err := option(movedb); err != nil {
            panic(err)
        }
    }

    return movedb
}

func (mdb *moveDBComponent) Supply() pbg.MoveDBInterface {
    return mdb
}

func (mdb *moveDBComponent) Retrieve(_ pbg.MoveDBInterface) {
    // Do nothing for now
}
