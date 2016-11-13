package mem

import (
    "time"

    "gopkg.in/mgo.v2/bson"
    "golang.org/x/crypto/bcrypt"

    "github.com/ar3s3ru/PokemonBattleGo/pbg"
)

func (dc *dataComponent) Log(v ...interface{}) {
    if dc.logger != nil {
        dc.logger.Println(v...)
    }
}

func (dc *dataComponent) GetPokèmon(id int) (pbg.Pokèmon, error) {
    resOk, resErr := make(chan pbg.Pokèmon, 1), make(chan error, 1)
    defer func() { close(resOk); close(resErr) }()

    dc.pokèmonReqs <- func(pokèmons []pbg.Pokèmon) {
        if inRange(id, len(pokèmons)) {
            resOk <- pokèmons[id - 1]
        } else {
            resErr <- pbg.ErrPokèmonNotFound
        }
    }

    select {
    case pokèmon := <-resOk:
        return pokèmon, nil
    case err := <-resErr:
        return nil, err
    }
}

//func (dc *dataComponent) AddPokèmon(pkm pbg.Pokèmon) error {
//    result := make(chan error, 1)
//    dc.pokèmonReqs <- func(pokèmons []pbg.Pokèmon) {
//        if pkm != nil {
//            pokèmons = append(pokèmons, pkm)
//            result <- nil
//        } else {
//            result <- ErrInvalidPokèmonType
//        }
//    }
//
//    return <-result
//}

func (dc *dataComponent) GetMove(id int) (pbg.Move, error) {
    resOk, resErr := make(chan pbg.Move, 1), make(chan error, 1)
    defer func() { close(resOk); close(resErr) }()

    dc.moveReqs <- func(moves []pbg.Move) {
        if inRange(id, len(moves)) {
            resOk <- moves[id]
        } else {
            resErr <- pbg.ErrMoveNotFound
        }
    }

    select {
    case move := <-resOk:
        return move, nil
    case err := <-resErr:
        return nil, err
    }
}

func (dc *dataComponent) trainerHandler(req func(chan<- interface{}, chan<- error, map[bson.ObjectId]pbg.Trainer)) (interface{}, error) {
    resOk, resErr := make(chan interface{}, 1), make(chan error, 1)
    defer func() { close(resOk); close(resErr) }()

    dc.trainerReqs <- func(trainers map[bson.ObjectId]pbg.Trainer) {
        req(resOk, resErr, trainers)
    }

    select {
    case ok := <-resOk:
        return ok, nil
    case err := <-resErr:
        return nil, err
    }
}

func (dc *dataComponent) trainerGet(req func(chan<- interface{}, chan<- error, map[bson.ObjectId]pbg.Trainer)) (pbg.Trainer, error) {
    trainer, err := dc.trainerHandler(func(ok chan<- interface{}, err chan<- error, trainers map[bson.ObjectId]pbg.Trainer) {
        req(ok, err, trainers)
    })

    if err != nil {
        return nil, err
    }
    // TODO: maybe type assertion?
    return trainer.(pbg.Trainer), err
}

func (dc *dataComponent) GetTrainerByName(name string) (pbg.Trainer, error) {
    return dc.trainerGet(func(ok chan<- interface{}, err chan<- error, trainers map[bson.ObjectId]pbg.Trainer) {
        for _, trainer := range trainers {
            if trainer.Name() == name {
                ok <- trainer
                return
            }
        }

        err <- pbg.ErrTrainerNotFound
    })
}

func (dc *dataComponent) GetTrainerById(id bson.ObjectId) (pbg.Trainer, error) {
    return dc.trainerGet(func(cok chan<- interface{}, err chan<- error, trainers map[bson.ObjectId]pbg.Trainer) {
        if trainer, ok := trainers[id]; !ok {
            err <- pbg.ErrTrainerNotFound
        } else {
            cok <- trainer
        }
    })
}

func (dc *dataComponent) AddTrainer(user, pass string) (bson.ObjectId, error) {
    id, err := dc.trainerHandler(func(ok chan<- interface{}, cerr chan<- error, trainers map[bson.ObjectId]pbg.Trainer) {
        for _, trainer := range trainers {
            if trainer.Name() == user {
                cerr <- pbg.ErrTrainerAlreadyExists
                return
            }
        }

        if pwd, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost); err != nil {
            cerr <- pbg.ErrPasswordSalting
        } else if trainer, err := dc.trainerFactory(WithTrainerName(user), WithTrainerPassword(pwd)); err != nil {
            cerr <- err
        } else {
            id := bson.NewObjectIdWithTime(time.Now())
            trainers[id] = trainer

            ok <- id
        }
    })

    dc.Log("Function sended onto the Trainer request queue")
    dc.Log("Got result from Trainer request:", id)

    if err != nil {
        return "", err
    }

    return id.(bson.ObjectId), err
}

func (dc *dataComponent) DeleteTrainer(id bson.ObjectId) error {
    _, err := dc.trainerHandler(func(_ chan<- interface{}, result chan<- error, trainers map[bson.ObjectId]pbg.Trainer) {
        if _, ok := trainers[id]; !ok {
            result <- pbg.ErrTrainerNotFound
        } else {
            delete(trainers, id)
            result <- nil
        }
    })

    return err
}

//func (dc *dataComponent) UpdateTrainer(id bson.ObjectId, newTr pbg.Trainer) error {
//    return dc.trainerHandl(func(result chan<- interface{}, trainers map[bson.ObjectId]pbg.Trainer) {
//        if newTr == nil {
//            result <- ErrInvalidTrainerType
//        } else if _, ok := trainers[id]; !ok {
//            result <- pbg.ErrTrainerNotFound
//        } else {
//            for kid, trainer := range trainers {
//                if trainer.GetName() == newTr.GetName() && kid != id {
//
//                }
//            }
//
//            trainers[id] = newTr
//            result <- nil
//        }
//    })
//}
