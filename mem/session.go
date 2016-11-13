package mem

import (
    "time"

    "github.com/ar3s3ru/PokemonBattleGo/pbg"
)

type session struct {
    User   pbg.Trainer `json:"user"`
    Tken  string      `json:"token"`
    expire time.Time
}

func (s *session) Reference() pbg.Trainer {
    return s.User
}

func (s *session) Token() string {
    return s.Tken
}

func (s *session) Expired() bool {
    return time.Now().After(s.expire)
}
