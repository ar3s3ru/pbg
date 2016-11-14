package mem

import (
    "github.com/ar3s3ru/PokemonBattleGo/pbg"
    "gopkg.in/mgo.v2/bson"
    "time"
)

type (
    synchronizedRequest func(chan<- interface{}, chan<- error, map[bson.ObjectId]pbg.Trainer)
)

func (tdb *trainerDBComponent) requestingWithReturn(req synchronizedRequest) (interface{}, error) {
    resOk, resErr := make(chan interface{}, 1), make(chan error, 1)
    defer func() { close(resOk); close(resErr) }()

    tdb.requests <- func(trainers map[bson.ObjectId]pbg.Trainer) {
        req(resOk, resErr, trainers)
    }

    select {
    case ok := <-resOk:
        return ok, nil
    case err := <-resErr:
        return nil, err
    }
}

func (tdb *trainerDBComponent) requestingGet(req synchronizedRequest) (pbg.Trainer, error) {
    trainer, err := tdb.requestingWithReturn(
        func(ok chan<- interface{}, bad chan<- error, trainers map[bson.ObjectId]pbg.Trainer) {
            req(ok, bad, trainers)
        },
    )

    if err != nil {
        return nil, err
    } else {
        return trainer.(pbg.Trainer), nil
    }
}

func (tdb *trainerDBComponent) Log(v ...interface{}) {
    if tdb.logger != nil {
        tdb.logger.Println(v...)
    }
}

func (tdb *trainerDBComponent) AddTrainer(options ...pbg.TrainerFactoryOption) (bson.ObjectId, error) {
    id, err := tdb.requestingWithReturn(
        func(ok chan<- interface{}, bad chan<- error, trainers map[bson.ObjectId]pbg.Trainer) {
            trainer, err := tdb.factory(options...)
            if err != nil {
                bad <- err
                return
            }

            for _, tr := range trainers {
                if trainer.Name() == tr.Name() {
                    bad <- pbg.ErrTrainerAlreadyExists
                    return
                }
            }

            id := bson.NewObjectIdWithTime(time.Now())
            trainers[id] = trainer

            ok <- id
        },
    )

    if err != nil {
        return "", err
    } else {
        return id.(bson.ObjectId), nil
    }
}

func (tdb *trainerDBComponent) GetTrainerByName(name string) (pbg.Trainer, error) {
    return tdb.requestingGet(
        func(ok chan<- interface{}, bad chan<- error, trainers map[bson.ObjectId]pbg.Trainer) {
            for _, trainer := range trainers {
                if trainer.Name() == name {
                    ok <- trainer
                    return
                }
            }

            bad <- pbg.ErrTrainerNotFound
        },
    )
}

func (tdb *trainerDBComponent) GetTrainerById(id bson.ObjectId) (pbg.Trainer, error) {
    return tdb.requestingGet(
        func(ok chan<- interface{}, bad chan<- error, trainers map[bson.ObjectId]pbg.Trainer) {
            if trainer, found := trainers[id]; !found {
                bad <- pbg.ErrTrainerNotFound
            } else {
                ok <- trainer
            }
        },
    )
}

func (tdb *trainerDBComponent) DeleteTrainer(id bson.ObjectId) error {
    _, err := tdb.requestingWithReturn(
        func(ok chan<- interface{}, bad chan<- error, trainers map[bson.ObjectId]pbg.Trainer) {
            if _, ok := trainers[id]; !ok {
                bad <- pbg.ErrTrainerNotFound
            } else {
                delete(trainers, id)
                bad <- nil
            }
        },
    )

    return err
}
