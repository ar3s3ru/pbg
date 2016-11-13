package mem

import (
    "log"
    //"runtime"

    "gopkg.in/mgo.v2/bson"
    "github.com/ar3s3ru/PokemonBattleGo/pbg"
)

type (
    pokèmonReq func([]pbg.Pokèmon)
    moveReq    func([]pbg.Move)
    trainerReq func(map[bson.ObjectId]pbg.Trainer)

    dataComponent struct {
        moves    []pbg.Move
        pokèmons []pbg.Pokèmon
        trainers map[bson.ObjectId]pbg.Trainer

        trainerFactory pbg.TrainerFactory

        moveReqs    chan moveReq
        pokèmonReqs chan pokèmonReq
        trainerReqs chan trainerReq

        logger *log.Logger
    }
)

func NewDataComponent(options ...DataComponentOption) pbg.DataComponent {
    dc := &dataComponent{
        moves:    make([]pbg.Move, 100),
        pokèmons: make([]pbg.Pokèmon, 100),
        trainers: make(map[bson.ObjectId]pbg.Trainer),

        trainerFactory: NewTrainer,

        moveReqs:    make(chan moveReq),
        pokèmonReqs: make(chan pokèmonReq),
        trainerReqs: make(chan trainerReq),
    }

    for _, option := range options {
        if err := option(dc); err != nil {
            panic(err)
        }
    }

    go dc.moveLoop()
    go dc.pokèmonLoop()
    go dc.trainersLoop()

    return dc
}

func (dc *dataComponent) pokèmonLoop() {
    for req := range dc.pokèmonReqs {
        dc.Log("Serving Pokèmon request ", req)
        req(dc.pokèmons)
    }
}

func (dc *dataComponent) moveLoop() {
    for req := range dc.moveReqs {
        dc.Log("Serving Move request ", req)
        req(dc.moves)
    }
}

func (dc *dataComponent) trainersLoop() {
    for req := range dc.trainerReqs {
        dc.Log("Serving Trainer request ", req)
        req(dc.trainers)
    }
}

func (dc *dataComponent) Supply() pbg.DataInterface {
    return dc
}

func (dc *dataComponent) Retrieve(di pbg.DataInterface) {
    // nothing
}
