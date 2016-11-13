package mem

import "github.com/ar3s3ru/PokemonBattleGo/pbg"

type pokèmonTeam struct {
    pbg.Pokèmon      `json:"pokèmon"`
    Movs [4]pbg.Move `json:"moves"`
    Levl int         `json:"level"`
    Ivs  [6]int      `json:"ivs"`
    Evs  [6]int      `json:"evs"`
}

func (pt *pokèmonTeam) Moves() [4]pbg.Move {
    return pt.Movs
}

func (pt *pokèmonTeam) Level() int {
    return pt.Levl
}

func (pt *pokèmonTeam) IVs() [6]int {
    return pt.Ivs
}

func (pt *pokèmonTeam) EVs() [6]int {
    return pt.Evs
}
