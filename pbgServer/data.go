package pbgServer

type IDataMechanism interface {
    AddTrainer(trainer Trainer) int
    RemoveTrainer(id int)

    GetPokemonById(id int)        Pokemon
    GetTrainerById(id int)        Trainer
    GetTrainerByName(name string) Trainer
}
