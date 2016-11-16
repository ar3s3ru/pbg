package mem

import "errors"

var (
	ErrInvalidDataSourceFile = errors.New(`Using "" as Pokèmon file is not allowed`)

	ErrInvalidMoveDataset    = errors.New("Invalid Move dataset used")
	ErrInvalidPokèmonDataset = errors.New("Invalid Pokèmon dataset used")

	// Session errors
	ErrInvalidTrainerValue = errors.New("Invalid Trainer reference specified")
	ErrInvalidTokenValue   = errors.New("Invalid Session token used")

	// PokèmonTeam factory errors
	ErrInvalidReferenceValue = errors.New("Invalid Pokèmon reference value inserted")
	ErrInvalidLevelValue     = errors.New("Invalid Pokèmon Level value inserted")
	ErrInvalidIVsValue       = errors.New("Invalid Pokèmon IVs values inserted")
	ErrInvalidEVsValue       = errors.New("Invalid Pokèmon EVs values inserted")
	ErrInvalidMovesValue     = errors.New("Invalid Pokèmon Moves, can't be all null")
)
