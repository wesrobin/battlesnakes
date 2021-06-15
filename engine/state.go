package engine

import "github.com/wesrobin/battlesnakes/model"

var (
	state map[model.Coord]model.GameObject
)

func updateState(board model.Board) {
	state = make(map[model.Coord]model.GameObject)
	for _, food := range board.Food {
		state[food] = model.Food
	}

	for _, snek := range board.Snakes {
		for _, body := range snek.Body {
			body := body
			state[body] = model.Snake
		}
	}
}
