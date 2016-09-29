package pbgServer

import "fmt"

type (
    Pokèmon interface {
        GetName()      string
        GetType()      PokèmonType
        GetPokèdex()   int
        GetBaseStats() [6]int
    }

    Move interface {
        GetName()      string
        GetType()      Type
        GetCategory()  Category
        GetDamage()    int
        GetPrecision() int
        GetPPs()       int
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
)

const (
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
)

const (
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

var (
    CategoryNames = [...]string {
        "Physique", "Special", "State",
    }
    TypeNames     = [...]string {
        "Normal", "Fire", "Fightning", "Water", "Flying", "Grass", "Poison", "Electric", "Ground", "Psychic", "Rock",
        "Ice", "Bug", "Dragon", "Ghost", "Dark", "Steel", "Fairy", "???",
    }
    ClassesNames  = [...]string {
        "Trainer", "Beauty", "Biker", "BirdKeeper", "Blackbelt", "Bug Catcher", "Burglar", "Channeler", "Cooltrainer",
        "Cue Ball", "Engineer", "Fisherman", "Gambler", "Gentleman", "Hiker", "Trainer Jr.", "Juggler", "Lass",
        "PokèManiac", "Psychic", "Rocker", "Rocket", "Sailor", "Scientist", "Super Nerd", "Swimmer", "Tamer",
        "Youngster", "Chief",
    }
)

func (c Category) String() string {
    if c == Physique || c == Special || c == State {
        return CategoryNames[c]
    } else {
        return "Undefined"
    }
}

func (t Type) String() string {
    if t >= Normal && t <= Undefined {
        return TypeNames[t]
    } else {
        return "Undefined"
    }
}

func (tc TrainerClass) String() string {
    if tc >= TrainerC && tc <= Chief {
        return ClassesNames[tc]
    } else {
        return "Undefined"
    }
}

func (pt PokèmonType) String() string {
    if pt[1] == -1 {
        return pt[0].String()
    } else {
        return fmt.Sprintf("%s/%s", pt[0].String(), pt[1].String())
    }
}
