package mem

import (
	"log"

	"github.com/ar3s3ru/pbg"
	"gopkg.in/mgo.v2/bson"
)

type (
	trainerRequest     func(map[bson.ObjectId]pbg.Trainer)
	TrainerDBComponent struct {
		trainers map[bson.ObjectId]pbg.Trainer
		requests chan trainerRequest
		factory  TrainerFactory
		logger   *log.Logger
	}
)

func NewTrainerDBComponent(options ...TrainerDBComponentOption) pbg.TrainerDBComponent {
	trainerdb := &TrainerDBComponent{
		trainers: make(map[bson.ObjectId]pbg.Trainer),
		requests: make(chan trainerRequest),
		factory:  NewTrainer,
	}

	for _, option := range options {
		if err := option(trainerdb); err != nil {
			panic(err)
		}
	}

	go trainerdb.requestDispatcher()
	return trainerdb
}

func (tdb *TrainerDBComponent) requestDispatcher() {
	for request := range tdb.requests {
		request(tdb.trainers)
	}
}

func (tdb *TrainerDBComponent) Supply() pbg.TrainerDBInterface {
	return tdb
}

func (tdb *TrainerDBComponent) Retrieve(_ pbg.TrainerDBInterface) {
	// Nothing for now
}
