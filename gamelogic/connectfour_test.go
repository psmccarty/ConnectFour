package gamelogic

import (
	"testing"
)

const DEFAULT_ROWS int = 6  // number of rows
const DEFAULT_COLS int = 7  // number of columns
const DEFAULT_P1 byte = 'X' // number of columns
const DEFAULT_P2 byte = 'O' // number of columns

// test the functionality of InitBoard
func TestInitBoard(t *testing.T) {
	var tests = []struct {
		n_rows    int
		n_cols    int
		player1   byte
		player2   byte
		expectErr bool
	}{
		{0, DEFAULT_COLS, DEFAULT_P1, DEFAULT_P2, true},
		{DEFAULT_ROWS, 0, DEFAULT_P1, DEFAULT_P2, true},
		{DEFAULT_ROWS, DEFAULT_COLS, DEFAULT_P1, DEFAULT_P1, true},
		{DEFAULT_ROWS, DEFAULT_COLS, DEFAULT_P1, DEFAULT_P2, false},
		{DEFAULT_ROWS + 3, DEFAULT_COLS + 1, DEFAULT_P1, DEFAULT_P2, false},
	}
	for _, test := range tests {
		got, err := InitBoard(test.n_rows, test.n_cols, test.player1, test.player2)

		if test.expectErr && err == nil {
			t.Errorf("InitBoard(%v, %v, %v, %v) expected error but got none\n",
				test.n_rows, test.n_cols, test.player1, test.player2)

		} else if !test.expectErr && err != nil {
			t.Errorf("InitBoard(%v, %v, %v, %v) expected unexpected error %q\n",
				test.n_rows, test.n_cols, test.player1, test.player2, err)

		} else {
			if test.expectErr == false && (got.Num_rows != test.n_rows || 
			got.Num_cols != test.n_cols || got.Players[0] != test.player1 ||
			got.Players[1] != test.player2 || len(got.Table) != test.n_rows ||
			len(got.Table[0]) != test.n_cols || got.CurrPlayer != 0 || 
			len(got.NextPlacement) != test.n_cols) {

			t.Errorf("InitBoard(%v, %v, %v, %v) got n_rows=(%v,%v), n_cols=(%v,%v), "+
				"players=(%c,%c), currPlayer=(%d), len(nextPlacement)=(%d)\n",
				test.n_rows, test.n_cols, test.player1, test.player2,
				got.Num_rows, len(got.Table), got.Num_cols, len(got.Table[0]),
				got.Players[0], got.Players[1], got.CurrPlayer, len(got.NextPlacement))
			}

			for i, val := range got.NextPlacement {
					if val != test.n_rows-1 {
						t.Errorf("got NextPlacement[%d] = %d for InitBoard(%d,%d,%s,%s)"+
							" when expected %d\n",
							i, val, test.n_rows, test.n_cols, string(test.player1),
							string(test.player2), test.n_rows-1)
					}
			}
		}
	}
}
