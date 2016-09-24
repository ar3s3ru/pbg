package data

type (
    IDataMechanism interface {
        AddTrainer(trainer *Trainer) int
        RemoveTrainer(id int)

        GetPokemonById(id int)        *Pokemon
        GetTrainerById(id int)        *Trainer
        GetTrainerByName(name string) *Trainer
    }

    TrainerType uint8
    PokemonType uint8
)

const (
    // TrainerType
    Scholar TrainerType = iota

    // PokemonType
    Fire PokemonType = iota
)