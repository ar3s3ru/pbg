package mem

import (
    "log"

    "github.com/ar3s3ru/PokemonBattleGo/pbg"
)

type (
    dataComponentOption func(*dataComponent)    error
    DataComponentOption func(pbg.DataComponent) error
)

func adaptDataComponentOption(option dataComponentOption) DataComponentOption {
    return func(dc pbg.DataComponent) error {
        if dcc, ok := dc.(*dataComponent); !ok {
            return pbg.ErrInvalidDataComponent
        } else {
            return option(dcc)
        }
    }
}

//func WithPokèmonDataset(pokèmons []pbg.Pokèmon) DataComponentOption {
//    return adaptDataComponentOption(func(dc *dataComponent) error {
//        dc.pokèmons = pokèmons
//        return nil
//    })
//}
//
//func WithMoveDataset(moves []pbg.Move) DataComponentOption {
//    return adaptDataComponentOption(func(dc *dataComponent) error {
//        dc.moves = moves
//        return nil
//    })
//}

func WithDatasetFile(file string) DataComponentOption {
    return adaptDataComponentOption(func(dc *dataComponent) error {
        sf, err := marshalSourceFile(file)
        if err != nil {
            return err
        }

        dc.pokèmons = convertLtoPL(sf.PList)
        dc.moves    = convertLtoML(sf.MList)

        return nil
    })
}

func WithDataLogger(logger *log.Logger) DataComponentOption {
    return adaptDataComponentOption(func (dc *dataComponent) error {
        dc.logger = logger
        return nil
    })
}
