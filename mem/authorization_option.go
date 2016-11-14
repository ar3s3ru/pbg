package mem

import (
    "log"
    "time"

    "github.com/ar3s3ru/PokemonBattleGo/pbg"
)

type (
    sessionComponentOption func(*sessionComponent)    error
    SessionComponentOption func(pbg.SessionComponent) error
)

func adaptSessionComponentOption(option sessionComponentOption) SessionComponentOption {
    return func(sc pbg.SessionComponent) error {
        if scc, ok := sc.(*sessionComponent); !ok {
            return pbg.ErrInvalidSessionComponent
        } else {
            return option(scc)
        }
    }
}

func WithSessionDBLogger(logger *log.Logger) SessionComponentOption {
    return adaptSessionComponentOption(func(sc *sessionComponent) error {
        sc.logger = logger
        return nil
    })
}

func WithPurgeTime(time time.Duration) SessionComponentOption {
    return adaptSessionComponentOption(func(sc *sessionComponent) error {
        sc.purgeTimer = time
        return nil
    })
}