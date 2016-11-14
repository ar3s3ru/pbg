package mem

import "github.com/ar3s3ru/PokemonBattleGo/pbg"

func (pdb *pokèmonDBComponent) Log(v ...interface{}) {
    if pdb.logger != nil {
        pdb.logger.Println(v...)
    }
}

func (pdb *pokèmonDBComponent) GetPokèmon(id int) (pbg.Pokèmon, error) {
    if inRange(id, len(pdb.pokèmons)) {
        return pdb.pokèmons[id - 1], nil
    } else {
        return nil, pbg.ErrPokèmonNotFound
    }
}

func (pdb *pokèmonDBComponent) GetPokèmons() []pbg.Pokèmon {
    return pdb.pokèmons
}
