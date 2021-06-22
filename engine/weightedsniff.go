package engine

import (
	"github.com/wesrobin/battlesnakes/model"
	"math"
	"math/rand"
)

type WeightedSniff struct {
	s state
}

var (
	myTail  = 50
	illegal = 0

	// Misc
	sniffRadius = 10
)

type moveWeight struct {
	mv     model.Move
	weight int
}

func (ws WeightedSniff) GetMove(s state, board model.Board) model.Move {
	ws.s = s
	sniffRadius = (board.Width + board.Height)/2
	cs := sniffedCoords(ws.s.me)
	weights := make(map[model.Coord]int)
	for _, c := range cs {
		w := ws.weightMyCoord(board, c)
		weights[c] = w
	}
	u, d, l, r := moveWeights(ws.s.me.Head, weights)

	// Don't do illegal stuff
	head := ws.s.me.Head
	if !legalCoord(ws.s, board, head.Move(model.Up)) {
		u.weight = 0
	}
	if !legalCoord(ws.s, board, head.Move(model.Down)) {
		d.weight = 0
	}
	if !legalCoord(ws.s, board, head.Move(model.Left)) {
		l.weight = 0
	}
	if !legalCoord(ws.s, board, head.Move(model.Right)) {
		r.weight = 0
	}

	mvs := []*moveWeight{&u, &d, &l, &r}
	for i, mv := range mvs {
		if mv.weight == 0 {
			continue
		}
		b := step(board, mv.mv)
		free := ws.moveableSquares(b)
		mvs[i].weight += free/2
	}

	return chooseMove(u.weight, d.weight, l.weight, r.weight)
}

func (ws WeightedSniff) moveableSquares(b model.Board) int {
	seen := make(map[model.Coord]bool)
	queue := make([]model.Coord, 0)
	total := 0

	head := b.Snakes[0].Head
	adj := getAdjacent(head)
	for _, coord := range adj {
		if inBounds(b, coord) {
			queue = append(queue, coord)
		}
	}

	for {
		if len(queue) == 0 {
			break
		}
		c := queue[0]
		queue = queue[1:]
		if ws.s.gobjs[c] == model.Body || !inBounds(b, c) || seen[c] {
			continue
		}
		seen[c] = true
		total++
		adjs := getAdjacent(c)
		for _, adj := range adjs {
			if !seen[adj] {
				queue = append(queue, adj)
			}
		}
	}

	return total
}

func getAdjacent(cell model.Coord) []model.Coord {
	return []model.Coord{
		{X: cell.X + 1, Y: cell.Y},
		{X: cell.X - 1, Y: cell.Y},
		{X: cell.X, Y: cell.Y + 1},
		{X: cell.X, Y: cell.Y - 1},
	}
}

func moveWeights(head model.Coord, weights map[model.Coord]int) (u, d, l, r moveWeight) {
	u = moveWeight{mv: model.Up}
	d = moveWeight{mv: model.Down}
	l = moveWeight{mv: model.Left}
	r = moveWeight{mv: model.Right}
	for c, w := range weights {
		//dist := float64(dist(head, c))
		//dist = math.Pow(dist, 2)
		dist := float64(distTaxi(head, c))
		if c.X > head.X {
			r.weight += int(float64(w) / dist)
		}
		if c.X < head.X {
			l.weight += int(float64(w) / dist)
		}
		if c.Y > head.Y {
			u.weight += int(float64(w) / dist)
		}
		if c.Y < head.Y {
			d.weight += int(float64(w) / dist)
		}
	}
	return
}

func withinRange(val, max int) bool {
	return float32(max-val)/float32(max) < 0.1
}

func chooseMove(u, d, l, r int) model.Move {
	max := int(math.Max(float64(u), math.Max(float64(d), math.Max(float64(l), float64(r)))))
	choices := make([]model.Move, 0)
	for i := 0 ; i < u && withinRange(u, max) ; i++ {
		choices = append(choices, model.Up)
	}
	for i := 0 ; i < d && withinRange(d, max) ; i++ {
		choices = append(choices, model.Down)
	}
	for i := 0 ; i < l && withinRange(l, max) ; i++ {
		choices = append(choices, model.Left)
	}
	for i := 0 ; i < r && withinRange(r, max) ; i++ {
		choices = append(choices, model.Right)
	}
	if len( choices) == 0 {
		return model.Up
	}
	return choices[rand.Intn(len(choices))]
}

func sniffedCoords(me model.Battlesnake) []model.Coord {
	head := me.Head
	neck := me.Body[1]
	var sniffs []model.Coord
	for x := head.X - sniffRadius; x <= head.X+sniffRadius; x++ {
		for y := head.Y - sniffRadius; y <= head.Y+sniffRadius; y++ {
			coord := model.Coord{X: x, Y: y}
			if coord == head || coord == neck {
				continue
			}
			sniffs = append(sniffs, coord)
		}
	}
	return sniffs
}

func (ws WeightedSniff) weightMyCoord(board model.Board, coord model.Coord) int {
	if !inBounds(board, coord) {
		return illegal
	}

	// Check if adjacent to other snek head - these are bad because we don't know what they will do
	for _, snek := range ws.s.otherSneks {
		for _, c := range getAdjacent(snek.Head) {
			if c == coord {
				return illegal
			}
		}
	}

	// Check is food
	if ws.s.gobjs[coord] == model.Food {
		return foodWeight(ws.s.me, board)
	} else if ws.s.gobjs[coord] == model.Body {
		return illegal
	} else if ws.isMyTail(board, coord) {
		return tailWeight(ws.s.me, board)
	}

	return 10
}

func foodWeight(me model.Battlesnake, board model.Board) int {
	if me.Health > 50 {
		return 2
	} else if me.Health > 30 {
		return 10
	}
	return 30
}

func tailWeight(me model.Battlesnake, board model.Board) int {
	if me.Health > 50 {
		return myTail
	} else if me.Health > 30 {
		return myTail/2
	}
	return 10
}

func (ws WeightedSniff) isMyTail(board model.Board, coord model.Coord) bool {
	if ws.s.gobjs[coord] != model.Tail {
		return false
	}
	// At the start the tail overlaps some segments, just check that it'sm not doing that too pls
	return coord == ws.s.me.Body[ws.s.me.Length-1] && coord != ws.s.me.Body[ws.s.me.Length-2]
}
