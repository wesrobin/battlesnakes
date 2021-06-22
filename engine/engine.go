package engine

import (
	"fmt"
	"github.com/wesrobin/battlesnakes/model"
)

type Strategy interface {
	GetMove(s state, board model.Board) model.Move
}

func GetMove(id string, board model.Board) string {
	s := board

	printMap(s)
	move := WeightedSniff{}.GetMove(sm[id], s)
	fmt.Println("chose", model.PossibleMoves[move])
	fmt.Println("*******END*******")

	return model.PossibleMoves[move]
}