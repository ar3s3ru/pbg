package mem

import (
    "github.com/ar3s3ru/PokemonBattleGo/pbg"
    "errors"
    "log"
)

type (
    moveDBComponentOption func(*moveDBComponent) error
)

var (
    ErrInvalidMoveDBComponent = errors.New("Invalid MoveDBComponent used, must be from mem package")
)

func adaptMoveDBComponentOption(option moveDBComponentOption) pbg.MoveDBComponentOption {
    return func(movedb pbg.MoveDBComponent) error {
        if covertedMoveDB, ok := movedb.(*moveDBComponent); !ok {
            return ErrInvalidMoveDBComponent
        } else {
            return option(covertedMoveDB)
        }
    }
}

func WithMovesFile(moveFile string) pbg.MoveDBComponentOption {
    return adaptMoveDBComponentOption(func(mdb *moveDBComponent) error {
        // TODO: finish this
        return nil
    })
}

func WithMoves(moves []pbg.Move) pbg.MoveDBComponentOption {
    return adaptMoveDBComponentOption(func(mdb *moveDBComponent) error {
        if moves == nil {
            return ErrInvalidMoveDataset
        }

        mdb.moves = moves
        return nil
    })
}

func WithMoveDBLogger(logger *log.Logger) pbg.MoveDBComponentOption {
    return adaptMoveDBComponentOption(func(mdb *moveDBComponent) error {
        if logger == nil {
            return pbg.ErrInvalidLogger
        }

        mdb.logger = logger
        return nil
    })
}