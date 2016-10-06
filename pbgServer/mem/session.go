package mem

import (
    "github.com/ar3s3ru/PokemonBattleGo/pbgServer"
    "time"
)

type (
    session struct {
        User   pbgServer.Trainer `json:"user"`
        Token  string            `json:"token"`
        expire time.Time
    }
)

func (s *session) GetUserReference() pbgServer.Trainer {
    return s.User
}

func (s *session) GetToken() string {
    return s.Token
}

func (s *session) IsExpired() bool {
    return time.Now().After(s.expire)
}
