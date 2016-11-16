package mem

import (
	"encoding/json"
	"time"

	"github.com/ar3s3ru/pbg"
)

var (
	sessionJSONBase = []byte(`{"user":,"token":""}`)
)

type Session struct {
	user   pbg.Trainer
	token  string
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

	lenght := len(sessionJSONBase) + len(user) + len(s.token)
	session := make([]byte, lenght)

	last := copy(session, sessionJSONBase[:8])
	last += copy(session[last:], user)
	last += copy(session[last:], sessionJSONBase[8:18])
	last += copy(session[last:], []byte(s.token))
	last += copy(session[last:], sessionJSONBase[18:])

	return session, nil
}
