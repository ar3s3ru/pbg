package mem

import "github.com/ar3s3ru/PokemonBattleGo/pbg"

type teamOption  func(*pokèmonTeam) error
type TeamFactory func()             pbg.PokèmonTeam

func NewTeamFactory() pbg.TeamFactory {
    return TeamFactory(func() pbg.PokèmonTeam {
        return &pokèmonTeam{
            Moves: [4]pbg.Move{nil, nil, nil, nil},
            IVs:   [6]int{0, 0, 0, 0, 0, 0},
            EVs:   [6]int{0, 0, 0, 0, 0, 0},
        }
    })
}

func (tf TeamFactory) Create(options ...pbg.TeamFactoryOption) (pbg.PokèmonTeam, error) {
    pokèmon := tf()
    for _, option := range options {
        if err := option(pokèmon); err != nil {
            return nil, err
        }
    }

    return pokèmon, nil
}

func adaptTeamOption(option teamOption) pbg.TeamFactoryOption {
    return func(pTeam pbg.PokèmonTeam) error {
        switch converted := pTeam.(type) {
        case *pokèmonTeam:
            return option(converted)
        default:
            return ErrInvalidPokèmonTeamType
        }


    }
}

func WithPokèmonReference(pkmn pbg.Pokèmon) pbg.TeamFactoryOption {
    return adaptTeamOption(func(pokèmon *pokèmonTeam) error {
        pokèmon.Pokèmon = pkmn
        return nil
    })
}

func WithPokèmonMoves(move1, move2, move3, move4 pbg.Move) pbg.TeamFactoryOption {
    return adaptTeamOption(func(pokèmon *pokèmonTeam) error {
        pokèmon.Moves[0] = move1
        pokèmon.Moves[1] = move2
        pokèmon.Moves[2] = move3
        pokèmon.Moves[3] = move4

        return nil
    })
}

func WithPokèmonLevel(level int) pbg.TeamFactoryOption {
    return adaptTeamOption(func(pokèmon *pokèmonTeam) error {
        pokèmon.Level = level
        return nil
    })
}

func WithPokèmonIVs(ivs [6]int) pbg.TeamFactoryOption {
    return adaptTeamOption(func(pokèmon *pokèmonTeam) error {
        for i := range pokèmon.IVs {
            pokèmon.IVs[i] = ivs[i]
        }

        return nil
    })
}

func WithPokèmonEVs(evs [6]int) pbg.TeamFactoryOption {
    return adaptTeamOption(func(pokèmon *pokèmonTeam) error {
        for i := range pokèmon.EVs {
            pokèmon.EVs[i] = evs[i]
        }

        return nil
    })
}
