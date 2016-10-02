package pbgServer

import (
    "gopkg.in/mgo.v2/bson"
    "errors"
)

// Interfaccia per l'accesso e la gestione dei dati sul server.
//
// Interface for data access and management onto the server.
type IDataMechanism interface {
    AddTrainer(trainer Trainer)     (bson.ObjectId, error)
    RemoveTrainer(id bson.ObjectId) error

    GetPokèmons() []Pokèmon
    GetMoves()    []Move

    GetMoveById(id int)              (Move, error)
    GetPokèmonById(id int)           (Pokèmon, error)
    GetTrainerById(id bson.ObjectId) (Trainer, error)
    GetTrainerByName(name string)    (Trainer, error)
}

var (
    ErrPokèmonNotFound    = errors.New("Pokèmon not found")
    ErrMoveNotFound       = errors.New("Move not found")
    ErrTrainerNotFound    = errors.New("Trainer not found")

    ErrIllegalTrainer     = errors.New("Trainer object is nil")
    ErrInvalidTrainerName = errors.New("Invalid Trainer name used")
)
