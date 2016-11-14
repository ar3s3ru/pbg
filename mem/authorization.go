package mem

import (
    "log"
    "time"

    "github.com/ar3s3ru/PokemonBattleGo/pbg"
)

type (
    sessionRequest func(map[string]pbg.Session)
    SessionDBComponent struct {
        sessions       map[string]pbg.Session
        sessionReqs    chan sessionRequest
        sessionFactory SessionFactory

        logger     *log.Logger
        purgeTimer time.Duration
    }
)

func NewSessionComponent(options ...SessionDBComponentOption) pbg.SessionComponent {
    sc := &SessionDBComponent{
        sessions:       make(map[string]pbg.Session),
        sessionReqs:    make(chan sessionRequest),
        sessionFactory: NewSession,
        purgeTimer:     time.Hour,
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

func (sc *SessionDBComponent) sessionLoop() {
    for req := range sc.sessionReqs {
        if sc.logger != nil {
            sc.logger.Println("Received Session request", req)
        }

        req(sc.sessions)
    }
}

func (sc *SessionDBComponent) purgingLoop() {
    for t := range time.NewTicker(sc.purgeTimer).C {
        if sc.logger != nil {
            sc.logger.Println(t, "- Executing session purging...")
        }

        sc.Purge()
    }
}

func (sc *SessionDBComponent) Supply() pbg.SessionInterface {
    return sc
}

func (sc *SessionDBComponent) Retrieve(si pbg.SessionInterface) {
    // nothing...
}

func (sc *SessionDBComponent) Purge() {
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
