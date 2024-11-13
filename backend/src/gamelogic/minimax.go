package gamelogic

import "math"

func Move(gs *Tictactoe) {
	res, _ := gs.GameOver()
	if res {
		return
	}

	var highScore int64 = math.MinInt64
	var move Position

	for i := range gs.Board.Tiles {
		for j := range gs.Board.Tiles[i] {
			if gs.Board.Tiles[i][j] == Empty {
				gs.Board.Tiles[i][j] = PlayerO
				score := minimax(&gs.Board, PlayerX)
				gs.Board.Tiles[i][j] = Empty
				if score > highScore {
					highScore = score
					move = Position{X: uint8(i), Y: uint8(j)}
				}
			}
		}
	}
	gs.Board.Tiles[move.X][move.Y] = PlayerO
	gs.Turn = PlayerX
}

func minimax(b *Board, p Player) int64 {
	if winner := b.CheckWinner(); winner != Empty {
		switch winner {
		case PlayerO:
			return 1
		case PlayerX:
			return -1
		}
	}
	if b.BoardIsFull() {
		return 0
	}

	if p == PlayerO {
		var highScore int64 = math.MinInt64
		for i := range b.Tiles {
			for j := range b.Tiles[i] {
				if b.Tiles[i][j] == Empty {
					b.Tiles[i][j] = PlayerO
					score := minimax(b, PlayerX)
					b.Tiles[i][j] = Empty
					highScore = max(score, highScore)
				}
			}
		}
		return highScore
	} else {
		var lowScore int64 = math.MaxInt64
		for i := range b.Tiles {
			for j := range b.Tiles[i] {
				if b.Tiles[i][j] == Empty {
					b.Tiles[i][j] = PlayerX
					score := minimax(b, PlayerO)
					b.Tiles[i][j] = Empty
					lowScore = min(score, lowScore)
				}
			}
		}
		return lowScore
	}
}
