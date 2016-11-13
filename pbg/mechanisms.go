package pbg

import "gopkg.in/mgo.v2/bson"

type (
    // Interfaccia per accedere al logger di un particolare componente del framework
    // Esegue un log.Logger.Println() sotto il tappeto (in genere)
    Logger interface {
        // Scrive le interfacce passate sul logger del componente
        Log(...interface{})
    }

    // Componente software che identifica il DB dei modelli del software
    // Permette di fornire un'interfaccia d'accesso al DB e di richiederla indietro
    // (utile nel caso di ObjectPools)
    DataComponent interface {
        // Fornisce un'interfaccia al DB dei modelli
        Supply() DataInterface
        // Richiede l'interfaccia fornita indietro
        Retrieve(DataInterface)
    }

    // Interfaccia software che permette di eseguire determinate operazioni
    // sul DB dei modelli del software.
    DataInterface interface {
        Logger

        // Operazioni CRUD ------------------------------------
        // ----------------------------------------------------

        GetPokèmon(id int)                    (Pokèmon, error)
        //AddPokèmon(pkm Pokèmon)               error
        //DeletePokèmon(id int)                 error
        //UpdatePokèmon(id int, newPkm Pokèmon) error

        GetMove(id int)                (Move, error)
        //AddMove(mv Move)               error
        //DeleteMove(id int)             error
        //UpdateMove(id int, newMv Move) error

        GetTrainerByName(name string)    (Trainer, error)
        GetTrainerById(id bson.ObjectId) (Trainer, error)
        AddTrainer(user, pass string)    (bson.ObjectId, error)
        DeleteTrainer(id bson.ObjectId)  error
        //UpdateTrainer(id bson.ObjectId, newTr Trainer) error
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

    // Interfaccia software che permette di eseguire determinate operazioni
    // sul DB delle sessioni del server.
    SessionInterface interface {
        Logger

        // Operazioni CRUD -------------------------
        // -----------------------------------------

        GetSession(token string)    (Session, error)
        AddSession(trainer Trainer) (Session, error)
        DeleteSession(token string) error
    }
)
