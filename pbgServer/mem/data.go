package mem

import (
    "github.com/ar3s3ru/PokemonBattleGo/pbgServer"
    "sync"
    "errors"
    "gopkg.in/mgo.v2/bson"
    "time"
)

type (
    memData struct {
        pokèmons   map[bson.ObjectId]pbgServer.Pokèmon
        trainers   map[bson.ObjectId]pbgServer.Trainer
        // NOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO!
        // EXTREME BOTTLENECK HERE!
        trainMutex sync.Mutex
    }

    memDataBuilder struct {
        pokèmonFile  string
        trainersFile string
    }

    DataBuilder interface {
        UsePokèmonFile(path string)  DataBuilder
        UseTrainersFile(path string) DataBuilder

        Build() pbgServer.IDataMechanism
    }
)

var (
    ErrPokèmonNotFound    = errors.New("Pokèmon not found")
    ErrTrainerNotFound    = errors.New("Trainer not found")
    ErrIllegalTrainer     = errors.New("Trainer object is nil")
    ErrInvalidTrainerName = errors.New("Invalid Trainer name used")
)

func NewDataBuilder() DataBuilder {
    return &memDataBuilder{}
}

func (builder *memDataBuilder) UsePokèmonFile(path string) DataBuilder {
    builder.pokèmonFile = path
    return builder
}

func (builder *memDataBuilder) UseTrainersFile(path string) DataBuilder {
    builder.trainersFile = path
    return builder
}

func (builder *memDataBuilder) Build() pbgServer.IDataMechanism {
    if builder.pokèmonFile == "" {
        panic("Using \"\" as Pokèmon file is not allowed")
    }

    return &memData{
        pokèmons: make(map[bson.ObjectId]pbgServer.Pokèmon),
        trainers: make(map[bson.ObjectId]pbgServer.Trainer),
    }
}

func (data *memData) AddTrainer(trainer pbgServer.Trainer) (bson.ObjectId, error) {
    if trainer == nil {
        return "", ErrIllegalTrainer
    }

    data.trainMutex.Lock()
    defer data.trainMutex.Unlock()

    id := bson.NewObjectIdWithTime(time.Now())
    data.trainers[id] = trainer
    return id, nil
}

func (data *memData) RemoveTrainer(id bson.ObjectId) error {
    data.trainMutex.Lock()
    defer data.trainMutex.Unlock()

    if trainer := data.trainers[id]; trainer == nil {
        return ErrTrainerNotFound
    } else {
        delete(data.trainers, id)
        return nil
    }
}

func (data *memData) GetPokèmonById(id bson.ObjectId) (pbgServer.Pokèmon, error) {
    if pokèmon := data.pokèmons[id]; pokèmon == nil {
        return nil, ErrPokèmonNotFound
    } else {
        return pokèmon, nil
    }
}

func (data *memData) GetTrainerById(id bson.ObjectId) (pbgServer.Trainer, error) {
    data.trainMutex.Lock()
    defer data.trainMutex.Unlock()

    if trainer := data.trainers[id]; trainer == nil {
        return nil, ErrTrainerNotFound
    } else {
        return trainer, nil
    }
}

func (data *memData) GetTrainerByName(name string) (pbgServer.Trainer, error) {
    if name == "" {
        return nil, ErrInvalidTrainerName
    }

    data.trainMutex.Lock()
    defer data.trainMutex.Unlock()

    for _, trainer := range data.trainers {
        if trainer.GetName() == name {
            return trainer, nil
        }
    }

    return nil, ErrTrainerNotFound
}
