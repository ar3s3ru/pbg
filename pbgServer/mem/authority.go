package mem

import (
    "github.com/ar3s3ru/PokemonBattleGo/pbgServer"
    "gopkg.in/mgo.v2/bson"
    "sync"
    "golang.org/x/crypto/bcrypt"
    "time"
)

type (
    IAuthority interface {
        pbgServer.IAuthMechanism
        pbgServer.ISessionMechanism
    }

    AuthorityBuilder interface {
        UseDataMechanism(pbgServer.IDataMechanism) AuthorityBuilder
        Build() IAuthority
    }

    authBuilder struct {
        dataMechanism pbgServer.IDataMechanism
    }

    memAuthority struct {
        dataMechanism pbgServer.IDataMechanism
        sessions      map[bson.ObjectId]pbgServer.Session
        sessionsMutex sync.Mutex
    }
)

func AuthBuilder() AuthorityBuilder {
    return &authBuilder{}
}

func (builder *authBuilder) UseDataMechanism(dm pbgServer.IDataMechanism) AuthorityBuilder {
    if dm == nil {
        panic(pbgServer.ErrInvalidDataMechanism)
    } else {
        builder.dataMechanism = dm
        return builder
    }
}

func (builder *authBuilder) Build() IAuthority {
    if builder.dataMechanism == nil {
        panic(pbgServer.ErrInvalidDataMechanism)
    } else {
        return &memAuthority{
            dataMechanism: builder.dataMechanism,
            sessions:      make(map[bson.ObjectId]pbgServer.Session),
            sessionsMutex: sync.Mutex{},
        }
    }
}

func (authority *memAuthority) AddSession(user pbgServer.User) pbgServer.Session {
    // TODO: finish this
    return nil
}

func (authority *memAuthority) GetSession(token string) (pbgServer.Session, error) {
    authority.sessionsMutex.Lock()
    defer authority.sessionsMutex.Unlock()

    for _, v := range authority.sessions {
        if v.GetToken() == token {
            return v, nil
        }
    }

    return nil, pbgServer.ErrSessionNotFound
}

func (authority *memAuthority) RemoveSession(token string) error {
    authority.sessionsMutex.Lock()
    defer authority.sessionsMutex.Unlock()

    for k, v := range authority.sessions {
        if v.GetToken() == token {
            delete(authority.sessions, k)
            return nil
        }
    }

    return pbgServer.ErrSessionNotFound
}

func (authority *memAuthority) Purge() {
    authority.sessionsMutex.Lock()
    defer authority.sessionsMutex.Unlock()

    for k, v := range authority.sessions {
        if v.IsExpired() {
            delete(authority.sessions, k)
        }
    }
}

func (authority *memAuthority) DoLogin(username string, password string) (pbgServer.Session, error) {
    if usr, err := authority.dataMechanism.GetTrainerByName(username); err != nil {
        return nil, err
    } else if err := bcrypt.CompareHashAndPassword([]byte(usr.GetPasswordHash()), []byte(password)); err != nil {
        return nil, err
    } else {
        return authority.AddSession(usr), nil
    }
}

func (authority *memAuthority) DoLogout(session pbgServer.Session) error {
    if session != nil {
        return authority.RemoveSession(session.GetToken())
    }

    return pbgServer.ErrInvalidSessionObject
}

func (authority *memAuthority) Register(username string, password string) (pbgServer.User, bson.ObjectId, error) {
    if _, err := authority.dataMechanism.GetTrainerByName(username); err == nil {
        // Error, probably already exists
        return nil, "", pbgServer.ErrUserAlreadyExists
    } else if pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err != nil {
        // Problem with salting...
        return nil, "", err
    } else {
        trainer := &trainer{
            name: username,
            hpwd: pwd,
            sgup: time.Now(),
            set:  false,
            tm:   [...]pbgServer.PokèmonTeam{ nil, nil, nil, nil, nil, nil },
            cls:  pbgServer.TrainerC,
        }

        if id, err := authority.dataMechanism.AddTrainer(trainer); err != nil {
            return nil, "", err
        } else {
            return trainer, id, nil
        }
    }
}

func (authority *memAuthority) Unregister(session pbgServer.Session) error {
    if session == nil {
        return pbgServer.ErrInvalidSessionObject
    }
    // TODO: finish this
    return nil
}
