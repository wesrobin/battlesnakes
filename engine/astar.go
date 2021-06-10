package engine

import (
	"fmt"
	"github.com/wesrobin/battlesnakes/model"
)

type AStar struct {
}

func (as AStar) getMove(board model.Board) model.Move {
	b := board // Don't touch me on my studio
	fmt.Print(b)
	return model.Right
}

