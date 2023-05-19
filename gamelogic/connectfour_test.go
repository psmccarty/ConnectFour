package gamelogic

import (
	"testing"
)

const NUM_ROWS int = 6 // number of rows
const NUM_COLS int = 7 // number of columns

const DEFAULT_ROWS int = 6  // number of rows
const DEFAULT_COLS int = 7  // number of columns
const DEFAULT_P1 byte = 'X' // number of columns
const DEFAULT_P2 byte = 'O' // number of columns

// test the functionality of InitBoard
func TestInitBoard(t *testing.T) {
	type initInput struct {
		n_rows    int
		n_cols    int
		player1   byte
		player2   byte
		expectErr bool
	}

	var inputs = []initInput{
		{0, DEFAULT_COLS, DEFAULT_P1, DEFAULT_P2, true},
		{DEFAULT_ROWS, 0, DEFAULT_P1, DEFAULT_P2, true},
		{DEFAULT_ROWS, DEFAULT_COLS, DEFAULT_P1, DEFAULT_P1, true},
		{DEFAULT_ROWS, DEFAULT_COLS, DEFAULT_P1, DEFAULT_P2, false},
		{DEFAULT_ROWS + 3, DEFAULT_COLS + 1, DEFAULT_P1, DEFAULT_P2, false},
	}

	for _, test := range inputs {
		gameBoard, err := InitBoard(test.n_rows, test.n_cols, test.player1, test.player2)

		if test.expectErr && err == nil {
			t.Errorf("expected error on input: (%d, %d, %s, %s), but got none\n",
				test.n_rows, test.n_cols, string(test.player1), string(test.player2))

		} else if !test.expectErr {

			if err != nil {
				t.Errorf("got error on input: (%d, %d, %s, %s), when expected none\n",
					test.n_rows, test.n_cols, string(test.player1), string(test.player2))
			}

			if gameBoard.Num_rows != test.n_rows || len(gameBoard.Table) != test.n_rows {
				t.Errorf("got num_rows = %d for InitBoard(%d, %d, %s, %s), when expected %d\n",
					gameBoard.Num_rows, test.n_rows, test.n_cols, string(test.player1),
					string(test.player2), test.n_rows)
			}

			if gameBoard.Num_cols != test.n_cols || len(gameBoard.Table[0]) != test.n_cols {
				t.Errorf("got num_cols = %d for InitBoard(%d, %d, %s, %s), when expected %d\n",
					gameBoard.Num_cols, test.n_rows, test.n_cols, string(test.player1),
					string(test.player2), test.n_cols)
			}

			if gameBoard.Players[0] != test.player1 || gameBoard.Players[1] != test.player2 ||
				gameBoard.CurrPlayer != 0 {
				t.Errorf("got (p1,p2) = (%s,%s) for InitBoard(%d,%d,%s,%s) when expected (%s,%s)\n",
					string(gameBoard.Players[0]), string(gameBoard.Players[1]), test.n_rows,
					test.n_cols, string(test.player1), string(test.player2), string(test.player1),
					string(test.player2))
			}

			if len(gameBoard.NextPlacement) != test.n_cols{
				t.Errorf("got len(NextPlacement) = %d for InitBoard(%d,%d,%s,%s)" +
				" when expected %d\n",
					len(gameBoard.NextPlacement), test.n_rows, test.n_cols,
					string(test.player1), string(test.player2), test.n_cols)

			} else {
				for _, val := range gameBoard.NextPlacement {
					if val != test.n_rows - 1 {
						t.Errorf("got NextPlacement[i] = %d for InitBoard(%d,%d,%s,%s)" +
							" when expected %d\n",
						val, test.n_rows, test.n_cols, string(test.player1), string(test.player2),
						test.n_rows - 1)
					}
				}
			}
		}
	}
}

