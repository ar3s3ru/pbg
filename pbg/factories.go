package pbg

type (
    PokèmonFactory interface {
        Create()                 Pokèmon
        CreateSlice(...Pokèmon) []Pokèmon
    }

    MoveFactory interface {
        Create()              Move
        CreateSlice(...Move) []Move
    }

    TrainerFactory interface {
        Create() Trainer
    }

    TeamFactory interface {
        Create() PokèmonTeam
    }

    SessionFactory interface {
        Create() Session
    }
)
