package mem

import (
    "github.com/ar3s3ru/PokemonBattleGo/pbgServer"
    "gopkg.in/mgo.v2/bson"
    "sync"
    "golang.org/x/crypto/bcrypt"
    "time"
    "github.com/satori/go.uuid"
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
        sessions      map[string]pbgServer.Session
        // As always, it's a bottleneck...
        sessionsMutex sync.RWMutex
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
            sessions:      make(map[string]pbgServer.Session),
            sessionsMutex: sync.RWMutex{},
        }
    }
}

func (authority *memAuthority) AddSession(user pbgServer.Trainer) (pbgServer.Session, error) {
    if user == nil {
        return nil, pbgServer.ErrInvalidUserObject
    }

    sess := &session{
        user:   user,
        token:  uuid.NewV4().String(),
        expire: time.Now().Add(30 * time.Hour), // Token dura per 30 ore
    }

    authority.sessions[sess.GetToken()] = sess
    return sess, nil
}

func (authority *memAuthority) GetSession(token string) (pbgServer.Session, error) {
    authority.sessionsMutex.RLock()
    defer authority.sessionsMutex.RUnlock()

    if s, ok := authority.sessions[token]; !ok {
        return nil, pbgServer.ErrSessionNotFound
    } else {
        return s, nil
    }
}

func (authority *memAuthority) RemoveSession(token string) error {
    authority.sessionsMutex.Lock()
    defer authority.sessionsMutex.Unlock()

    if _, ok := authority.sessions[token]; !ok {
        return pbgServer.ErrSessionNotFound
    } else {
        delete(authority.sessions, token)
        return nil
    }
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
        return authority.AddSession(usr)
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
        // New trainer object
        trainer := &trainer{
            Name: username,
            hpwd: pwd,
            Sgup: time.Now(),
            set:  false,
            Tm:   [6]pbgServer.PokèmonTeam{ nil, nil, nil, nil, nil, nil },
            Cls:  pbgServer.TrainerC,
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
