package engine

import "github.com/wesrobin/battlesnakes/model"

var (
	state map[model.Coord]model.GameObject
)

func genState(board model.Board) map[model.Coord]model.GameObject {
	state := make(map[model.Coord]model.GameObject)
	for _, food := range board.Food {
		state[food] = model.Food
	}

	for _, snek := range board.Snakes {
		for i, body := range snek.Body {
			body := body
			if i < len(snek.Body)-1 {
				state[body] = model.Snake
			} else {
				// When we eat an apple or on a new turn the tail can == body, and is always considered last
				if _, ok := state[body] ; !ok {
					state[body] = model.Tail
				}
			}
		}
	}
	return state
}

func UpdateState(board model.Board) {
	state = genState(board)
}
