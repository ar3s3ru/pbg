package mem

import (
    "log"
    "time"

    "github.com/ar3s3ru/PokemonBattleGo/pbg"
)

type (
    SessionDBComponentOption func(pbg.SessionComponent) error
)

func WithSessionDBLogger(logger *log.Logger) SessionDBComponentOption {
    return func(sc *SessionDBComponent) error {
        sc.logger = logger
        return nil
    }
}

func WithPurgeTime(time time.Duration) SessionDBComponentOption {
    return func(sc *SessionDBComponent) error {
        sc.purgeTimer = time
        return nil
    }
}