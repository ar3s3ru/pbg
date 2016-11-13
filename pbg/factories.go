package pbg

type (
    // Opzioni funzionali per i PokèmonFactory
    PokèmonFactoryOption func(Pokèmon) error
    // Factory method per gli oggetti di tipo Pokèmon
    PokèmonFactory       func(...PokèmonFactoryOption) (Pokèmon, error)

    // Opzioni funzionali per i MoveFactory
    MoveFactoryOption func(Move) error
    // Factory method per gli oggetti di tipo Move
    MoveFactory       func(...MoveFactoryOption) (Move, error)

    // Opzioni funzionali per i TrainerFactory
    TrainerFactoryOption func(Trainer) error
    // Factory method per gli oggetti di tipo Trainer
    TrainerFactory       func(...TrainerFactoryOption) (Trainer, error)

    // Opzioni funzionali per i TeamFactory
    TeamFactoryOption func(PokèmonTeam) error
    // Factory method per gli oggetti di tipo PokèmonTeam
    TeamFactory       func(...TeamFactoryOption) (PokèmonTeam, error)

    // Opzioni funzionali per i SessionFactory
    SessionFactoryOption func(Session) error
    // Factory method per gli oggetti di tipo Session
    SessionFactory       func(...SessionFactoryOption) (Session, error)
)
