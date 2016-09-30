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

    // Interfaccia che rappresenta un utente registrato al server.
    // Possiede tutte le informazioni necessarie al login e all'identificazione univoca.
    //
    // Interface that represents an user registered onto the server.
    // It holds all the necessary information for login and unique identification.
    User interface {
        GetName()         string
        GetPasswordHash() string
        GetSignUpDate()   time.Time
    }
)
