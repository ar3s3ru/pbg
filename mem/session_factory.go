package mem

import (
	"time"

	"github.com/ar3s3ru/PokemonBattleGo/pbg"
)

type (
	SessionFactory func(...SessionOption) (pbg.Session, error)
	SessionOption  func(*Session) error
)

func NewSession(options ...SessionOption) (pbg.Session, error) {
	session := &Session{user: nil, expire: time.Now()}

	for _, option := range options {
		if err := option(session); err != nil {
			return nil, err
		}
	}

	return session, nil
}

func WithReference(user pbg.Trainer) SessionOption {
	return func(session *Session) error {
		if user == nil {
			return ErrInvalidTrainerValue
		}

		session.user = user
		return nil
	}
}

func WithToken(token string) SessionOption {
	return func(session *Session) error {
		if token == "" {
			return ErrInvalidTokenValue
		}

		session.token = token
		return nil
	}
}

func WithExpiringDate(expiring time.Time) SessionOption {
	return func(session *Session) error {
		// Time can't be nil or invalid, by language promise
		session.expire = expiring
		return nil
	}
}
