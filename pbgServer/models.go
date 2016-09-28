package pbgServer

type (
    Pokèmon interface {
        GetName()      string
        GetType()      PokèmonType
        GetPokèdex()   int
        GetBaseStats() [6]int
    }

    Move interface {
        GetName()      string
        GetType() Type
        GetCategory()  bool
        GetDamage()    int
        GetPrecision() int
    }

    Trainer interface {
        User
        GetType() TrainerClass
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

    Category     int
    Type         int
    TrainerClass int
    PokèmonType  [2]Type
)

const (
    Physique Category = iota
    Special
    State

    Normal Type = iota
    Fire
    Fightning
    Water
    Flying
    Grass
    Poison
    Electric
    Ground
    PsychicT
    Rock
    Ice
    Bug
    Dragon
    Ghost
    Dark
    Steel
    Fairy
    Undefined

    TrainerC TrainerClass = iota
    Beauty
    Biker
    BirdKeeper
    Blackbelt
    BugCatcher
    Burglar
    Channeler
    Cooltrainer
    CueBall
    Engineer
    Fisherman
    Gambler
    Gentleman
    Hiker
    JrTrainer
    Juggler
    Lass
    PokèManiac
    PsychicC
    Rocker
    Rocket
    Sailor
    Scientist
    SuperNerd
    Swimmer
    Tamer
    Youngster
    Chief
)