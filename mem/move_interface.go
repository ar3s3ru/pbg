package mem

import "github.com/ar3s3ru/pbg"

func (mdb *MoveDBComponent) Log(v ...interface{}) {
	if mdb.logger != nil {
		mdb.logger.Println(v...)
	}
}

func (mdb *MoveDBComponent) GetMove(id int) (pbg.Move, error) {
	if inRange(id, len(mdb.moves)) {
		return mdb.moves[id-1], nil
	} else {
		return nil, pbg.ErrMoveNotFound
	}
}

func (mdb *MoveDBComponent) GetMoves() []pbg.Move {
	return mdb.moves
}
