package pbg

import "gopkg.in/mgo.v2/bson"

type (
    // Interfaccia per accedere al logger di un particolare componente del framework
    // Esegue un log.Logger.Println() sotto il tappeto (in genere)
    Logger interface {
        // Scrive le interfacce passate sul logger del componente
        Log(...interface{})
    }

    // Componente software che identifica il DB dei modelli Pokèmon
    // Permette di fornire un'interfaccia d'accesso al DB e di richiederla indietro
    // (utile nel caso di ObjectPools)
    PokèmonDBComponent interface {
        Supply() PokèmonDBInterface
        Retrieve(PokèmonDBInterface)
    }

    PokèmonDBComponentOption func(PokèmonDBComponent) error

    // Interfaccia software che permette di eseguire operazioni CRUD
    // sul DB dei Pokèmon
    PokèmonDBInterface interface {
        Logger

        GetPokèmon(id int) (Pokèmon, error)
        GetPokèmons()      []Pokèmon
    }

    // Componente software che identifica il DB dei modelli Mosse
    // Permette di fornire un'interfaccia d'accesso al DB e di richiederla indietro
    // (utile nel caso di ObjectPools)
    MoveDBComponent interface {
        Supply() MoveDBInterface
        Retrieve(MoveDBInterface)
    }

    MoveDBComponentOption func(MoveDBComponent) error

    // Interfaccia software che permette di eseguire operazioni CRUD
    // sul DB delle Mosse
    MoveDBInterface interface {
        Logger

        GetMove(id int) (Move, error)
        GetMoves()      []Move
    }

    // Componente software che identifica il DB dei modelli Allenatori
    // Permette di fornire un'interfaccia d'accesso al DB e di richiederla indietro
    // (utile nel caso di ObjectPools)
    TrainerDBComponent interface {
        Supply() TrainerDBInterface
        Retrieve(TrainerDBInterface)
    }

    TrainerDBComponentOption func(TrainerDBComponent) error

    // Interfaccia software che permette di eseguire operazioni CRUD
    // sul DB degli Allenatori
    TrainerDBInterface interface {
        Logger

        AddTrainer(options ...TrainerFactoryOption) (bson.ObjectId, error)
        GetTrainerByName(name string)               (Trainer, error)
        GetTrainerById(id bson.ObjectId)            (Trainer, error)
        DeleteTrainer(id bson.ObjectId)             error
    }

    // Componente software che identifica il DB delle sessioni del server
    // Permette di fornire e richiedere indietro un'interfaccia di accesso al DB
    // (utile nel caso di ObjectPools)
    SessionComponent interface {
        // Fornisce un'interfaccia al DB delle sessioni
        Supply() SessionInterface
        // Richiede indietro l'interfaccia fornita in precedenza
        Retrieve(SessionInterface)

        // Elimina tutte le sessioni scadute dal DB
        Purge()
    }

    SessionComponentOption func(SessionComponent) error

    // Interfaccia software che permette di eseguire determinate operazioni
    // sul DB delle sessioni del server.
    SessionInterface interface {
        Logger

        // Operazioni CRUD -------------------------
        // -----------------------------------------

        AddSession(options ...SessionFactoryOption) (Session, error)
        GetSession(token string)                    (Session, error)
        DeleteSession(token string)                 error
    }
)
