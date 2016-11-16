package mem

import (
	"github.com/ar3s3ru/PokemonBattleGo/pbg"
	"log"
)

type (
	TrainerDBComponentOption func(*TrainerDBComponent) error
)

func WithTrainerDBLogger(logger *log.Logger) TrainerDBComponentOption {
	return func(tdb *TrainerDBComponent) error {
		if logger == nil {
			return pbg.ErrInvalidLogger
		}

		tdb.logger = logger
		return nil
	}
}
