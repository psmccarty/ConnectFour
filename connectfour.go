package main

import (
	"fmt"
	"strconv"
)


const NUM_ROWS int = 6			// number of rows
const NUM_COLS int = 7			// number of columns
var board [NUM_ROWS][NUM_COLS]byte 	// board for this game
var currPlayer byte = 'X'		// current player
var nextPlacement [NUM_COLS]int 

func main(){
	for i := 0; i < NUM_COLS; i++ {
		nextPlacement[i] = NUM_ROWS - 1
	}
	
	for {
		printBoard()
		x, y := playerMove()
		if x == -1 || y == -1 {
			fmt.Println("That is not a valid move")
			continue
		}
		if hasWon(x, y) {
			break
		}
		if currPlayer == 'X' {
			currPlayer = 'O'
		} else {
			currPlayer = 'X'
		}
	}
	printBoard()
	fmt.Printf("Player %q has won!\n", currPlayer)
}

// print the board to the screen
func printBoard(){
	fmt.Print("  ")
	for i := 0; i < NUM_COLS * 2 + 1; i++ {
		fmt.Print("-")
	}
	fmt.Print("\n")
	for i := 0; i < NUM_ROWS; i++ {
		for j := 0; j < NUM_COLS; j++ {
			if j == 0 {
				fmt.Print(" | ") 
			}
			fmt.Printf("%c ", board[i][j])
			if j == NUM_COLS - 1 {
				fmt.Print("|")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("  ")
	for i := 0; i < NUM_COLS * 2 + 1; i++ {
		fmt.Print("-")
	}
	fmt.Print("\n")
}

func playerMove() (int, int){
	var str string
	var move int
	fmt.Print("Input your move: ")
	fmt.Scanln(&str)
	move, err := strconv.Atoi(str)
	if err != nil || move < 0 || move > NUM_COLS - 1 || nextPlacement[move] < 0 {
		return -1, -1
	}

	
	board[nextPlacement[move]][move] = currPlayer
	nextPlacement[move]--

	return nextPlacement[move] + 1, move
}

func hasWon(x, y int) bool{
	var connected int = 1

	// check right horizontal 
	for j := y + 1; j < NUM_COLS; j++ {
		if board[x][j] == currPlayer {
			connected++
			if connected == 4{
				return true
			}
		} else{
			break;
		}
	}	

	// check left horizontal 
	for j := y - 1; j >= 0; j-- {
		if board[x][j] == currPlayer {
			connected++
			if connected == 4{
				return true
			}
		} else{
			break
		}
	}

	connected = 1
	// check vertical
	for i := x + 1; i < NUM_ROWS; i++ {
		if board[i][y] == currPlayer{
			connected++
			if connected == 4 {
				return true
			}
		} else {
			break
		}
	}


	// check increasing diagonal
	connected = 1
	for i, j := x + 1, y + 1; i < NUM_ROWS && j < NUM_COLS; i, j = i+1, j+1 {
	//for i, j := x + 1, y + 1; i < NUM_ROWS && j < NUM_COLS; i++, j++ {
		if board[i][j] == currPlayer{
			connected++
			if connected == 4 {
				return true
			}
		} else {
			break
		}
	}

	for i, j := x - 1, y - 1; i >= 0 && j >= 0; i, j = i-1, j-1{
		if board[i][j] == currPlayer{
			connected++
			if connected == 4 {
				return true
			}
		} else {
			break
		}
	}

	// check decreasing diagonal
	connected = 1
	for i, j := x + 1, y + 1; i < NUM_ROWS && j < NUM_COLS; i, j = i+1, j+1{
		if board[i][j] == currPlayer{
			connected++
			if connected == 4 {
				return true
			}
		} else {
			break
		}
	}

	for i, j := x - 1, y - 1; i >= 0 && j >= 0; i, j = i-1, j-1{
		if board[i][j] == currPlayer{
			connected++
			if connected == 4 {
				return true
			}
		} else {
			break
		}
	}
	return false
}
