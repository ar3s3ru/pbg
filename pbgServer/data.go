package pbgServer

import (
    "gopkg.in/mgo.v2/bson"
    "errors"
)

// Interfaccia per l'accesso e la gestione dei dati sul server.
//
// Interface for data access and management onto the server.
type IDataMechanism interface {
    AddTrainer(trainer Trainer)                      (bson.ObjectId, error)
    AddPokèmonTeam(int, [4]int, int, [6]int, [6]int) (PokèmonTeam, error)
    RemoveTrainer(id bson.ObjectId) error

    GetPokèmons() []Pokèmon
    GetMoves()    []Move

    GetMoveById(id int)              (Move, error)
    GetPokèmonById(id int)           (Pokèmon, error)
    GetTrainerById(id bson.ObjectId) (Trainer, error)
    GetTrainerByName(name string)    (Trainer, error)
}

var (
    ErrInvalidPokemonLevel = errors.New("Invalid Pokèmon level, must be 1 to 100")
    ErrInvalidPokemonIVs   = errors.New("Invalid Pokèmon IVs, every field must go from 0 to 31, for a max of 6*31")
    ErrInvalidPokemonEVs   = errors.New("Invalid Pokèmon EVs, every field must go from 0 to 255, with a maximum of 510 summed")
    ErrInvalidPokemonId    = errors.New("Invalid Pokèmon Id used")
    ErrInvalidMoveId       = errors.New("Invalid Move Id used")
    ErrInvalidFirstPokemon = errors.New("You must add a Pokèmon to the 1st position at least")

    ErrPokèmonNotFound    = errors.New("Pokèmon not found")
    ErrMoveNotFound       = errors.New("Move not found")
    ErrTrainerNotFound    = errors.New("Trainer not found")

    ErrIllegalTrainer     = errors.New("Trainer object is nil")
    ErrInvalidTrainerName = errors.New("Invalid Trainer name used")
)
