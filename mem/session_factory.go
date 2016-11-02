package mem

import (
    "time"

    "github.com/ar3s3ru/PokemonBattleGo/pbg"
)

type sessionOption  func(*session) error
type SessionFactory func()         pbg.Session

func NewSessionFactory() pbg.SessionFactory {
    return SessionFactory(func() pbg.Session {
        return &session{User: nil, expire: time.Now()}
    })
}

func (sf SessionFactory) Create(options ...pbg.SessionFactoryOption) (pbg.Session, error) {
    session := sf()
    for _, option := range options {
        if err := option(session); err != nil {
            return nil, err
        }
    }

    return session, nil
}

func adaptSessionFactoryOption(option sessionOption) pbg.SessionFactoryOption {
    return func(sess pbg.Session) error {
        switch converted := sess.(type) {
        case *session:
            return option(converted)
        default:
            return ErrInvalidSessionType
        }
    }
}

func WithUserReference(user pbg.Trainer) pbg.SessionFactoryOption {
    return adaptSessionFactoryOption(func(session *session) error {
        session.User = user
        return nil
    })
}

func WithToken(token string) pbg.SessionFactoryOption {
    return adaptSessionFactoryOption(func(session *session) error {
        session.Token = token
        return nil
    })
}

func WithExpiringDate(expiring time.Time) pbg.SessionFactoryOption {
    return adaptSessionFactoryOption(func(session *session) error {
        session.expire = expiring
        return nil
    })
}
