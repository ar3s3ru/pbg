package pbgServer

type (
    ISessionsMechanism interface {
        AddSession(trainer Trainer) string
        GetSession(token string)    Session
        RemoveSession(token string)
    }

    Session struct {
        // TODO
    }
)
