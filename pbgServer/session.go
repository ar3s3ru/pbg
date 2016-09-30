package pbgServer

import "time"

type (
    // TODO: add error handling?
    // Interfaccia che gestisce le sessioni utente all'interno del server.
    //
    // Interfaces that handles user's sessions inside the server.
    ISessionMechanism interface {
        AddSession(user User)    string
        GetSession(token string) Session
        RemoveSession(token string)

        Purge()
    }

    Session interface {
        GetUserReference() User
        GetToken()  string
        GetExpire() time.Time
        IsExpired() bool
    }
)
