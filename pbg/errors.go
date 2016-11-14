package pbg

import "errors"

var (
    ErrMoveNotFound    = errors.New("Move not found")
    ErrPokèmonNotFound = errors.New("Pokèmon not found")
    ErrTrainerNotFound = errors.New("Trainer not found")
    ErrSessionNotFound = errors.New("Session not found")
    ErrSessionExpired  = errors.New("Session has expired")

    // Trainer creation errors
    ErrTrainerAlreadyExists = errors.New("Trainer already exists")
    ErrInvalidPasswordUsed  = errors.New("Username or password are invalid")

    // Authorization errors
    ErrPasswordSalting            = errors.New("Password chosen cannot be used, try a different one")
    ErrInvalidAuthorizationHeader = errors.New("Invalid Authorization header (not Basic)")

    // Server errors
    ErrInvalidServerType = errors.New("Invalid Server type used")
    ErrInvalidHTTPPort   = errors.New("Invalid HTTP port value used")
    ErrInvalidLogger     = errors.New("Invalid logger used")

    ErrInvalidAPIEndpoint      = errors.New(`Invalid API endpoint value (must be at least 2 character and start with an "/"`)
    ErrInvalidAPIResponser     = errors.New("Invalid APIResponser value used in the builder")
    ErrInvalidSessionComponent = errors.New("Invalid SessionComponent object used")

    ErrInvalidMoveDBComponent    = errors.New("Invalid MoveDBComponent value passed to the Server")
    ErrInvalidPokèmonDBComponent = errors.New("Invalid PokèmonDBComponent value passed to the Server")
    ErrInvalidTrainerDBComponent = errors.New("Invalid TrainerDBComponent value passed to the Server")
)
