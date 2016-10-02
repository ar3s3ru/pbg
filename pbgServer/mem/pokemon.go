package mem

import (
    pbg "github.com/ar3s3ru/PokemonBattleGo/pbgServer"
)

type (
    pokèmon struct {
        Name   string          `json:"name"`
        Typ    pbg.PokèmonType `json:"type"`
        Pdn    int             `json:"pokedex"`
        Base   [6]int          `json:"baseStats"`
        Sprite sprite          `json:"sprites"`
    }

    sprite struct {
        Front string `json:"front"`
        Back  string `json:"back"`
    }

    nature struct {
        name string `json:"name"`
    }

    ability struct {
        name string `json:"name"`
    }
)

func convertLtoPL(pkdx []pokèmon) []pbg.Pokèmon {
    if pkdx == nil {
        panic("Must use a valid pokèmon list!")
    }

    list := make([]pbg.Pokèmon, len(pkdx), len(pkdx))
    for i := range pkdx {
        list[i] = pbg.Pokèmon(&pkdx[i])
    }

    return list
}

func (pkm *pokèmon) GetName() string {
    return pkm.Name
}

func (pkm *pokèmon) GetType() pbg.PokèmonType {
    return pkm.Typ
}

func (pkm *pokèmon) GetPokèdex() int {
    return pkm.Pdn
}

func (pkm *pokèmon) GetBaseStats() [6]int {
    return pkm.Base
}

func (pkm *pokèmon) GetFrontSprite() string {
    return pkm.Sprite.Front
}

func (pkm *pokèmon) GetBackSprite() string {
    return pkm.Sprite.Back
}
