package mem

import "github.com/ar3s3ru/PokemonBattleGo/pbg"

type move struct {
    Name string       `json:"name"`
    Typ  pbg.Type     `json:"type"`
    Cat  pbg.Category `json:"category"`
    PPs  int          `json:"pps"`
    Pwr  int          `json:"power"`
    Accy int          `json:"accuracy"`
    Prio int          `json:"priority"`
}

func (m *move) GetName() string {
    return m.Name
}

func (m *move) GetType() pbg.Type {
    return m.Typ
}

func (m *move) GetCategory() pbg.Category {
    return m.Cat
}

func (m *move) GetPriority() int {
    return m.Prio
}

func (m *move) GetPower() int {
    return m.Pwr
}

func (m *move) GetAccuracy() int {
    return m.Accy
}

func (m *move) GetPPs() int {
    return m.PPs
}
