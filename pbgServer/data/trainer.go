package data

type (
    PokemonSquad struct {
        Pokemon *Pokemon
        Nature  uint8
        Moves   [4]*Move
        Level   uint8
        EVs     [6]uint8
        IVs     [6]uint8
    }

    Trainer struct {
        Name         string
        PasswordHash string
        SignUpDate   uint64

        Type TrainerType
    }
)

// PokemonSquad functions
func (pkm *PokemonSquad) GetStatValue(stat Statistic) uint8 {
    switch basePkm := pkm.Pokemon; stat {
    case PS:
        return ((pkm.IVs[PS] + (2 * basePkm.BaseStat[PS]) + pkm.EVs[PS]) * ((pkm.Level) / 100)) + 10 + pkm.Level
    case Atk:
    case Def:
    case AtkSp:
    case DefSp:
    case Vel:
        return (((pkm.IVs[stat] + (2 * basePkm.BaseStat[stat]) + pkm.EVs[stat]) * ((pkm.Level) / 100)) + 5) * pkm.Nature
    }

    return 0
}