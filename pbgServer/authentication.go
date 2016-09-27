package pbgServer

import "time"

type (
    // TODO: add error handling?
    IAuthMechanism interface {
        DoLogin(string, string) Session
        DoLogout(Session)
    }

    IAuthBuilder interface {
        UseSessionMechanism(ISessionMechanism) IAuthBuilder
        Build() IAuthMechanism
    }

    User interface {
        GetName()         string
        GetPasswordHash() string
        GetSignUpDate()   time.Time
    }
)
