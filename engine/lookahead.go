package engine

import (
	"fmt"
	"github.com/wesrobin/battlesnakes/model"
	"math/rand"
)

// Lookahead strategy plays out each of the three (or less!) available moves to the snake. Whichever lasts longest wins.
// If neither strategy wins, chooses one at random after `searchDepth` steps

const searchDepth = 20

func getLookaheadMove(board model.Board) model.Move {
	return bfs(board)
}

// Returns a list of possible moves after iterating `searchDepth` times
func bfs(board model.Board) model.Move {
	possMvs := getPossibleMoves(board)
	moveScores := make(map[model.Move]int)
	for _, mv := range possMvs {
		mv := mv
		d := bfsUtil(board, mv, 0)
		moveScores[mv] = d
		if d < searchDepth {
			fmt.Println(model.PossibleMoves[mv], " dies")
		}
	}

	max := -1
	var finalMove model.Move
	for mv, d := range moveScores {
		if d > max {
			max = d
			finalMove = mv
		}
	}
	return finalMove
}

func bfsUtil(board model.Board, mv model.Move, d int) int {
	d++
	if d == searchDepth {
		return d
	}
	b2 := step(board, mv)
	if b2.Snakes[0].Health == 0 {
		return d
	}
	mvs := getPossibleMoves(board)
	if len(mvs) == 0 {
		return d
	}
	return bfsUtil(b2, mvs[rand.Intn(len(mvs))], d) // FIXME: Don't just move randomly
}

