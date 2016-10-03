package mem

import (
    "time"
    "github.com/ar3s3ru/PokemonBattleGo/pbgServer"
    "errors"
)

type (
    trainer struct {
        Name string                   `json:"username"`
        hpwd []byte                   `json:"-"`
        Sgup time.Time                `json:"sign_up"`
        // Trainer related fieds
        set  bool                     `json:"-"`
        Cls  pbgServer.TrainerClass   `json:"class"`
        Tm   [6]pbgServer.PokèmonTeam `json:"team"`
    }
)

func (t *trainer) GetName() string {
    return t.Name
}

func (t *trainer) GetPasswordHash() []byte {
    return t.hpwd
}

func (t *trainer) GetSignUpDate() time.Time {
    return t.Sgup
}

func (t *trainer) IsSet() bool {
    return t.set
}

func (t *trainer) GetClass() (pbgServer.TrainerClass, error) {
    if !t.IsSet() {
        return -1, errors.New("Not set!")
    }

    return t.Cls, nil
}

func (t *trainer) GetTeam() ([6]pbgServer.PokèmonTeam, error) {
    if !t.IsSet() {
        // TODO: change t.tm in nil of some sort
        return t.Tm, errors.New("Not set!")
    }

    return t.Tm, nil
}
