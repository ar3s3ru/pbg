package mem

import (
    "gopkg.in/mgo.v2/bson"
    "github.com/ar3s3ru/PokemonBattleGo/pbg"
    "log"
)

type (
    trainerRequest func(map[bson.ObjectId]pbg.Trainer)
    trainerDBComponent struct {
        trainers map[bson.ObjectId]pbg.Trainer
        requests chan trainerRequest
        factory  pbg.TrainerFactory
        logger   *log.Logger
    }
)

func NewTrainerDBComponent(options ...pbg.TrainerDBComponentOption) pbg.TrainerDBComponent {
    trainerdb := &trainerDBComponent{
        trainers: make(map[bson.ObjectId]pbg.Trainer),
        requests: make(chan trainerRequest),
        factory:  NewTrainer,
    }

    for _, option := range options {
        if err := option(trainerdb); err != nil {
            panic(err)
        }
    }

    go trainerdb.requestDispatcher()
    return trainerdb
}

func (tdb *trainerDBComponent) requestDispatcher() {
    for request := range tdb.requests {
        request(tdb.trainers)
    }
}

func (tdb *trainerDBComponent) Supply() pbg.TrainerDBInterface {
    return tdb
}

func (tdb *trainerDBComponent) Retrieve(_ pbg.TrainerDBInterface) {
    // Nothing for now
}


