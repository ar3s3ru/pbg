package pbgServer

import "gopkg.in/mgo.v2/bson"

type IDataMechanism interface {
    AddTrainer(trainer Trainer)     (bson.ObjectId, error)
    RemoveTrainer(id bson.ObjectId) error

    GetPokèmonById(id bson.ObjectId) (Pokèmon, error)
    GetTrainerById(id bson.ObjectId) (Trainer, error)
    GetTrainerByName(name string)    (Trainer, error)
}
