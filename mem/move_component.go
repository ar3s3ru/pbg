package mem

import (
    "github.com/ar3s3ru/PokemonBattleGo/pbg"
    "log"
)

type (
    MoveDBComponent struct {
        moves    []pbg.Move
        logger   *log.Logger
    }
)

func NewMoveDBComponent(options ...MoveDBComponentOption) pbg.MoveDBComponent {
    movedb := &MoveDBComponent{}

    for _, option := range options {
        if err := option(movedb); err != nil {
            panic(err)
        }
    }

    return movedb
}

func (mdb *MoveDBComponent) Supply() pbg.MoveDBInterface {
    return mdb
}

func (mdb *MoveDBComponent) Retrieve(_ pbg.MoveDBInterface) {
    // Do nothing for now
}
