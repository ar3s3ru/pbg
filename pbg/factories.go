package pbg

type (
    PokèmonFactoryOption func(Pokèmon) error
    PokèmonFactory interface {
        Create(...PokèmonFactoryOption) (Pokèmon, error)
        CreateSlice(...Pokèmon)         ([]Pokèmon, error)
    }

    MoveFactoryOption func(Move) error
    MoveFactory interface {
        Create(...MoveFactoryOption) (Move, error)
        CreateSlice(...Move)         ([]Move, error)
    }

    TrainerFactoryOption func(Trainer) error
    TrainerFactory interface {
        Create(...TrainerFactoryOption) (Trainer, error)
    }

    TeamFactoryOption func(PokèmonTeam) error
    TeamFactory interface {
        Create(...TeamFactoryOption) (PokèmonTeam, error)
    }

    SessionFactoryOption func(Session) error
    SessionFactory interface {
        Create(...SessionFactoryOption) (Session, error)
    }
)
