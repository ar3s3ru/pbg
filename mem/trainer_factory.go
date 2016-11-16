package mem

import (
	"golang.org/x/crypto/bcrypt"
	"time"

	"github.com/ar3s3ru/pbg"
)

type (
	TrainerFactory       func(...TrainerFactoryOption) (pbg.Trainer, error)
	TrainerFactoryOption func(*Trainer) error
)

func NewTrainer(options ...TrainerFactoryOption) (pbg.Trainer, error) {
	trainer := &Trainer{
		signUp: time.Now(),
		setted: false,
		team:   [6]pbg.PokèmonTeam{nil, nil, nil, nil, nil, nil},
		class:  pbg.TrainerC,
	}

	for _, option := range options {
		if err := option(trainer); err != nil {
			return nil, err
		}
	}

	return trainer, nil
}

func WithTrainerName(name string) TrainerFactoryOption {
	return func(trainer *Trainer) error {
		trainer.name = name
		return nil
	}
}

func WithTrainerPassword(pass string) TrainerFactoryOption {
	return func(trainer *Trainer) error {
		pwd, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
		if err != nil {
			return pbg.ErrPasswordSalting
		}

		trainer.hashedPass = pwd
		return nil
	}
}
