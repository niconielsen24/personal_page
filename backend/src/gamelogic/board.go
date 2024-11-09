package gamelogic

import (
	"errors"
	"fmt"
)

type Player rune

var (
	ErrPosistionOutOfBounds = errors.New("position is out of bounds")
	ErrTileAlreadyUsed      = errors.New("tile is already used")
)

const (
	PlayerX Player = 'X'
	PlayerO Player = 'O'
	Empty   Player = ' '
)

type Position struct {
	x uint8
	y uint8
}

type Board struct {
	Tiles [3][3]Player `json:"tiles"`
}

func (p *Position) OutOfBounds() bool {
  return IsOutOfBounds(p.x, 0, 2) || IsOutOfBounds(p.y, 0, 2)
}

func (b *Board) InitBoard() {
	for i := range b.Tiles {
		for j := range b.Tiles[i] {
			b.Tiles[i][j] = Empty
		}
	}
}

func (b *Board) BoardIsFull() bool {
	for _, row := range b.Tiles {
		for _, tile := range row {
			if tile == Empty {
				return false
			}
		}
	}
	return true
}

func (b *Board) CheckWinner() Player {
	for i := range b.Tiles {
		if b.Tiles[i][0] == PlayerX && b.Tiles[i][1] == PlayerX && b.Tiles[i][2] == PlayerX {
			return PlayerX
		}
		if b.Tiles[i][0] == PlayerO && b.Tiles[i][1] == PlayerO && b.Tiles[i][2] == PlayerO {
			return PlayerO
		}
	}
	for i := range b.Tiles {
		if b.Tiles[0][i] == PlayerX && b.Tiles[1][i] == PlayerX && b.Tiles[2][i] == PlayerX {
			return PlayerX
		}
		if b.Tiles[0][i] == PlayerO && b.Tiles[1][i] == PlayerO && b.Tiles[2][i] == PlayerO {
			return PlayerO
		}
	}
	if b.Tiles[0][0] == PlayerX && b.Tiles[1][1] == PlayerX && b.Tiles[2][2] == PlayerX {
		return PlayerX
	}
	if b.Tiles[0][0] == PlayerO && b.Tiles[1][1] == PlayerO && b.Tiles[2][2] == PlayerO {
		return PlayerO
	}
	if b.Tiles[2][0] == PlayerX && b.Tiles[1][1] == PlayerX && b.Tiles[0][2] == PlayerX {
		return PlayerX
	}
	if b.Tiles[2][0] == PlayerO && b.Tiles[1][1] == PlayerO && b.Tiles[0][2] == PlayerO {
		return PlayerO
	}
	return Empty
}

func IsOutOfBounds(val uint8, min uint8, max uint8) bool {
	return !(min <= val && max >= val)
}

func (b *Board) TurnTile(p Player, pos Position) error {
	if IsOutOfBounds(pos.x, 0, 2) {
		return ErrPosistionOutOfBounds
	}
	if IsOutOfBounds(pos.y, 0, 2) {
		return ErrPosistionOutOfBounds
	}
	if b.Tiles[pos.x][pos.y] != Empty {
		return ErrTileAlreadyUsed
	}
	b.Tiles[pos.x][pos.y] = p
	return nil
}

func (b *Board) PrintBoard() {
	for i := range b.Tiles {
		fmt.Print("|")
		for j := range b.Tiles[i] {
			fmt.Printf("%c |", b.Tiles[i][j])
		}
		fmt.Print("\n")
	}
}
