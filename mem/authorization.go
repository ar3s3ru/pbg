package mem

import (
    "log"
    "time"

    "github.com/ar3s3ru/PokemonBattleGo/pbg"
)

type (
    sessionRequest func(map[string]pbg.Session)
    sessionComponent struct {
        sessions       map[string]pbg.Session
        sessionReqs    chan sessionRequest
        sessionFactory pbg.SessionFactory

        logger     *log.Logger
        purgeTimer time.Duration
    }
)

func NewSessionComponent(options ...SessionComponentOption) pbg.SessionComponent {
    sc := &sessionComponent{
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

func (sc *sessionComponent) sessionLoop() {
    for req := range sc.sessionReqs {
        if sc.logger != nil {
            sc.logger.Println("Received Session request", req)
        }

        req(sc.sessions)
    }
}

func (sc *sessionComponent) purgingLoop() {
    for t := range time.NewTicker(sc.purgeTimer).C {
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
