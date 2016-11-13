package mem

import (
    "time"
    "strconv"
    "encoding/json"

    "github.com/ar3s3ru/PokemonBattleGo/pbg"
)

type trainer struct {
    name       string
    hashedPass []byte
    signUp     time.Time

    setted bool
    class  pbg.TrainerClass
    team   [6]pbg.PokèmonTeam
}

func (t *trainer) Name() string {
    return t.name
}

func (t *trainer) PasswordHash() []byte {
    return t.hashedPass
}

func (t *trainer) SignUpDate() time.Time {
    return t.signUp
}

func (t *trainer) Set() bool {
    return t.setted
}

func (t *trainer) Class() pbg.TrainerClass {
    return t.class
}

func (t *trainer) Team() [6]pbg.PokèmonTeam {
    return t.team
}

func (t *trainer) SetTrainer(team [6]pbg.PokèmonTeam, class pbg.TrainerClass) error {
    t.team   = team
    t.class  = class
    t.setted = true

    return nil
}

func (t *trainer) UpdateTrainer(team [6]pbg.PokèmonTeam) error {
    t.setted = true
    t.team   = team

    return nil
}

// Implements the Marshaler interface for JSON mashaling
func (t *trainer) MarshalJSON() ([]byte, error) {
    set  := strconv.FormatBool(t.setted)
    base := `{"name":"` + t.Name() + `","sign_up":"` + t.signUp.String() + `","set":` + set

    if !t.Set() {
        return []byte(base + `}`), nil
    } else if team, err := json.Marshal(t.Team()); err != nil {
        return nil, err
    } else {
        return []byte(base + `,"team":` + string(team) + `,"class":"` + t.Class().String() + `"}`), nil
    }
}
