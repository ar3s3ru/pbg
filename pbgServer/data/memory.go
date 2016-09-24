package data

type MemoryDataMechanism struct {
    pokemonList map[int]*Pokemon
    trainerList map[int]*Trainer
}

func MemoryDataMechanismBuilder(jsonFile string) *MemoryDataMechanism {
    return &MemoryDataMechanism{
        pokemonList: make(map[int]*Pokemon),
        trainerList: make(map[int]*Trainer),
    }
}

func (mdm *MemoryDataMechanism) AddTrainer(trainer *Trainer) int {
    newId := len(mdm.trainerList) + 1
    mdm.trainerList[newId] = trainer
    return newId
}

func (mdm *MemoryDataMechanism) RemoveTrainer(id int) {
    delete(mdm.trainerList, id)
}

func (mdm *MemoryDataMechanism) GetPokemonById(id int) *Pokemon {
    return mdm.pokemonList[id]
}

func (mdm *MemoryDataMechanism) GetTrainerById(id int) *Trainer {
    return mdm.trainerList[id]
}

func (mdm *MemoryDataMechanism) GetTrainerByName(name string) *Trainer {
    for _, trainer := range mdm.trainerList {
        if trainer.Name == name {
            return trainer
        }
    }
    return nil
}
