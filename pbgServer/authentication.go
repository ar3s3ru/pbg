package pbgServer

import (
    "time"
    "gopkg.in/mgo.v2/bson"
    "errors"
)

type (
    // Interfaccia che realizza la logica di autenticazione nel server.
    //
    // Interface that realizes authentication logic into the server.
    IAuthMechanism interface {
        DoLogin(string, string) (Session, error)
        DoLogout(Session)       error

        Register(string, string) (User, bson.ObjectId, error)
        Unregister(Session)      error
    }

    // Interfaccia che rappresenta un utente registrato al server.
    // Possiede tutte le informazioni necessarie al login e all'identificazione univoca.
    //
    // Interface that represents an user registered onto the server.
    // It holds all the necessary information for login and unique identification.
    User interface {
        GetName()         string
        GetPasswordHash() []byte
        GetSignUpDate()   time.Time
    }
)

var (
    ErrInvalidSessionObject = errors.New("Invalid session object used: <nil>")
    ErrUserAlreadyExists    = errors.New("User already exists")
)
