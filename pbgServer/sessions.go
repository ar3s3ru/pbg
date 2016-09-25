package pbgServer

import "time"

type (
    ISessionsMechanism interface {
        AddSession(trainer Trainer) string
        GetSession(token string)    Session
        RemoveSession(token string)

        Purge()
    }

    Session interface {
        GetTrainerReference() Trainer
        GetToken()  string
        GetExpire() time.Time
        IsExpired() bool
    }
)
