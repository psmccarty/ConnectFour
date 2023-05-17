package gamelogic

import (
	"testing"
)

const NUM_ROWS int = 6 // number of rows
const NUM_COLS int = 7 // number of columns

func TestEmptyPrintBoard(t *testing.T) {

	var testBoard [][]byte

	// empty board
	val, err := PrintBoard(testBoard)
	if val != "" || err == nil {
		t.Fatalf("Fail on empty board")
	}

	// board with no columns
	testBoard = make([][]byte, NUM_ROWS)
	val, err = PrintBoard(testBoard)
	if val != "" || err == nil {
		t.Fatalf("Fail on all empty columns")
	}

	// not all columns filled
	for i := 0; i < NUM_ROWS-1; i++ {
		testBoard[i] = make([]byte, NUM_COLS)
		val, err = PrintBoard(testBoard)
		if val != "" || err == nil {
			t.Fatalf("Fail on some empty columns")
		}
	}

	// columns of different sizes
	testBoard[NUM_ROWS-1] = make([]byte, NUM_COLS-1)
	val, err = PrintBoard(testBoard)
	if val != "" || err == nil {
		t.Fatalf("Fail on column of wrong size")
	}
}

func TestValidPrintBoard(t *testing.T) {
	testBoard := make([][]byte, NUM_ROWS)
	for i := range testBoard {
		testBoard[i] = make([]byte, NUM_COLS)
	}

	val, err := PrintBoard(testBoard)
	if val == "" || err != nil || len(val) != (NUM_ROWS+2)*(2*(NUM_COLS)+5) {
		t.Fatalf("Fail on valid blank board")
	}

}
