package mem

import (
    "github.com/ar3s3ru/PokemonBattleGo/pbgServer"
    "errors"
    "gopkg.in/mgo.v2/bson"
    "sync"
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

var (
    ErrInvalidDataMechanism = errors.New("Invalid DataMechanism value used in AuthorityBuilder (nil)")
)

func AuthBuilder() AuthorityBuilder {
    return &authBuilder{}
}

func (builder *authBuilder) UseDataMechanism(dm pbgServer.IDataMechanism) AuthorityBuilder {
    if dm == nil {
        panic(ErrInvalidDataMechanism)
    } else {
        builder.dataMechanism = dm
        return builder
    }
}

func (builder *authBuilder) Build() IAuthority {
    if builder.dataMechanism == nil {
        panic(ErrInvalidDataMechanism)
    } else {
        return &memAuthority{
            dataMechanism: builder.dataMechanism,
            sessions:      make(map[bson.ObjectId]pbgServer.Session),
            sessionsMutex: sync.Mutex{},
        }
    }
}

func (authority *memAuthority) AddSession(user pbgServer.User) string {
    // TODO: finish this
    return string(bson.NewObjectId())
}

func (authority *memAuthority) GetSession(token string) pbgServer.Session {
    // TODO: finish this
    return nil
}

func (authority *memAuthority) RemoveSession(token string) {
    // TODO: finish this
}

func (authority *memAuthority) Purge() {
    // TODO: finish this
}

func (authority *memAuthority) DoLogin(username string, password string) pbgServer.Session {
    // TODO: finish this
    return nil
}

func (authority *memAuthority) DoLogout(pbgServer.Session) {
    // TODO: finish this
}
