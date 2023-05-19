package gamelogic

import (
	"errors"
)

type Board struct {
	Num_rows, Num_cols int      // how many rows and columns this board has
	CurrPlayer         int      // the current player
	Players            [2]byte  // characters representing the players
	NextPlacement      []int    // where each move should go on a given row
	Table              [][]byte // represents the actual board
}

// create a board with the specified attributes
func InitBoard(n_rows, n_cols int, player1, player2 byte) (Board, error) {

	var b Board

	if n_rows <= 0 || n_cols <= 0 {
		return b, errors.New("impossible board attributes")
	}

	if player1 == player2 {
		return b, errors.New("players must be distinct")
	}

	players := [2]byte{player1, player2}

	n_placement := make([]int, n_cols)
	for i := range n_placement {
		n_placement[i] = n_rows - 1
	}

	table := make([][]byte, n_rows)
	for i := range table {
		table[i] = make([]byte, n_cols)
	}

	b = Board{n_rows, n_cols, 0, players, n_placement, table}
	return b, nil
}

// return the board representation as a string
func (b *Board) String() string {
	var str string

	if b.Num_rows == 0 {
		return str
	}

	for i := 0; i < b.Num_rows; i++ {
		if len(b.Table[i]) == 0 {
			return str
		} else if len(b.Table[i]) != b.Num_cols {
			return str
		}
	}

	str += "  "
	for i := 0; i < b.Num_cols*2+1; i++ {
		str += "-"
	}
	str += " \n"
	for i := 0; i < b.Num_rows; i++ {
		for j := 0; j < b.Num_cols; j++ {
			if j == 0 {
				str += " |"
			}
			str += " "
			if b.Table[i][j] == 0 {
				str += " "
			} else {
				str += string(b.Table[i][j])
			}

			if j == b.Num_cols-1 {
				str += " |"
			}
		}
		str += "\n"
	}
	str += "  "
	for i := 0; i < b.Num_cols*2+1; i++ {
		str += "-"
	}
	str += " \n"
	return str
}

// place the players move on the board. returns an error if move is not allowed
func (b *Board) Move(y int) (int, int, error) {

	if y < 0 || y > b.Num_cols-1 {
		return -1, -1, errors.New("out of bounds move")
	}

	if b.NextPlacement[y] < 0 {
		return -1, -1, errors.New("column full")
	}

	b.Table[b.NextPlacement[y]][y] = b.Players[b.CurrPlayer]
	b.NextPlacement[y]--
	return b.NextPlacement[y] + 1, y, nil
}

// returns the outcome of the game. 0 for tie, 1 for win and 2 for keep playing
func (b *Board) Outcome(x, y int) int {
	var connected int = 1

	// check right horizontal
	for j := y + 1; j < b.Num_cols; j++ {
		if b.Table[x][j] == b.Players[b.CurrPlayer] {
			connected++
			if connected == 4 {
				return 1
			}
		} else {
			break
		}
	}

	// check left horizontal
	for j := y - 1; j >= 0; j-- {
		if b.Table[x][j] == b.Players[b.CurrPlayer] {
			connected++
			if connected == 4 {
				return 1
			}
		} else {
			break
		}
	}

	// check vertical
	connected = 1
	for i := x + 1; i < b.Num_rows; i++ {
		if b.Table[i][y] == b.Players[b.CurrPlayer] {
			connected++
			if connected == 4 {
				return 1
			}
		} else {
			break
		}
	}

	// check increasing diagonal
	connected = 1
	for i, j := x+1, y+1; i < b.Num_rows && j < b.Num_cols; i, j = i+1, j+1 {
		if b.Table[i][j] == b.Players[b.CurrPlayer] {
			connected++
			if connected == 4 {
				return 1
			}
		} else {
			break
		}
	}

	for i, j := x-1, y-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if b.Table[i][j] == b.Players[b.CurrPlayer] {
			connected++
			if connected == 4 {
				return 1
			}
		} else {
			break
		}
	}

	// check decreasing diagonal
	connected = 1
	for i, j := x+1, y+1; i < b.Num_rows && j < b.Num_cols; i, j = i+1, j+1 {
		if b.Table[i][j] == b.Players[b.CurrPlayer] {
			connected++
			if connected == 4 {
				return 1
			}
		} else {
			break
		}
	}

	for i, j := x-1, y-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if b.Table[i][j] == b.Players[b.CurrPlayer] {
			connected++
			if connected == 4 {
				return 1
			}
		} else {
			break
		}
	}
	//TODO: check for tie
	for _, val := range b.NextPlacement {
		if val >= 0 {
			return 2
		}
	}
	return 0
}

// move to the next player
func (b *Board) NextPlayer() {
	b.CurrPlayer = (b.CurrPlayer + 1) % 2
}
