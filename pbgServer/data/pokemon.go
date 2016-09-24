package data

type (
    Statistic uint8

    Move struct {
        Name string
        Type PokemonType
    }

    Pokemon struct {
        Name     string
        Type     [2]PokemonType
        BaseStat [6]uint8
    }
)

const (
    PS Statistic = iota
    Atk
    Def
    AtkSp
    DefSp
    Vel
)