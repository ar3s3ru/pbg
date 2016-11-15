package mem

import (
    "log"
    "github.com/ar3s3ru/PokemonBattleGo/pbg"
)

type (
    PokèmonDBComponentOption func(*PokèmonDBComponent) error
)

func WithPokèmons(pokèmons []pbg.Pokèmon) PokèmonDBComponentOption {
    return func (pdb *PokèmonDBComponent) error {
        if pokèmons == nil {
            return ErrInvalidPokèmonDataset
        }

        pdb.pokèmons = pokèmons
        return nil
    }
}

func WithPokèmonDBLogger(logger *log.Logger) PokèmonDBComponentOption {
    return func(pdb *PokèmonDBComponent) error {
        pdb.logger = logger
        return nil
    }
}