package pbgServer

import "fmt"

type (
    // Interfaccia che rappresenta un Pokèmon.
    //
    // Interface that represent a Pokèmon.
    Pokèmon interface {
        GetName()      string
        GetType()      PokèmonType
        GetPokèdex()   int
        GetBaseStats() [6]int
        // Sprites
        GetFrontSprite() string
        GetBackSprite()  string
    }

    // Interfaccia che rappresenta una mossa Pokèmon.
    //
    // Interface that represents a Pokèmon move.
    Move interface {
        GetName()     string
        GetType()     Type
        GetCategory() Category
        GetPriority() int
        GetPower()    int
        GetAccuracy() int
        GetPPs()      int
    }

    // Interfaccia che rappresenta un Allenatore Pokèmon.
    // L'ALlenatore è un modello dell'utente che interagisce col server - indi per cui ha una relazione di composizione
    // con l'interfaccia Utente.
    //
    // Interface that represent a Pokèmon Trainer.
    // The Trainer is a model for the user that interacts with the server - as such, it has a composition relationship
    // with the User interface.
    Trainer interface {
        User
        IsSet()   bool
        GetType() TrainerClass
        GetTeam() [6]PokèmonTeam
    }

    // Interfaccia che rappresenta un Pokèmon all'interno di una squadra.
    // Ha una relazione di composizione con un oggetto Pokèmon, ma estende un Pokèmon base con possibile Nickname,
    // un numero di mosse (da 1 a 4 possibili), un certo livello, una determinata Natura e Abilità, e un numero di IVs
    // ed EVs.
    //
    // Interface that represent a Pokèmon inside a team.
    // It has a composition relationship with a Pokèmon object, but extends the latter with a possible Nickname,
    // a certain number of moves (from 1 to 4), a certain level, a specified Nature and Ability, IVs and EVs.
    PokèmonTeam interface {
        Pokèmon
        GetMoves()   [4]Move
        GetLevel()   int
        GetNature()  Nature
        GetAbility() Ability

        GetIVs() [6]int
        GetIV()  int

        GetEVs() [6]int
        GetEV()  int
    }

    // Interfaccia che rappresenta una Natura.
    //
    // Interface that represent a Nature.
    Nature interface {
        GetName() string
        // TODO
    }

    // Interfaccia che rappresenta un'Abilità.
    //
    // Interface that represent an Ability.
    Ability interface {
        GetName() string
        // TODO
    }

    // Categoria di mossa Pokèmon.
    //
    // Pokèmon move Category.
    Category int

    // Tipo di mossa Pokèmon.
    //
    // Pokèmon move Type.
    Type int

    // Classe di un ALlenatore Pokèmon.
    //
    // Pokèmon Trainer Class.
    TrainerClass int

    // Tipo di un Pokèmon; può avere fino a 2 tipi di mossa Pokèmon.
    //
    // Pokèmon Type; can have up to 2 Pokèmon move type.
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
    Fighting
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
    // Nomi per le varie categorie
    CategoryNames = [...]string {
        "Physique", "Special", "State",
    }

    // Nomi per i vari tipi
    TypeNames = [...]string {
        "Normal", "Fire", "Fightning", "Water", "Flying", "Grass", "Poison", "Electric", "Ground", "Psychic", "Rock",
        "Ice", "Bug", "Dragon", "Ghost", "Dark", "Steel", "Fairy", "???",
    }

    // Nomi per le varie classi di allenatori
    ClassesNames = [...]string {
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
