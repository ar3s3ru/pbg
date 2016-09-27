package pbgServer

import "time"

type (
    // TODO: add error handling?
    ISessionMechanism interface {
        AddSession(user User)    string
        GetSession(token string) Session
        RemoveSession(token string)

        Purge()
    }

    ISessBuilder interface {
        UseDataMechanism(IDataMechanism) IDataMechanism
        Build() ISessionMechanism
    }

    Session interface {
        GetUserReference() User
        GetToken()  string
        GetExpire() time.Time
        IsExpired() bool
    }
)
