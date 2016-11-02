package mem

import (
    "time"

    "github.com/ar3s3ru/PokemonBattleGo/pbg"
)

type session struct {
    User   pbg.Trainer `json:"user"`
    Token  string      `json:"token"`
    expire time.Time
}

func (s *session) GetUserReference() pbg.Trainer {
    return s.User
}

func (s *session) GetToken() string {
    return s.Token
}

func (s *session) IsExpired() bool {
    return time.Now().After(s.expire)
}
