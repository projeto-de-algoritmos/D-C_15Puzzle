package game

import (
	"log"
	"math/rand"
	"strconv"
)

type Game struct {
	Puzzle [][]string
	size int
	selected point
	empty point
}

func NewGame(size int) *Game {
	game := new(Game)
	game.size = size
	game.empty.x = size-1
	game.empty.y = size-1

	var newPuzzle []int

	for {
		newPuzzle = rand.Perm(size*size-1)
		if isValidGame(newPuzzle) {
			break
		}
		log.Println("Invalid Game")
	}

	game.Puzzle = make([][]string, size)
	for i := range game.Puzzle {
		game.Puzzle[i] = make([]string, size)
	}
	for i := 0; i<size*size-1; i++ {
		game.Puzzle[i/size][i%size] =  strconv.Itoa(newPuzzle[i]+1)
	}
	game.Puzzle[size-1][size-1] = "X"

	return game
}

func isValidGame(puzzle []int) bool {
	_, c := CountSliceInversions(puzzle)	
	return c%2 == 0
}
