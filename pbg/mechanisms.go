package pbg

import "gopkg.in/mgo.v2/bson"

type (
    DataMechanism interface {
        GetPokèmon(int) (Pokèmon, error)
        GetMove(int)    (Move, error)

        ListPokèmon() []Pokèmon
        ListMoves()   []Move

        GetTrainerById(bson.ObjectId) (Trainer, error)
        GetTrainerByName(string)      (Trainer, error)

        AddTrainer(string, string)   (bson.ObjectId, error)
        RemoveTrainer(bson.ObjectId) error

        GetPokèmonFactory() PokèmonFactory
        GetMoveFactory()    MoveFactory
        GetTrainerFactory() TrainerFactory
        GetTeamFactory()    TeamFactory
    }

    AuthorizationMechanism interface {
        CheckAuthorization([]byte) (statusCode int, session Session, err error)
    }

    SessionMechanism interface {
        AddSession(Trainer)   (Session, error)
        GetSession(string)    (Session, error)
        RemoveSession(string) error
        Purge()

        GetSessionFactory() SessionFactory
    }

    Session interface {
        GetUserReference() Trainer
        GetToken()         string
        IsExpired()        bool
    }
)
