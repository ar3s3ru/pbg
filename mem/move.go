package mem

import "github.com/ar3s3ru/PokemonBattleGo/pbg"

type move struct {
    Nam  string       `json:"name"`
    Typ  pbg.Type     `json:"type"`
    Cat  pbg.Category `json:"category"`
    Pps  int          `json:"pps"`
    Pwr  int          `json:"power"`
    Accy int          `json:"accuracy"`
    Prio int          `json:"priority"`
}

func (m *move) Name() string {
    return m.Nam
}

func (m *move) Type() pbg.Type {
    return m.Typ
}

func (m *move) Category() pbg.Category {
    return m.Cat
}

func (m *move) Priority() int {
    return m.Prio
}

func (m *move) Power() int {
    return m.Pwr
}

func (m *move) Accuracy() int {
    return m.Accy
}

func (m *move) PPs() int {
    return m.Pps
}
