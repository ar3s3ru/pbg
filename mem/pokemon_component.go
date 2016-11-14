package mem

import (
    "github.com/ar3s3ru/PokemonBattleGo/pbg"
    "log"
)

type (
    pokèmonDBComponent struct {
        pokèmons []pbg.Pokèmon
        logger   *log.Logger
    }
)

func NewPokèmonDBComponent(options ...pbg.PokèmonDBComponentOption) pbg.PokèmonDBComponent {
    pokedb := &pokèmonDBComponent{}

    for _, option := range options {
        if err := option(pokedb); err != nil {
            panic(err)
        }
    }

    return pokedb
}

func (pdb *pokèmonDBComponent) Supply() pbg.PokèmonDBInterface {
    return pdb
}

func (pdb *pokèmonDBComponent) Retrieve(_ pbg.PokèmonDBInterface) {
    // Do nothing for now
}
