package pbgServer

import "time"

type (
    IAuthMechanism interface {
        // TODO
    }

    User interface {
        GetName()         string
        GetPasswordHash() string
        GetSignUpDate()   time.Time
    }
)
