package mem

import (
    "time"

    "github.com/ar3s3ru/PokemonBattleGo/pbg"
)

type (
    trainerFactoryOption func(*trainer) error
)

func NewTrainer(options ...pbg.TrainerFactoryOption) (pbg.Trainer, error) {
    trainer := &trainer{
        signUp:     time.Now(),
        setted:     false,
        team:       [6]pbg.PokèmonTeam{nil, nil, nil, nil, nil, nil},
        class:      pbg.TrainerC,
    }

    for _, option := range options {
        if err := option(trainer); err != nil {
            return nil, err
        }
    }

    return trainer, nil
}

func adaptTrainerFactoryOption(option trainerFactoryOption) pbg.TrainerFactoryOption {
    return func(tr pbg.Trainer) error {
        switch converted := tr.(type) {
        case *trainer:
            return option(converted)
        default:
            return ErrInvalidTrainerType
        }
    }
}

func WithTrainerName(name string) pbg.TrainerFactoryOption {
    return adaptTrainerFactoryOption(func(trainer *trainer) error {
        trainer.name = name
        return nil
    })
}

func WithTrainerPassword(pass []byte) pbg.TrainerFactoryOption {
    return adaptTrainerFactoryOption(func(trainer *trainer) error {
        trainer.hashedPass = pass
        return nil
    })
}
