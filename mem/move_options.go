package mem

import (
	"github.com/ar3s3ru/pbg"
	"log"
)

type (
	MoveDBComponentOption func(*MoveDBComponent) error
)

func WithMovesFile(moveFile string) MoveDBComponentOption {
	return func(mdb *MoveDBComponent) error {
		// TODO: finish this
		return nil
	}
}

func WithMoves(moves []pbg.Move) MoveDBComponentOption {
	return func(mdb *MoveDBComponent) error {
		if moves == nil {
			return ErrInvalidMoveDataset
		}

		mdb.moves = moves
		return nil
	}
}

func WithMoveDBLogger(logger *log.Logger) MoveDBComponentOption {
	return func(mdb *MoveDBComponent) error {
		mdb.logger = logger
		return nil
	}
}
