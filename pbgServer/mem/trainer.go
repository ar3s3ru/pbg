package mem

import (
    "time"
    "github.com/ar3s3ru/PokemonBattleGo/pbgServer"
    "strconv"
)

type (
    trainer struct {
        Name string
        hpwd []byte
        Sgup time.Time
        // Trainer related fieds
        set  bool
        Cls  pbgServer.TrainerClass
        Tm   [6]pbgServer.PokèmonTeam
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

// Implements the Marshaler interface for JSON mashaling
func (t *trainer) MarshalJSON() ([]byte, error) {
    set := strconv.FormatBool(t.set)
    return []byte(`{"name":"` + t.Name + `","sign_up":"` + t.Sgup.String() + `","set":` + set + `}`), nil
}
