package mem

import "github.com/ar3s3ru/PokemonBattleGo/pbg"

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
