package engine

import (
	"github.com/wesrobin/battlesnakes/model"
)

func GetMove(state model.Board) string {
	s := state
	printMap(s)
	move := getLookaheadMove(s)

	// --- Base case ---
	//// Choose a random direction to move in
	//for move, moveS := range model.PossibleMoves {
	//	coord := getCoordAfterMove(state.Snakes[0].Head, move)
	//	if legalCoord(state, coord) {
	//		fmt.Printf("Chose move: %s to coord: (%d,%d)\n", moveS, coord.X, coord.Y)
	//		return moveS
	//	}
	//}
	//move := model.PossibleMoves[model.Move(rand.Intn(len(model.PossibleMoves)))]
	//fmt.Printf("Chose move: %s\n", model.PossibleMoves[move])
	//// Just return random move lol
	return model.PossibleMoves[move]
}