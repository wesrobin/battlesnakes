package engine

import (
	"fmt"
	"github.com/wesrobin/battlesnakes/model"
	"testing"
)

func TestMoveableSquares(t *testing.T) {
	snek := model.Battlesnake{
		Body: []model.Coord{{9, 10}, {9, 9}, {10, 9}, {10, 8}, {9, 8}, {9, 7}},
		Head: model.Coord{X: 9, Y: 10},
	}
	board := model.Board{
		Snakes: []model.Battlesnake{snek},
		Width: 11,
		Height: 11,
	}
	//printMap(board)
	//fmt.Println(moveableSquares(board))
	//
	fmt.Println("\n***RIGHT***")
	rB := step(board, model.Right)
	UpdateState(rB)
	printMap(rB)
	fmt.Println(moveableSquares(rB))
	//
	//fmt.Println("\n***LEFt***")
	//lB := step(board, model.Left)
	//printMap(lB)
	//fmt.Println(moveableSquares(lB))
}
