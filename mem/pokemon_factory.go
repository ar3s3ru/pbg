package mem

import "github.com/ar3s3ru/PokemonBattleGo/pbg"

type PokèmonFactory func() pbg.Pokèmon

func NewPokèmonFactory() PokèmonFactory {
    return PokèmonFactory(func() pbg.Pokèmon {
        return &pokèmon{}
    })
}

func (pf PokèmonFactory) Create(_ ...pbg.PokèmonFactoryOption) (pbg.Pokèmon, error) {
    // Discarding options for now
    return pf(), nil
}

func (_ PokèmonFactory) CreateSlice(pokèmons ...pbg.Pokèmon) ([]pbg.Pokèmon, error) {
    slice := make([]pbg.Pokèmon, len(pokèmons), len(pokèmons))
    for i, pokèmon := range pokèmons {
        slice[i] = pokèmon
    }

    // Ignoring errors for now
    return slice, nil
}
