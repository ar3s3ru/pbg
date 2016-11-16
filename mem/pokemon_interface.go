package mem

import "github.com/ar3s3ru/PokemonBattleGo/pbg"

func (pdb *PokèmonDBComponent) Log(v ...interface{}) {
	if pdb.logger != nil {
		pdb.logger.Println(v...)
	}
}

func (pdb *PokèmonDBComponent) GetPokèmon(id int) (pbg.Pokèmon, error) {
	if inRange(id, len(pdb.pokèmons)) {
		return pdb.pokèmons[id-1], nil
	} else {
		return nil, pbg.ErrPokèmonNotFound
	}
}

func (pdb *PokèmonDBComponent) GetPokèmons() []pbg.Pokèmon {
	return pdb.pokèmons
}
