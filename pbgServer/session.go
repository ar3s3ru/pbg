package pbgServer

import "time"

type (
    // TODO: add error handling?
    // Interfaccia che gestisce le sessioni utente all'interno del server.
    //
    // Interface that handles user's sessions inside the server.
    ISessionMechanism interface {
        // Gestione sessioni
        AddSession(user User)    string
        GetSession(token string) Session
        RemoveSession(token string)

        // Rimuove tutte le sessioni scadute
        Purge()
    }

    // Interfaccia che rappresenta una sessione utente sul server.
    // Una sessione viene identificata dal proprio token, e ha un certo periodo di validità.
    //
    // Interface that represent an user session onto the server.
    // A session is identified by its token, and it has a certain validity period.
    Session interface {
        GetUserReference() User
        GetToken()  string
        GetExpire() time.Time
        IsExpired() bool
    }
)
