package mem

import (
    "errors"
    "github.com/ar3s3ru/PokemonBattleGo/pbg"
    "log"
)

type (
    pokèmonDBComponentOption func(*pokèmonDBComponent) error
)

var (
    ErrInvalidPokèmonDBComponent = errors.New("Invalid PokèmonDBComponent used, must be from mem package")
)

func adaptPokèmonDBComponentOption(option pokèmonDBComponentOption) pbg.PokèmonDBComponentOption {
    return func(pokedb pbg.PokèmonDBComponent) error {
        if convertedPokeDB, ok := pokedb.(*pokèmonDBComponent); !ok {
            return ErrInvalidPokèmonDBComponent
        } else {
            return option(convertedPokeDB)
        }
    }
}

func WithPokèmons(pokèmons []pbg.Pokèmon) pbg.PokèmonDBComponentOption {
    return adaptPokèmonDBComponentOption(func (pdb *pokèmonDBComponent) error {
        if pokèmons == nil {
            return ErrInvalidPokèmonDataset
        }

        pdb.pokèmons = pokèmons
        return nil
    })
}

func WithPokèmonDBLogger(logger *log.Logger) pbg.PokèmonDBComponentOption {
    return adaptPokèmonDBComponentOption(func(pdb *pokèmonDBComponent) error {
        if logger == nil {
            return pbg.ErrInvalidLogger
        }

        pdb.logger = logger
        return nil
    })
}