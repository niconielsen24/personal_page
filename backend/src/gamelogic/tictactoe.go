package gamelogic

import (
	"errors"
	"github.com/google/uuid"
)

var (
	ErrNotPlayerTurn = errors.New("not this players turn")
)

type Tictactoe struct {
	Board  Board
	Turn   Player
	ID     uuid.UUID
	Winner Player
}

func (t *Tictactoe) InitTictactoe() error {
	t.Board.InitBoard()
	t.Turn = PlayerX
  t.Winner = Empty
	uuid, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	t.ID = uuid
	return nil
}

func (t *Tictactoe) GameOver() (bool, Player) {
	if t.Board.BoardIsFull() {
		return true, Empty
	}
	winner := t.Board.CheckWinner()
	if winner == Empty {
		return false, Empty
	}
  t.Winner = winner
	return true, winner
}

func (t *Tictactoe) MakeMove(p Player, pos Position) error {
	if p != t.Turn {
		return ErrNotPlayerTurn
	}
	err := t.Board.TurnTile(p, pos)
	if err != nil {
		return err
	}

	if t.Turn == PlayerX {
		t.Turn = PlayerO
	} else {
		t.Turn = PlayerX
	}

	return nil
}
