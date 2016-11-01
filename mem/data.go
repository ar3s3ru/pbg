package mem

import (
    "sync"
    
    "gopkg.in/mgo.v2/bson"
    "github.com/ar3s3ru/PokemonBattleGo/pbg"
)

type (
    DataBuilder interface {
        WithSourceFile(string)   DataBuilder
        WithTrainersFile(string) DataBuilder

        Build() pbg.DataMechanism
    }

    dataBuilder func(string, string) pbg.DataMechanism

    data struct {
        pokèmons []pbg.Pokèmon
        moves    []pbg.Move

        trainers map[bson.ObjectId]pbg.Trainer
        mutex    sync.RWMutex
    }

    sourceFile struct {
        Generation int       `json:"generation"`
        PNumbers   int       `json:"pokemon_count"`
        MNumbers   int       `json:"move_count"`
        //PList      []pokèmon `json:"pokemons"`
        //MList      []move    `json:"moves"`
    }
)

func NewDataBuilder() DataBuilder {
    return func(source string, trainers string) pbg.DataMechanism {
        return &data{}
    }
}

func (db dataBuilder) WithSourceFile(path string) DataBuilder {
    return func(_ string, trainers string) pbg.DataMechanism {
        return db(path, trainers)
    }
}

func (db dataBuilder) WithTrainersFile(path string) DataBuilder {
    return func(source string, _ string) pbg.DataMechanism {
        return db(source, path)
    }
}

func (db dataBuilder) Build() pbg.DataMechanism {
    return db("", "")
}
