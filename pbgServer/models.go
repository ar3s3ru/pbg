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

func (c Category) String() string {
    switch c {
    case Physique:
        return "Physique"
    case Special:
        return "Special"
    case State:
        return "State"
    default:
        return "Undefined"
    }
}

func (t Type) String() string {
    switch t {
    case Normal:
        return "Normal"
    case Fire:
        return "Fire"
    case Fightning:
        return "Fightning"
    case Water:
        return "Water"
    case Flying:
        return "Flying"
    case Grass:
        return "Grass"
    case Poison:
        return "Poison"
    case Electric:
        return "Electric"
    case Ground:
        return "Ground"
    case PsychicT:
        return "Psychic"
    case Rock:
        return "Rock"
    case Ice:
        return "Ice"
    case Bug:
        return "Bug"
    case Dragon:
        return "Dragon"
    case Ghost:
        return "Ghost"
    case Dark:
        return "Dark"
    case Steel:
        return "Steel"
    case Fairy:
        return "Fairy"
    case Undefined:
        return "???"
    default:
        return "Undefined"
    }
}

func (tc TrainerClass) String() string {
    switch tc {
    case TrainerC:
        return "Trainer"
    case Beauty:
        return "Beauty"
    case Biker:
        return "Biker"
    case BirdKeeper:
        return "Bird Keeper"
    case Blackbelt:
        return "Blackbelt"
    case BugCatcher:
        return "Bug Catcher"
    case Burglar:
        return "Burglar"
    case Channeler:
        return "Channeler"
    case Cooltrainer:
        return "Cooltrainer"
    case CueBall:
        return "Cue Ball"
    case Engineer:
        return "Engineer"
    case Fisherman:
        return "Fisherman"
    case Gambler:
        return "Gambler"
    case Gentleman:
        return "Gentleman"
    case Hiker:
        return "Hiker"
    case JrTrainer:
        return "Trainer Jr."
    case Juggler:
        return "Juggler"
    case Lass:
        return "Lass"
    case PokèManiac:
        return "PokèManiac"
    case PsychicC:
        return "Psychic"
    case Rocker:
        return "Rocker"
    case Rocket:
        return "Team Rocket"
    case Sailor:
        return "Sailor"
    case Scientist:
        return "Scientist"
    case SuperNerd:
        return "Super Nerd"
    case Swimmer:
        return "Swimmer"
    case Tamer:
        return "Tamer"
    case Youngster:
        return "Youngster"
    case Chief:
        return "Chief"
    default:
        return "Undefined"
    }
}

func (pt PokèmonType) String() string {
    //if fType >= Normal && fType <= Undefined {
    //    str += fType.String()
    //}
    //
    //if sType >= Normal && sType <= Undefined {
    //    str += "/" + sType.String()
    //}
    return fmt.Sprintf("%s/%s", pt[0].String(), pt[1].String())
}
