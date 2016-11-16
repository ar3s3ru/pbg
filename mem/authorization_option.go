package mem

import (
	"log"
	"time"
)

type (
	SessionDBComponentOption func(*SessionDBComponent) error
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
