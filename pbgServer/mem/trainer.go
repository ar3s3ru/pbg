package mem

import (
    "time"
    "github.com/ar3s3ru/PokemonBattleGo/pbgServer"
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
        return -1, pbgServer.ErrTrainerNotSet
    }

    return t.Cls, nil
}

func (t *trainer) GetTeam() ([6]pbgServer.PokèmonTeam, error) {
    if !t.IsSet() {
        // TODO: change t.tm in nil of some sort
        return t.Tm, pbgServer.ErrTrainerNotSet
    }

    return t.Tm, nil
}

func (t *trainer) SetTrainer(team [6]pbgServer.PokèmonTeam, class pbgServer.TrainerClass) error {
    // TODO: finish this
    return nil
}

func (t *trainer) UpdateTrainer(team [6]pbgServer.PokèmonTeam) error {
    // TODO: finish this
    return nil
}
