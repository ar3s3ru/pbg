package mem

import "github.com/ar3s3ru/PokemonBattleGo/pbg"

type MoveFactory func() pbg.Move

func NewMoveFactory() MoveFactory {
    return MoveFactory(func() pbg.Move {
        return &move{}
    })
}

func (mf MoveFactory) Create(_ ...pbg.MoveFactoryOption) (pbg.Move, error) {
    // Discarding options for now
    return mf(), nil
}

func (_ MoveFactory) CreateSlice(moves ...pbg.Move) ([]pbg.Move, error) {
    // TODO: finish this
    return nil, nil
}
