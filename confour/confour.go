package main

import (
	"fmt"
	"strconv"

	"psmccarty.com/gamelogic"
)

 const NUM_ROWS int = 6
 const NUM_COLS int = 7
const PLAYER1 byte = 'X'
const PLAYER2 byte = 'O'

func main() {

	gameBoard, err := gamelogic.InitBoard(NUM_ROWS, NUM_COLS, PLAYER1, PLAYER2)

	if err != nil {
		panic("cannot initialize board")
	}

	for {
		fmt.Println(gameBoard.String())
		move, err := getMove()
		if err != nil {
			fmt.Println("Invalid move. Try again")
			continue
		}

		x, y, err := gameBoard.Move(move)
		if err != nil {
			fmt.Println(err)
			continue
		}

		result := gameBoard.Outcome(x, y)

		if result == 0 {
			fmt.Println(gameBoard.String())
			fmt.Println("Its a tie!")
			break
		} else if result == 1 {
			fmt.Println(gameBoard.String())
			fmt.Printf("Player %q wins!\n", gameBoard.Players[gameBoard.CurrPlayer])
			break
		}
		gameBoard.NextPlayer()
	}
}

// get the players move
func getMove() (int, error) {
	var str string
	fmt.Print("Input your move: ")
	fmt.Scanln(&str)

	return strconv.Atoi(str)
}
