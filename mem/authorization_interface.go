package mem

import (
	"time"

	"github.com/ar3s3ru/pbg"
	"github.com/satori/go.uuid"
)

func (sc *SessionDBComponent) Log(v ...interface{}) {
	if sc.logger != nil {
		sc.logger.Println(v...)
	}
}

func (sc *SessionDBComponent) GetSession(token string) (pbg.Session, error) {
	resOk, resErr := make(chan pbg.Session, 1), make(chan error, 1)
	defer func() { close(resOk); close(resErr) }()

	sc.sessionReqs <- func(sessions map[string]pbg.Session) {
		if s, ok := sessions[token]; !ok {
			resErr <- pbg.ErrSessionNotFound
		} else {
			resOk <- s
		}
	}

	select {
	case session := <-resOk:
		return session, nil
	case err := <-resErr:
		return nil, err
	}
}

func (sc *SessionDBComponent) AddSession(trainer pbg.Trainer, expire time.Time) (pbg.Session, error) {
	session, err := sc.sessionFactory(
		WithReference(trainer), WithToken(uuid.NewV4().String()), WithExpiringDate(expire),
	)

	if err != nil {
		return nil, err
	}

	done := make(chan pbg.Session, 1)
	defer close(done)

	sc.sessionReqs <- func(sessions map[string]pbg.Session) {
		sessions[session.Token()] = session
		done <- session
	}

	return <-done, nil
}

func (sc *SessionDBComponent) DeleteSession(token string) error {
	if len(token) != 36 {
		return ErrInvalidTokenValue
	}

	sc.sessionReqs <- func(sessions map[string]pbg.Session) {
		delete(sessions, token)
	}

	return nil
}
