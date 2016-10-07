package mem

import pbg "github.com/ar3s3ru/PokemonBattleGo/pbgServer"

type (
    pokèmonTeam struct {
        pbg.Pokèmon        `json:"pokèmon"`
        Moves  [4]pbg.Move `json:"moves"`
        Level  int         `json:"level"`
        IVs    [6]int      `json:"ivs"`
        EVs    [6]int      `json:"evs"`
    }
)

func checkIvs(ivs [6]int) bool {
    sum := 0
    for _, v := range ivs {
        if v < 0 || v > 31 {
            return false
        }

        sum += v
    }

    if sum < 0 || sum > (6 * 31) {
        return false
    }

    return true
}

func checkEvs(evs [6]int) bool {
    sum := 0
    for _, v := range evs {
        if v < 0 || v > 255 {
            return false
        }

        sum += v
    }

    if sum < 0 || sum > 510 {
        return false
    }

    return true
}

func checkLevel(level int) bool {
    return level >= 1 && level <= 100
}

func (md *memData) AddPokèmonTeam(pkmnId int, moves [4]int, level int, ivs [6]int, evs [6]int) (pbg.PokèmonTeam, error) {

    if !checkIvs(ivs) {
        return nil, pbg.ErrInvalidPokemonIVs
    }

    if !checkEvs(evs) {
        return nil, pbg.ErrInvalidPokemonEVs
    }

    if !checkLevel(level) {
        return nil, pbg.ErrInvalidPokemonLevel
    }

    if pkmn, err := md.GetPokèmonById(pkmnId); err != nil {
        return nil, pbg.ErrInvalidPokemonId
    } else {
        // Check moves
        var iMoves [4]pbg.Move
        for i, v := range moves {
            if move, err := md.GetMoveById(v); err != nil {
                // Invalid move
                return nil, pbg.ErrInvalidMoveId
            } else {
                // Valid move
                iMoves[i] = move
            }
        }

        return &pokèmonTeam{
            Pokèmon: pkmn,
            Moves:   iMoves,
            Level:   level,
            IVs:     ivs,
            EVs:     evs,
        }, nil
    }
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
