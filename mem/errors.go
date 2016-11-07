package mem

import "errors"

var (
    ErrInvalidMoveType        = errors.New("Invalid type used as Move interface")
    ErrInvalidTrainerType     = errors.New("Invalid type used as Trainer interface")
    ErrInvalidPokèmonType     = errors.New("Invalid type used as Pokèmon interface")
    ErrInvalidPokèmonTeamType = errors.New("Invalid type used as PokèmonTeam interface")
    ErrInvalidSessionType     = errors.New("Invalid type used as Session interface")

    ErrInvalidDataSourceFile = errors.New(`Using "" as Pokèmon file is not allowed`)

    ErrTrainerConversion  = errors.New("Cannot convert interface{} to Trainer interface")
    ErrObjectIdConversion = errors.New("Cannot convert interface{} to bson.ObjectId type")

    // PokèmonTeam factory errors
    ErrInvalidReferenceValue = errors.New("Invalid Pokèmon reference value inserted")
    ErrInvalidLevelValue     = errors.New("Invalid Pokèmon Level value inserted")
    ErrInvalidIVsValue       = errors.New("Invalid Pokèmon IVs values inserted")
    ErrInvalidEVsValue       = errors.New("Invalid Pokèmon EVs values inserted")
    ErrInvalidMovesValue     = errors.New("Invalid Pokèmon Moves, can't be all null")
)
