package engine

import "github.com/wesrobin/battlesnakes/model"

var (
	sm map[string]state
)

func init() {
	sm = make(map[string]state)
}

type state struct {
	gobjs      map[model.Coord]model.GameObject
	me         model.Battlesnake
	otherSneks []model.Battlesnake
}

func newState() state {
	return state{
		gobjs: make(map[model.Coord]model.GameObject),
	}
}

func CreateState(id string, board model.Board, me model.Battlesnake) {
	sm[id] = createState(board, me)
}

func createState(board model.Board, me model.Battlesnake) state {
	s := newState()
	s.me = me
	s = updateState(s, board)
	return s
}

func UpdateState(id string, board model.Board, me model.Battlesnake) {
	if _, ok := sm[id]; !ok { // My client sucks, so just do this
		sm[id] = createState(board, me)
	}
	sm[id] = updateState(sm[id], board)
}

func updateState(old state, board model.Board) state {
	s := newState()
	for _, snek := range board.Snakes {
		snek := snek
		if snek.ID == old.me.ID {
			s.me = snek
		} else {
			s.otherSneks = append(s.otherSneks, snek)
		}
	}

	for _, food := range board.Food {
		s.gobjs[food] = model.Food
	}

	for _, snek := range board.Snakes {
		for i, body := range snek.Body {
			body := body
			if i == 0 {
				s.gobjs[body] = model.Head
			} else if i < len(snek.Body)-1 {
				s.gobjs[body] = model.Body
			} else {
				// When we eat an apple or on a new turn the tail can == body, and is always considered last
				if _, ok := s.gobjs[body]; !ok {
					s.gobjs[body] = model.Tail
				}
			}
		}
	}
	return s
}

func DeleteState(id string) {
	delete(sm, id)
}
