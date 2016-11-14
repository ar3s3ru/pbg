package mem

import (
    "github.com/ar3s3ru/PokemonBattleGo/pbg"
    "log"
)

type (
    PokèmonDBComponent struct {
        pokèmons []pbg.Pokèmon
        logger   *log.Logger
    }
)

func NewPokèmonDBComponent(options ...PokèmonDBComponentOption) pbg.PokèmonDBComponent {
    pokedb := &PokèmonDBComponent{}

    for _, option := range options {
        if err := option(pokedb); err != nil {
            panic(err)
        }
    }

    return pokedb
}

func (pdb *PokèmonDBComponent) Supply() pbg.PokèmonDBInterface {
    return pdb
}

func (pdb *PokèmonDBComponent) Retrieve(_ pbg.PokèmonDBInterface) {
    // Do nothing for now
}
