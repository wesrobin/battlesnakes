package engine

import (
	"github.com/wesrobin/battlesnakes/model"
	"math/rand"
)

// Lookahead strategy plays out each of the three (or less!) available moves to the snake. Whichever lasts longest wins.
// If neither strategy wins, chooses one at random after `searchDepth` steps

const searchDepth = 10

func getLookaheadMove(board model.Board) model.Move {
	return bfs(board)
}

// Returns a list of possible moves after iterating `searchDepth` times
func bfs(board model.Board) model.Move {
	possMvs := getPossibleMoves(board)
	moveScores := make(map[model.Move]int)
	for _, mv := range possMvs {
		mv := mv
		d := bfsUtil(board, mv, searchDepth)
		moveScores[mv] = d
	}

	min := searchDepth + 1
	var finalMove model.Move
	for mv, d := range moveScores {
		if d < min {
			min = d
			finalMove = mv
		}
	}
	return finalMove
}

func bfsUtil(board model.Board, mv model.Move, d int) int {
	d--
	if d == 0 {
		return 0
	}
	//printMap(board)
	b2 := step(board, mv)
	//printMap(b2)
	mvs := getPossibleMoves(board)
	if len(mvs) == 0 {
		return d
	}
	return bfsUtil(b2, mvs[rand.Intn(len(mvs))], d) // FIXME: Don't just move randomly
}

