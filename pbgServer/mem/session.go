package mem

import (
    "github.com/ar3s3ru/PokemonBattleGo/pbgServer"
    "time"
)

type (
    session struct {
        user   pbgServer.User
        token  string
        expire time.Time
    }
)

func (s *session) GetUserReference() pbgServer.User {
    return s.user
}

func (s *session) GetToken() string {
    return s.token
}

func (s *session) IsExpired() bool {
    return time.Now().After(s.expire)
}
