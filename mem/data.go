package mem

import (
    "sync"
    "time"
    "io/ioutil"
    "encoding/json"

    "golang.org/x/crypto/bcrypt"
    "gopkg.in/mgo.v2/bson"

    "github.com/ar3s3ru/PokemonBattleGo/pbg"
)

type (
    DataBuilder interface {
        WithSourceFile(string)   DataBuilder
        WithTrainersFile(string) DataBuilder

        Build() pbg.DataMechanism
    }

    dataBuilder     func(string, string) pbg.DataMechanism
    dataLockHandler func()               (interface{}, error)

    data struct {
        // Pokèmon and Move "database"
        pokèmons []pbg.Pokèmon
        moves    []pbg.Move
        // Trainer map
        trainers map[bson.ObjectId]pbg.Trainer
        mutex    sync.RWMutex
        // Factories
        teamFactory    pbg.TeamFactory
        moveFactory    pbg.MoveFactory
        pokèmonFactory pbg.PokèmonFactory
        trainerFactory pbg.TrainerFactory
    }
)

func NewDataBuilder() DataBuilder {
    return dataBuilder(func(source string, trainers string) pbg.DataMechanism {
        if source == "" {
            panic(ErrInvalidDataSourceFile)
        } else if file, err := ioutil.ReadFile(source); err != nil {
            panic(err)
        } else {
            source := sourceFile{}
            if err := json.Unmarshal(file, &source); err != nil {
                panic(err)
            }

            _pokèmonFactory := NewPokèmonFactory()
            _moveFactory    := NewMoveFactory()

            return &data{
                pokèmons: convertLtoPL(source.PList),
                moves:    convertLtoML(source.MList),
                // Trainer map
                trainers: make(map[bson.ObjectId]pbg.Trainer),
                // Factories
                teamFactory:    NewTeamFactory(),
                moveFactory:    _moveFactory,
                pokèmonFactory: _pokèmonFactory,
                trainerFactory: NewTrainerFactory(),
            }
        }
    })
}

func (db dataBuilder) WithSourceFile(path string) DataBuilder {
    return dataBuilder(func(_ string, trainers string) pbg.DataMechanism {
        return db(path, trainers)
    })
}

func (db dataBuilder) WithTrainersFile(path string) DataBuilder {
    return dataBuilder(func(source string, _ string) pbg.DataMechanism {
        return db(source, path)
    })
}

func (db dataBuilder) Build() pbg.DataMechanism {
    return db("", "")
}

func (d *data) GetPokèmon(id int) (pbg.Pokèmon, error) {
    if inRange(id, len(d.pokèmons)) {
        return nil, pbg.ErrPokèmonNotFound
    } else {
        return d.pokèmons[id - 1], nil
    }
}

func (d *data) GetMove(id int) (pbg.Move, error) {
    if inRange(id, len(d.moves)) {
        return nil, pbg.ErrMoveNotFound
    } else {
        return d.moves[id - 1], nil
    }
}

func (d *data) ListPokèmon() []pbg.Pokèmon {
    return d.pokèmons
}

func (d *data) ListMoves() []pbg.Move {
    return d.moves
}

func (d *data) GetTrainerById(id bson.ObjectId) (pbg.Trainer, error) {
    return d.readDataLockedRegion(
        d.handlerTrainerById(id),
    )
}

func (d *data) GetTrainerByName(name string) (pbg.Trainer, error) {
    return d.readDataLockedRegion(
        d.handlerTrainerByName(name),
    )
}

func (d *data) AddTrainer(user, pass string) (bson.ObjectId, error) {
    if _, err := d.GetTrainerByName(user); err == nil {
        return "", pbg.ErrTrainerAlreadyExists
    } else if pwd, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost); err != nil {
        return "", pbg.ErrPasswordSalting
    } else {
        return d.writeDataLockedRegion(
            d.handlerAddTrainer(
                d.GetTrainerFactory().Create(
                    WithTrainerName(user),
                    WithTrainerPassword(pwd),
                ),
            ),
        )
    }
}

func (d *data) RemoveTrainer(id bson.ObjectId) error {
    _, err := d.writeDataLockedRegion(
        d.handlerRemoveTrainer(id),
    )

    return err
}

func (d *data) GetPokèmonFactory() pbg.PokèmonFactory {
    return d.pokèmonFactory
}

func (d *data) GetMoveFactory() pbg.MoveFactory {
    return d.moveFactory
}

func (d *data) GetTrainerFactory() pbg.TrainerFactory {
    return d.trainerFactory
}

func (d *data) GetTeamFactory() pbg.TeamFactory {
    return d.teamFactory
}

func (d *data) readDataLockedRegion(handler dataLockHandler) (pbg.Trainer, error) {
    d.mutex.RLock()
    defer d.mutex.RUnlock()

    tr, err := handler()
    if convTr, ok := tr.(pbg.Trainer); ok {
        return convTr, err
    }

    // Can't convert to Trainer type, error
    return nil, ErrTrainerConversion
}

func (d *data) writeDataLockedRegion(handler dataLockHandler) (bson.ObjectId, error) {
    d.mutex.Lock()
    defer d.mutex.Unlock()

    id, err := handler()
    if convId, ok := id.(bson.ObjectId); ok {
        return convId, err
    }

    // Can't convert to ObjectId type, error
    return "", ErrObjectIdConversion
}

func (d *data) handlerAddTrainer(trainer pbg.Trainer, err error) dataLockHandler {
    return func() (interface{}, error) {
        // Errors from TrainerFactory creation
        if err != nil {
            return nil, err
        }

        id := bson.NewObjectIdWithTime(time.Now())
        d.trainers[id] = trainer

        return id, nil
    }
}

func (d *data) handlerRemoveTrainer(id bson.ObjectId) dataLockHandler {
    return func() (interface{}, error) {
        if _, ok := d.trainers[id]; !ok {
            return nil, pbg.ErrTrainerNotFound
        } else {
            delete(d.trainers, id)
            return nil, nil
        }
    }
}

func (d *data) handlerTrainerById(id bson.ObjectId) dataLockHandler {
    return func() (interface{}, error) {
        if trainer, ok := d.trainers[id]; !ok {
            return nil, pbg.ErrTrainerNotFound
        } else {
            return trainer, nil
        }
    }
}

func (d *data) handlerTrainerByName(name string) dataLockHandler {
    return func() (interface{}, error) {
        for _, trainer := range d.trainers {
            if trainer.GetName() == name {
                return trainer, nil
            }
        }

        return nil, pbg.ErrTrainerNotFound
    }
}
