package mem

import (
    "time"
    "encoding/json"

    "github.com/ar3s3ru/PokemonBattleGo/pbg"
)

var (
    sessionMarshalBase = []byte(`{"user":,"token":""`)
)

type Session struct {
    user   pbg.Trainer `json:"user"`
    token  string      `json:"token"`
    expire time.Time
}

func (s *Session) Reference() pbg.Trainer {
    return s.user
}

func (s *Session) Token() string {
    return s.token
}

func (s *Session) Expired() bool {
    return time.Now().After(s.expire)
}

func (s *Session) MarshalJSON() ([]byte, error) {
    user, err := json.Marshal(s.user)
    if err != nil {
        return nil, err
    }

    session := make([]byte, len(sessionMarshalBase) + len(user) + len(s.token))

    session = append(session, sessionMarshalBase[:8]...)
    session = append(session, user...)
    session = append(session, sessionMarshalBase[8:18]...)
    session = append(session, []byte(s.token)...)
    session = append(session, sessionMarshalBase[18:]...)

    return session, nil
}