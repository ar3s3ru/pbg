package pbgServer

type (
    Pokèmon interface {
        GetName()      string
        GetType()      PokèmonType
        GetPokèdex()   int
        GetBaseStats() [6]int
    }

    Move interface {
        GetName() string
        GetType() MoveType
    }

    Trainer interface {
        User
        GetType() TrainerType
        GetTeam() [6]PokèmonTeam
    }

    PokèmonTeam interface {
        GetNickname() string
        GetPokemon()  Pokèmon
        GetMoves()    [4]Move
        GetLevel()    int
        GetNature()   Nature
        GetAbility()  Ability

        GetIVs() [6]int
        GetIV()  int

        GetEVs() [6]int
        GetEV()  int
    }

    Nature interface {
        GetName() string
        // TODO
    }

    Ability interface {
        GetName() string
        // TODO
    }

    MoveType    int
    TrainerType int
    PokèmonType [2]MoveType
)
