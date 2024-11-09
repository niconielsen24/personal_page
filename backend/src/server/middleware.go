package server

import (
	"errors"

	"github.com/google/uuid"
)

var (
  ErrInvalidUUID = errors.New("Invalid uuid")
  ErrUuidDoesNotExist = errors.New("Uuid does not exist")
  ErrInvalidPosition = errors.New("Position is invalid")
)

func ValidateMoveRequest(ts *TicTacToeServer, mr MoveRequest) error {
	if mr.Id == uuid.Nil {
		return ErrInvalidUUID
	}
	res := false
	for i := range ts.games {
		res = res && mr.Id == ts.games[i].ID
	}
	if !res {
	  return ErrUuidDoesNotExist 
  }
	if mr.Position.OutOfBounds() {
		return ErrInvalidPosition
	}
	return nil
}
