package mem

import "github.com/ar3s3ru/PokemonBattleGo/pbg"

func convertLtoML(moves []move) []pbg.Move {
    if moves == nil {
        panic("Must use a valid move list!")
    }

    list := make([]pbg.Move, len(moves), len(moves))
    for i := range moves {
        list[i] = pbg.Move(&moves[i])
    }

    return list
}

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

func inRange(i int, upperBound int) bool {
    return i < 1 || i > upperBound
}

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
