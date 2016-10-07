package mem

import (
    "time"
    "github.com/ar3s3ru/PokemonBattleGo/pbgServer"
    "strconv"
    "encoding/json"
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

func (t *trainer) GetClass() pbgServer.TrainerClass {
    return t.Cls
}

func (t *trainer) GetTeam() [6]pbgServer.PokèmonTeam {
    return t.Tm
}

func (t *trainer) SetTrainer(team [6]pbgServer.PokèmonTeam, class pbgServer.TrainerClass) error {
    t.Tm  = team
    t.Cls = class
    t.set = true

    return nil
}

func (t *trainer) UpdateTrainer(team [6]pbgServer.PokèmonTeam) error {
    t.set = true
    t.Tm  = team

    return nil
}

// Implements the Marshaler interface for JSON mashaling
func (t *trainer) MarshalJSON() ([]byte, error) {
    set  := strconv.FormatBool(t.set)
    base := `{"name":"` + t.Name + `","sign_up":"` + t.Sgup.String() + `","set":` + set

    if !t.IsSet() {
        return []byte(base + `}`), nil
    } else if team, err := json.Marshal(t.GetTeam()); err != nil {
        return nil, err
    } else {
        return []byte(base + `,"team":` + string(team) + `,"class":"` + t.GetClass().String() + `"}`), nil
    }
}
