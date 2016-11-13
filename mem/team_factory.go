package mem

import "github.com/ar3s3ru/PokemonBattleGo/pbg"

type (
    teamOption func(*pokèmonTeam) error
)

func NewPokèmonTeam(options ...pbg.TeamFactoryOption) (pbg.PokèmonTeam, error) {
    pokèmon := &pokèmonTeam{}

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
        if pkmn == nil {
            return ErrInvalidReferenceValue
        }

        pokèmon.Pokèmon = pkmn
        return nil
    })
}

func WithPokèmonMoves(move1, move2, move3, move4 pbg.Move) pbg.TeamFactoryOption {
    return adaptTeamOption(func(pokèmon *pokèmonTeam) error {
        if move1 == nil && move2 == nil && move3 == nil && move4 == nil {
            return ErrInvalidMovesValue
        }

        pokèmon.Movs[0] = move1
        pokèmon.Movs[1] = move2
        pokèmon.Movs[2] = move3
        pokèmon.Movs[3] = move4

        return nil
    })
}

func WithPokèmonLevel(level int) pbg.TeamFactoryOption {
    return adaptTeamOption(func(pokèmon *pokèmonTeam) error {
        if !checkLevel(level) {
            return ErrInvalidLevelValue
        }

        pokèmon.Levl = level
        return nil
    })
}

func WithPokèmonIVs(ivs [6]int) pbg.TeamFactoryOption {
    return adaptTeamOption(func(pokèmon *pokèmonTeam) error {
        if !checkIvs(ivs) {
            return ErrInvalidIVsValue
        }

        for i := range pokèmon.Ivs {
            pokèmon.Ivs[i] = ivs[i]
        }

        return nil
    })
}

func WithPokèmonEVs(evs [6]int) pbg.TeamFactoryOption {
    return adaptTeamOption(func(pokèmon *pokèmonTeam) error {
        if !checkEvs(evs) {
            return ErrInvalidEVsValue
        }

        for i := range pokèmon.Evs {
            pokèmon.Evs[i] = evs[i]
        }

        return nil
    })
}
