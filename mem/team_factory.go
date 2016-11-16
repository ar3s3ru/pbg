package mem

import "github.com/ar3s3ru/pbg"

type (
	PokèmonTeamOption func(*PokèmonTeam) error
)

func NewPokèmonTeam(options ...PokèmonTeamOption) (pbg.PokèmonTeam, error) {
	pokèmon := &PokèmonTeam{}

	for _, option := range options {
		if err := option(pokèmon); err != nil {
			return nil, err
		}
	}

	return pokèmon, nil
}

func WithPokèmonReference(pkmn pbg.Pokèmon) PokèmonTeamOption {
	return func(pokèmon *PokèmonTeam) error {
		if pkmn == nil {
			return ErrInvalidReferenceValue
		}

		pokèmon.Pokèmon = pkmn
		return nil
	}
}

func WithPokèmonMoves(move1, move2, move3, move4 pbg.Move) PokèmonTeamOption {
	return func(pokèmon *PokèmonTeam) error {
		if move1 == nil && move2 == nil && move3 == nil && move4 == nil {
			return ErrInvalidMovesValue
		}

		pokèmon.moves[0] = move1
		pokèmon.moves[1] = move2
		pokèmon.moves[2] = move3
		pokèmon.moves[3] = move4

		return nil
	}
}

func WithPokèmonLevel(level int) PokèmonTeamOption {
	return func(pokèmon *PokèmonTeam) error {
		if !checkLevel(level) {
			return ErrInvalidLevelValue
		}

		pokèmon.level = level
		return nil
	}
}

func WithPokèmonIVs(ivs [6]int) PokèmonTeamOption {
	return func(pokèmon *PokèmonTeam) error {
		if !checkIvs(ivs) {
			return ErrInvalidIVsValue
		}

		for i := range pokèmon.ivs {
			pokèmon.ivs[i] = ivs[i]
		}

		return nil
	}
}

func WithPokèmonEVs(evs [6]int) PokèmonTeamOption {
	return func(pokèmon *PokèmonTeam) error {
		if !checkEvs(evs) {
			return ErrInvalidEVsValue
		}

		for i := range pokèmon.evs {
			pokèmon.evs[i] = evs[i]
		}

		return nil
	}
}
