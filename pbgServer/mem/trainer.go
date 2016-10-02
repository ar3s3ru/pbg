package mem

import (
    "time"
    "github.com/ar3s3ru/PokemonBattleGo/pbgServer"
    "errors"
)

type (
    trainer struct {
        name string
        hpwd []byte
        sgup time.Time
        // Trainer related fieds
        set  bool
        cls  pbgServer.TrainerClass
        tm   [6]pbgServer.PokèmonTeam
    }
)

func (t *trainer) GetName() string {
    return t.name
}

func (t *trainer) GetPasswordHash() []byte {
    return t.hpwd
}

func (t *trainer) GetSignUpDate() time.Time {
    return t.sgup
}

func (t *trainer) IsSet() bool {
    return t.set
}

func (t *trainer) GetClass() (pbgServer.TrainerClass, error) {
    if !t.IsSet() {
        return -1, errors.New("Not set!")
    }

    return t.cls, nil
}

func (t *trainer) GetTeam() ([6]pbgServer.PokèmonTeam, error) {
    if !t.IsSet() {
        // TODO: change t.tm in nil of some sort
        return t.tm, errors.New("Not set!")
    }

    return t.tm, nil
}
