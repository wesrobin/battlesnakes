package engine

import "github.com/wesrobin/battlesnakes/model"

func getCoordAfterMove(coord model.Coord, move model.Move) model.Coord {
	switch move {
	case model.Up:
		return model.Coord{X: coord.X, Y: coord.Y + 1}
	case model.Down:
		return model.Coord{X: coord.X, Y: coord.Y - 1}
	case model.Left:
		return model.Coord{X: coord.X - 1, Y: coord.Y}
	case model.Right:
		return model.Coord{X: coord.X + 1, Y: coord.Y}
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

// Pls only call with legal moves <3
func step(board model.Board, mv model.Move) model.Board {
	snek := board.Snakes[0]
	newHead := getCoordAfterMove(snek.Head, mv)

	l := len(snek.Body)
	var hazCheez bool
	for _, cheezes := range board.Food {
		if newHead == cheezes {
			hazCheez = true
			l++
			break
		}
	}

	newBod := make([]model.Coord, l)
	newBod[0] = newHead
	for i := 1; i < len(snek.Body); i++ {
		newBod[i] = snek.Body[i-1]
	}
	if hazCheez {
		newBod[len(newBod)-1] = snek.Body[len(snek.Body)-1]
	}
	snek.Body = newBod
	return board
}

func getPossibleMoves(board model.Board) []model.Move {
	var pMvs []model.Move
	for mv := range model.PossibleMoves {
		coord := getCoordAfterMove(board.Snakes[0].Head, mv)
		if legalCoord(board, coord) {
			pMvs = append(pMvs, mv)
		}
	}
	return pMvs
}
