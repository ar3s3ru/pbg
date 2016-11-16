package mem

import "github.com/ar3s3ru/pbg"

type (
	pokèmon struct {
		Nam    string          `json:"name"`
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

func (pkm *pokèmon) Name() string {
	return pkm.Nam
}

func (pkm *pokèmon) Type() pbg.PokèmonType {
	return pkm.Typ
}

func (pkm *pokèmon) Index() int {
	return pkm.Pdn
}

func (pkm *pokèmon) BaseStats() [6]int {
	return pkm.Base
}

func (pkm *pokèmon) FrontSprite() string {
	return pkm.Sprite.Front
}

func (pkm *pokèmon) BackSprite() string {
	return pkm.Sprite.Back
}
