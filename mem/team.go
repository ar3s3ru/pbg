package mem

import "github.com/ar3s3ru/PokemonBattleGo/pbg"

type pokèmonTeam struct {
    pbg.Pokèmon        `json:"pokèmon"`
    Moves  [4]pbg.Move `json:"moves"`
    Level  int         `json:"level"`
    IVs    [6]int      `json:"ivs"`
    EVs    [6]int      `json:"evs"`
}

func (pt *pokèmonTeam) GetMoves() [4]pbg.Move {
    return pt.Moves
}

func (pt *pokèmonTeam) GetLevel() int {
    return pt.Level
}

func (pt *pokèmonTeam) GetIVs() [6]int {
    return pt.IVs
}

func (pt *pokèmonTeam) GetEVs() [6]int {
    return pt.EVs
}
