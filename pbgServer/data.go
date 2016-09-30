package pbgServer

import "gopkg.in/mgo.v2/bson"

// TODO: add error handling
// Interfaccia per l'accesso e la gestione dei dati sul server.
//
// Interface for data access and management onto the server.
type IDataMechanism interface {
    AddTrainer(trainer Trainer)     (bson.ObjectId, error)
    RemoveTrainer(id bson.ObjectId) error

    GetPokèmonById(id int)           (Pokèmon, error)
    GetTrainerById(id bson.ObjectId) (Trainer, error)
    GetTrainerByName(name string)    (Trainer, error)
}
