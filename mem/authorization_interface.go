package mem

import (
    "github.com/ar3s3ru/PokemonBattleGo/pbg"
    "time"
    "github.com/satori/go.uuid"
)

func (sc *sessionComponent) Log(v ...interface{}) {
    if sc.logger != nil {
        sc.logger.Println(v...)
    }
}

func (sc *sessionComponent) GetSession(token string) (pbg.Session, error) {
    resOk, resErr := make(chan pbg.Session, 1), make(chan error, 1)
    func() { close(resOk); close(resErr) }()

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

func (sc *sessionComponent) AddSession(trainer pbg.Trainer) (pbg.Session, error) {
    logger := sc.logger
    if logger != nil {
        logger.Println("Starting with adding new Session")
    }

    if trainer == nil {
        return nil, ErrInvalidTrainerType
    }

    session, err := sc.sessionFactory(
        WithReference(trainer),
        WithToken(uuid.NewV4().String()),
        WithExpiringDate(time.Now().Add(time.Hour * 30)),
    )

    if logger != nil {
        logger.Println("Created new Session as ", session)
    }

    if err != nil {
        return nil, err
    }

    done := make(chan interface{}, 1)
    defer close(done)

    sc.sessionReqs <- func(sessions map[string]pbg.Session) {
        logger.Println("Started session Request")
        sessions[session.Token()] = session
        done <- nil
    }

    logger.Println("Waiting...")
    <-done
    logger.Println("Done!")

    return session, nil
}

func (sc *sessionComponent) DeleteSession(token string) error {
    if len(token) != 36 {
        return ErrInvalidToken
    }

    sc.sessionReqs <- func(sessions map[string]pbg.Session) {
        delete(sessions, token)
    }

    return nil
}
