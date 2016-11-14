package mem

import (
    "errors"
    "github.com/ar3s3ru/PokemonBattleGo/pbg"
    "log"
)

type (
    trainerDBComponentOption func(*trainerDBComponent) error
)

var (
    ErrInvalidTrainerDBComponent = errors.New("Invalid TrainerDBComponent used, must be from mem package")
)

func adaptTrainerDBComponentOption(option trainerDBComponentOption) pbg.TrainerDBComponentOption {
    return func(trainerdb pbg.TrainerDBComponent) error {
        if convertedTrainerDB, ok := trainerdb.(*trainerDBComponent); !ok {
            return ErrInvalidTrainerDBComponent
        } else {
            return option(convertedTrainerDB)
        }
    }
}

func WithTrainerDBLogger(logger *log.Logger) pbg.TrainerDBComponentOption {
    return adaptTrainerDBComponentOption(func(tdb *trainerDBComponent) error {
        if logger == nil {
            return pbg.ErrInvalidLogger
        }

        tdb.logger = logger
        return nil
    })
}
