package pbgServer

import "time"

type (
    // TODO: add error handling?
    // Interfaccia che realizza la logica di autenticazione nel server.
    //
    // Interface that realizes authentication logic into the server.
    IAuthMechanism interface {
        DoLogin(string, string) Session
        DoLogout(Session)
    }

    User interface {
        GetName()         string
        GetPasswordHash() string
        GetSignUpDate()   time.Time
    }
)
