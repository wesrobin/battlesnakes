package engine

import (
	"fmt"
	"github.com/wesrobin/battlesnakes/model"
	"math"
	"sync"
)

// Lookahead strategy plays out each of the three (or less!) available moves to the snake. Whichever lasts longest wins.
// If neither strategy wins, chooses one at random after `searchDepth` steps

type Lookahead struct {
	cache sync.Map
	s state
}

const searchDepth = 10

func (la *Lookahead) GetMove(board model.Board) model.Move {
	b := board // Don't touch the original
	return la.dfs(b)
}


// Returns a list of possible moves after iterating `searchDepth` times
func (la *Lookahead) dfs(board model.Board) model.Move {
	possMvs := getPossibleMoves(la.s, board)
	moveScores := make(map[model.Move]int)
	for _, mv := range possMvs {
		board := board
		mv := mv
		d := la.dfsUtil(&board, mv, 0)
		moveScores[mv] = d
		if d == searchDepth {
			break
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

func makeKey(board model.Board, move model.Move) string {
	return fmt.Sprintf("%v:%v", board.Hash(), move)
}

func (la *Lookahead) dfsUtil(board *model.Board, mv model.Move, d int) int {
	k := makeKey(*board, mv)
	if val, ok := la.cache.Load(k); ok {
		return val.(int)
	}
	d++
	if d == searchDepth {
		return d
	}
	b := step(*board, mv)
	if b.Snakes[0].Health == 0 {
		return d
	}
	mvs := getPossibleMoves(la.s, b)
	if len(mvs) == 0 {
		return d
	}
	max := float64(-1)
	for _, mv := range mvs {
		mv := mv
		d := la.dfsUtil(&b, mv, d)
		max = math.Max(max, float64(d))
	}
	la.cache.Store(k, d)
	return int(max)
}

