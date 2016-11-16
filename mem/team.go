package mem

import (
	"encoding/json"
	"strconv"

	"github.com/ar3s3ru/pbg"
)

var (
	pokèmonTeamJSONBase = []byte(`{"pokemon":,"moves":,"level":,"ivs":,"evs":}`)
)

type PokèmonTeam struct {
	pbg.Pokèmon `json:"pokèmon"`

	moves [4]pbg.Move `json:"moves"`
	level int         `json:"level"`
	ivs   [6]int      `json:"ivs"`
	evs   [6]int      `json:"evs"`
}

func (pt *PokèmonTeam) Moves() [4]pbg.Move {
	return pt.moves
}

func (pt *PokèmonTeam) Level() int {
	return pt.level
}

func (pt *PokèmonTeam) IVs() [6]int {
	return pt.ivs
}

func (pt *PokèmonTeam) EVs() [6]int {
	return pt.evs
}

func (pt *PokèmonTeam) MarshalJSON() ([]byte, error) {
	entities := []interface{}{
		pt.Pokèmon, pt.moves, pt.ivs, pt.evs,
	}

	converted := [][]byte{
		nil, nil, nil, nil, []byte(strconv.Itoa(pt.level)),
	}

	lenght := len(pokèmonTeamJSONBase) + len(converted[4])
	for i, entity := range entities {
		marshaled, err := json.Marshal(entity)
		if err != nil {
			return nil, err
		}

		lenght += len(marshaled)
		converted[i] = marshaled
	}

	pokèmon := make([]byte, lenght)

	last := copy(pokèmon, pokèmonTeamJSONBase[:11])
	last += copy(pokèmon[last:], converted[0])
	last += copy(pokèmon[last:], pokèmonTeamJSONBase[11:20])
	last += copy(pokèmon[last:], converted[1])
	last += copy(pokèmon[last:], pokèmonTeamJSONBase[20:29])
	last += copy(pokèmon[last:], converted[4])
	last += copy(pokèmon[last:], pokèmonTeamJSONBase[29:36])
	last += copy(pokèmon[last:], converted[2])
	last += copy(pokèmon[last:], pokèmonTeamJSONBase[36:43])
	last += copy(pokèmon[last:], converted[3])
	last += copy(pokèmon[last:], pokèmonTeamJSONBase[43:])

	return pokèmon, nil
}
