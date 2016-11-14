package mem

import (
    "github.com/ar3s3ru/PokemonBattleGo/pbg"
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
    case err := <- resErr:
        return nil, err
    }
}

func (sc *SessionDBComponent) AddSession(options ...pbg.SessionFactoryOption) (pbg.Session, error) {
    sc.Log("Starting with adding new Session")

    session, err := sc.sessionFactory(options...)
    if err != nil {
        return nil, err
    }

    sc.Log("Created new Session at", session)

    done := make(chan interface{}, 1)
    defer close(done)

    sc.sessionReqs <- func(sessions map[string]pbg.Session) {
        sc.Log("Started session Request")
        sessions[session.Token()] = session
        done <- nil
    }

    sc.Log("Waiting...")
    <-done
    sc.Log("Done!")

    return session, nil
}

func (sc *SessionDBComponent) DeleteSession(token string) error {
    if len(token) != 36 {
        return ErrInvalidToken
    }

    sc.sessionReqs <- func(sessions map[string]pbg.Session) {
        delete(sessions, token)
    }

    return nil
}
