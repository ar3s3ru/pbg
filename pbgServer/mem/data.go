package mem

import (
    "github.com/ar3s3ru/PokemonBattleGo/pbgServer"
    "sync"
    "errors"
    "gopkg.in/mgo.v2/bson"
    "time"
    "encoding/json"
    "io/ioutil"
)

type (
    DataBuilder interface {
        UsePokèmonFile(path string)  DataBuilder
        UseTrainersFile(path string) DataBuilder

        Build() pbgServer.IDataMechanism
    }

    memData struct {
        pokèdx   *pokèdex
        trainers map[bson.ObjectId]pbgServer.Trainer
        // NOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO!
        // EXTREME BOTTLENECK HERE!
        trainMutex sync.Mutex
    }

    memDataBuilder struct {
        pokèmonFile  string
        trainersFile string
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

    var pkms pokèdex

    // TODO: pokèmon file unmarshalling
    if file, err := ioutil.ReadFile(builder.pokèmonFile); err != nil {
        panic(err)
    } else {
        pkms = pokèdex{}
        if err := json.Unmarshal(file, &pkms); err != nil {
            panic(err)
        }
    }

    return &memData{
        pokèdx:   &pkms,
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

func (data *memData) GetPokèmonById(id int) (pbgServer.Pokèmon, error) {
    if id <= 0 || id > len(data.pokèdx.List) {
        return nil, ErrPokèmonNotFound
    } else {
        return &data.pokèdx.List[(id - 1)], nil
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
