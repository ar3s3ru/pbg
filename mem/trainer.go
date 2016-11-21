package mem

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/ar3s3ru/pbg"
)

var (
	trainerJSONBase = []byte(`{"name":"","sign_up":"","set":,"class":,"team":}`)
)

type Trainer struct {
	name       string
	hashedPass []byte
	signUp     time.Time

	setted bool
	class  pbg.TrainerClass
	team   [6]pbg.PokèmonTeam
}

func (t *Trainer) Name() string {
	return t.name
}

func (t *Trainer) PasswordHash() []byte {
	return t.hashedPass
}

func (t *Trainer) SignUpDate() time.Time {
	return t.signUp
}

func (t *Trainer) Set() bool {
	return t.setted
}

func (t *Trainer) Class() pbg.TrainerClass {
	return t.class
}

func (t *Trainer) Team() [6]pbg.PokèmonTeam {
	return t.team
}

func (t *Trainer) SetTrainer(team [6]pbg.PokèmonTeam, class pbg.TrainerClass) error {
	t.team = team
	t.class = class
	t.setted = true

	return nil
}

func (t *Trainer) UpdateTrainer(team [6]pbg.PokèmonTeam) error {
	t.setted = true
	t.team = team

	return nil
}

// Implements the Marshaler interface for JSON mashaling
func (t *Trainer) MarshalJSON() ([]byte, error) {
	class, _ := json.Marshal(t.Class())
	converted := [][]byte{
		[]byte(t.Name()),
		[]byte(t.SignUpDate().String()),
		[]byte(strconv.FormatBool(t.Set())),
		class,
		[]byte("null"),
	}

	if t.Set() {
		team, err := json.Marshal(t.Team())
		if err != nil {
			return nil, err
		} else {
			converted[4] = team
		}
	}

	lenght := len(trainerJSONBase) + argsLen(converted)
	trainer := make([]byte, lenght)

	last := copy(trainer, trainerJSONBase[:9])
	last += copy(trainer[last:], converted[0])
	last += copy(trainer[last:], trainerJSONBase[9:22])
	last += copy(trainer[last:], converted[1])
	last += copy(trainer[last:], trainerJSONBase[22:30])
	last += copy(trainer[last:], converted[2])
	last += copy(trainer[last:], trainerJSONBase[30:39])
	last += copy(trainer[last:], converted[3])
	last += copy(trainer[last:], trainerJSONBase[39:47])
	last += copy(trainer[last:], converted[4])
	last += copy(trainer[last:], trainerJSONBase[47:])

	return trainer, nil
}
