package pbg

import "errors"

var (
    // Server building
    ErrUnspecifiedConfig = errors.New("Configuration not set in the builder")
    ErrUnspecifiedDataM  = errors.New("Data mechanism not set in the builder")
    ErrUnspecifiedSessM  = errors.New("Session mechanism not set in the builder")
    ErrUnspecifiedAuthM  = errors.New("Authorization mechanism not set in the builder")
)
