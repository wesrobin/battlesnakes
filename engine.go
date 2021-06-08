package main

import (
	"fmt"
	"github.com/wesrobin/battlesnakes/model"
	"math/rand"
)

func getMove(state model.Board) string {
	// Choose a random direction to move in
	for _, move := range model.PossibleMoves {
		coord := getCoordAfterMove(state.Snakes[0].Head, move)
		if legalCoord(state, coord) {
			fmt.Printf("Chose move: %s to coord: (%d,%d)\n", move, coord.X, coord.Y)
			return move.String()
		}
	}
	move := model.PossibleMoves[rand.Intn(len(model.PossibleMoves))].String()
	fmt.Printf("Chose move at random: %s\n", move)
	// Just return random move lol
	return move
}

func getCoordAfterMove(head model.Coord, move model.Move) model.Coord {
	switch move {
	case model.Up:
		return model.Coord{X: head.X, Y: head.Y+1}
	case model.Down:
		return model.Coord{X: head.X, Y: head.Y-1}
	case model.Left:
		return model.Coord{X: head.X-1, Y: head.Y}
	case model.Right:
		return model.Coord{X: head.X+1, Y: head.Y}
	}
	return model.Coord{} // Hiss
}

func legalCoord(state model.Board, coord model.Coord) bool {
	if coord.X < 0 {
		return false
	}
	if coord.Y < 0 {
		return false
	}
	if coord.Y >= state.Height {
		return false
	}
	if coord.X >= state.Width {
		return false
	}

	for _, part := range state.Snakes[0].Body {
		if coord == part {
			return false
		}
	}

	return true // If it makes it here we are ok.
}