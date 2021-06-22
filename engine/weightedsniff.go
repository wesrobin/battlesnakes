package engine

import (
	"fmt"
	"github.com/wesrobin/battlesnakes/model"
	"math"
	"math/rand"
)

// Author's note: This is some of the most bad code I have written, sorry future self. Code standards would be good.

type WeightedSniff struct {
	s state
}

var (
	myTail  = 15
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
	sniffRadius = (board.Width + board.Height) / 2
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
		mvs[i].weight += free / 2

	}

	for i, mv := range mvs {
		// Check if adjacent to other snek head - these are bad because we don't know what they will do
		for _, snek := range s.otherSneks {
			if len(s.me.Body) > len(snek.Body) {
				continue
			}
			if isAdjacent(ws.s.me.Head.Move(mv.mv), snek.Head) {
				mvs[i].weight = 0
			}
		}
	}

	return chooseMove(u.weight, d.weight, l.weight, r.weight)
}

func (ws WeightedSniff) moveableSquares(b model.Board) int {
	seen := make(map[model.Coord]bool)
	queue := make([]model.Coord, 0)
	total := 0

	head := b.Snakes[0].Head
	adj := adjacentCells(head)
	for _, coord := range adj {
		if legalCoord(ws.s, b, coord) {
			queue = append(queue, coord)
		}
	}

	for {
		if len(queue) == 0 {
			break
		}
		c := queue[0]
		queue = queue[1:]
		if legalCoord(ws.s, b, c) || seen[c] {
			continue
		}
		seen[c] = true
		total++
		adjs := adjacentCells(c)
		for _, adj := range adjs {
			if !seen[adj] {
				queue = append(queue, adj)
			}
		}
	}

	return total
}

func adjacentCells(cell model.Coord) []model.Coord {
	return getLines(cell, 1)
}

func getLines(cell model.Coord, distance int) []model.Coord {
	var cs []model.Coord
	for i := 1; i <= distance; i++ {
		cs = append(cs,
			model.Coord{X: cell.X + i, Y: cell.Y},
			model.Coord{X: cell.X - i, Y: cell.Y},
			model.Coord{X: cell.X, Y: cell.Y + i},
			model.Coord{X: cell.X, Y: cell.Y - i},
		)
	}
	return cs
}

func isAdjacent(a, b model.Coord) bool {
	return isAdjacentWithMoves(a, b, 1)
}

func isAdjacentWithMoves(a, b model.Coord, moves int) bool {
	return distTaxi(a, b) <= int64(moves)
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

// Ensure that the given val is within 10% of the given max
func withinRange(val, max int) bool {
	return float32(max-val)/float32(max) < 0.1
}

func chooseMove(u, d, l, r int) model.Move {
	fmt.Printf("u %d\nd %d\nl %d\nr %d", u, d, l, r)
	max := int(math.Max(float64(u), math.Max(float64(d), math.Max(float64(l), float64(r)))))
	choices := make([]model.Move, 0)
	for i := 0; i < u && withinRange(u, max); i++ {
		choices = append(choices, model.Up)
	}
	for i := 0; i < d && withinRange(d, max); i++ {
		choices = append(choices, model.Down)
	}
	for i := 0; i < l && withinRange(l, max); i++ {
		choices = append(choices, model.Left)
	}
	for i := 0; i < r && withinRange(r, max); i++ {
		choices = append(choices, model.Right)
	}
	if len(choices) == 0 {
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

	for _, snek := range ws.s.otherSneks {
		if len(ws.s.me.Body) < len(snek.Body) {
			continue
		}
		for _, c := range adjacentCells(snek.Head) {
			if c == coord && ws.s.gobjs[c] != model.Body {
				return 50
			}
		}
	}

	if ws.s.gobjs[coord] == model.Food {
		return foodWeight(ws.s.me)
	} else if ws.s.gobjs[coord] == model.Body {
		return illegal
	} else if ws.isMyTail(coord) {
		return tailWeight(ws.s.me)
	}

	return 10
}

func foodWeight(me model.Battlesnake) int {
	if me.Health > 50 {
		return 20
	} else if me.Health > 30 {
		return 30
	} else if me.Health > 10 {
		return 50
	}
	return 100
}

func tailWeight(me model.Battlesnake) int {
	if me.Health > 50 {
		return myTail
	} else if me.Health > 30 {
		return myTail / 2
	}
	return 2
}

func (ws WeightedSniff) isMyTail(coord model.Coord) bool {
	if ws.s.gobjs[coord] != model.Tail {
		return false
	}
	// At the start the tail overlaps some segments, just check that it's not doing that too pls
	return coord == ws.s.me.Body[ws.s.me.Length-1] && coord != ws.s.me.Body[ws.s.me.Length-2]
}
