package mem

import "github.com/ar3s3ru/PokemonBattleGo/pbgServer"

type (
    move struct {
        Name string             `json:"name"`
        Typ  pbgServer.Type     `json:"type"`
        Cat  pbgServer.Category `json:"category"`
        PPs  int                `json:"pps"`
        Pwr  int                `json:"power"`
        Accy int                `json:"accuracy"`
        Prio int                `json:"priority"`
    }
)

func convertLtoML(moves []move) []pbgServer.Move {
    if moves == nil {
        panic("Must use a valid move list!")
    }

    list := make([]pbgServer.Move, len(moves), len(moves))
    for i := range moves {
        list[i] = pbgServer.Move(&moves[i])
    }

    return list
}

func (m *move) GetName() string {
    return m.Name
}

func (m *move) GetType() pbgServer.Type {
    return m.Typ
}

func (m *move) GetCategory() pbgServer.Category {
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
