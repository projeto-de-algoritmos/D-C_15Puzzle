package game

import (
	"log"
	"math/rand"
)

type Game struct {
	Puzzle []int
}

func NewGame() *Game {
	game := new(Game)
	var newPuzzle []int

	for {
		newPuzzle = rand.Perm(15)
		if isValidGame(newPuzzle) {
			break
		}
		log.Println("Invalid Game")
	}
	game.Puzzle = newPuzzle
	return game
}

func isValidGame(puzzle []int) bool {
	_, c := CountSliceInversions(puzzle)	
	if c%2 == 0 {
		return true
	}	
	return false
}
