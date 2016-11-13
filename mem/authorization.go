package mem

import (
    "log"
    "time"

    "github.com/ar3s3ru/PokemonBattleGo/pbg"
)

//type (
//    sessionMechanism struct {
//        sessions       map[string]pbg.Session
//        sessionMutex   sync.RWMutex
//        sessionFactory pbg.SessionFactory
//    }
//
//    authLockHandler func() (pbg.Session, error)
//)
//
//func NewSessionMechanism() pbg.SessionMechanism {
//    return &sessionMechanism{
//        sessions:       make(map[string]pbg.Session),
//        sessionFactory: NewSessionFactory(),
//    }
//}
//
//func (sm *sessionMechanism) AddSession(trainer pbg.Trainer) (pbg.Session, error) {
//    session, err := sm.GetSessionFactory().Create(
//        WithUserReference(trainer),
//        WithToken(uuid.NewV4().String()),
//        WithExpiringDate(time.Now().Add(24 * time.Hour)),
//    )
//
//    if err != nil {
//        return nil, err
//    }
//
//    return sm.writeAuthLockedRegion(
//        sm.handleAddSession(session),
//    )
//}
//
//func (sm *sessionMechanism) GetSession(token string) (pbg.Session, error) {
//    return sm.writeAuthLockedRegion(
//        sm.handleGetSession(token),
//    )
//}
//
//func (sm *sessionMechanism) RemoveSession(token string) error {
//    _, err := sm.writeAuthLockedRegion(
//        sm.handleRemoveSession(token),
//    )
//
//    return err
//}
//
//func (sm *sessionMechanism) Purge() {
//    sm.writeAuthLockedRegion(
//        sm.handlePurge(),
//    )
//}
//
//func (sm *sessionMechanism) GetSessionFactory() pbg.SessionFactory {
//    return sm.sessionFactory
//}
//
//func (sm *sessionMechanism) CheckAuthorization(header []byte) (statusCode int, sess pbg.Session, err error) {
//    // Session initial value
//    sess = nil
//    // Decoding Authorization header
//    token, err := basicAuthorization(header)
//    if err == pbg.ErrInvalidAuthorizationHeader {
//        statusCode = fasthttp.StatusBadRequest
//        return
//    }
//
//    // Checking session
//    if sess, err = sm.GetSession(string(token)); err == pbg.ErrSessionExpired || err == pbg.ErrSessionNotFound {
//        statusCode = fasthttp.StatusUnauthorized
//        return
//    } else if err != nil {
//        statusCode = fasthttp.StatusInternalServerError
//        return
//    }
//    // Everything went well!
//    statusCode = fasthttp.StatusOK
//    return
//}
//
//func (sm *sessionMechanism) readAuthLockedRegion(handler authLockHandler) (pbg.Session, error) {
//    sm.sessionMutex.RLock()
//    defer sm.sessionMutex.RUnlock()
//
//    return handler()
//}
//
//func (sm *sessionMechanism) writeAuthLockedRegion(handler authLockHandler) (pbg.Session, error) {
//    sm.sessionMutex.Lock()
//    defer sm.sessionMutex.Unlock()
//
//    return handler()
//}
//
//func (sm *sessionMechanism) handleAddSession(session pbg.Session) authLockHandler {
//    return func() (pbg.Session, error) {
//        sm.sessions[session.GetToken()] = session
//        return session, nil
//    }
//}
//
//func (sm *sessionMechanism) handleGetSession(token string) authLockHandler {
//    return func() (pbg.Session, error) {
//        if sess, ok := sm.sessions[token]; ok {
//            return sess, nil
//        }
//
//        return nil, pbg.ErrSessionNotFound
//    }
//}
//
//func (sm *sessionMechanism) handleRemoveSession(token string) authLockHandler {
//    return func() (pbg.Session, error) {
//        if _, ok := sm.sessions[token]; ok {
//            delete(sm.sessions, token)
//            return nil, nil
//        }
//
//        return nil, pbg.ErrSessionNotFound
//    }
//}
//
//func (sm *sessionMechanism) handlePurge() authLockHandler {
//    return func() (pbg.Session, error) {
//        for key, session := range sm.sessions {
//            // Session is expired
//            if session.IsExpired() {
//                // remove key from sessions map
//                delete(sm.sessions, key)
//            }
//        }
//        // We don't mind about return types at all
//        return nil, nil
//    }
//}
type (
    sessionRequest func(map[string]pbg.Session)
    sessionComponent struct {
        sessions       map[string]pbg.Session
        sessionReqs    chan sessionRequest
        sessionFactory pbg.SessionFactory

        logger *log.Logger
    }
)

func NewSessionComponent(options ...SessionComponentOption) pbg.SessionComponent {
    sc := &sessionComponent{
        sessions:       make(map[string]pbg.Session),
        sessionReqs:    make(chan sessionRequest),
        sessionFactory: NewSession,
    }

    for _, option := range options {
        if err := option(sc); err != nil {
            panic(err)
        }
    }

    go sc.sessionLoop()
    go sc.purgingLoop()

    return sc
}

func (sc *sessionComponent) sessionLoop() {
    for req := range sc.sessionReqs {
        if sc.logger != nil {
            sc.logger.Println("Received Session request", req)
        }

        req(sc.sessions)
    }
}

func (sc *sessionComponent) purgingLoop() {
    // TODO: add option to specify timing
    for t := range time.NewTicker(time.Hour).C {
        if sc.logger != nil {
            sc.logger.Println(t, "- Executing session purging...")
        }

        sc.Purge()
    }
}

func (sc *sessionComponent) Supply() pbg.SessionInterface {
    return sc
}

func (sc *sessionComponent) Retrieve(si pbg.SessionInterface) {
    // nothing...
}

func (sc *sessionComponent) Purge() {
    sc.sessionReqs <- func(sessions map[string]pbg.Session) {
        for key, session := range sessions {
            if session.Expired() {
                delete(sessions, key)
            }
        }

        if sc.logger != nil {
            sc.logger.Println("Purge terminated!")
        }
    }
}
