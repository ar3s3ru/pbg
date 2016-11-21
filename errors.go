package pbg

import "errors"

var (
	ErrMoveNotFound    = errors.New("move not found")
	ErrPokèmonNotFound = errors.New("pokèmon not found")
	ErrTrainerNotFound = errors.New("trainer not found")
	ErrSessionNotFound = errors.New("session not found")
	ErrSessionExpired  = errors.New("session has expired")

	// Trainer creation errors
	ErrTrainerAlreadyExists = errors.New("trainer already exists")
	ErrInvalidPasswordUsed  = errors.New("username or password are invalid")

	// Authorization errors
	ErrPasswordSalting            = errors.New("password chosen cannot be used, try a different one")
	ErrInvalidAuthorizationHeader = errors.New("invalid Authorization header (not Basic)")

	// Server errors
	ErrInvalidFastHTTPServer = errors.New("invalid fasthttp Server used, must be not nil")
	ErrInvalidServerType 	 = errors.New("invalid Server type used")
	ErrInvalidHTTPPort   	 = errors.New("invalid HTTP port value used")
	ErrInvalidLogger     	 = errors.New("invalid logger used")

	ErrInvalidAPIEndpoint      = errors.New(`invalid API endpoint value (must be at least 2 character and start with an "/"`)
	ErrInvalidAPIResponser     = errors.New("invalid APIResponser value used in the builder")
	ErrInvalidSessionComponent = errors.New("invalid SessionComponent object used")

	ErrInvalidMoveDBComponent    = errors.New("invalid MoveDBComponent value passed to the Server")
	ErrInvalidPokèmonDBComponent = errors.New("invalid PokèmonDBComponent value passed to the Server")
	ErrInvalidTrainerDBComponent = errors.New("invalid TrainerDBComponent value passed to the Server")
)
