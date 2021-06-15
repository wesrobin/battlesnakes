package engine

import (
	"github.com/wesrobin/battlesnakes/model"
	"math"
	"math/rand"
)

type WeightedSniff struct {
}

var (
	// Weights [0,200)
	myTail  = 20
	illegal = 0

	// Misc
	sniffRadius = 5
)

func (ws WeightedSniff) getMove(board model.Board) model.Move {
	cs := sniffedCoords(board)
	weights := make(map[model.Coord]int)
	for _, c := range cs {
		w := weightMyCoord(board, c)
		weights[c] = w
	}
	u, d, l, r := moveWeights(board.Snakes[0].Head, weights)

	// Don't do illegal stuff
	head := board.Snakes[0].Head
	if !legalCoord(board, head.Move(model.Up)) {
		u = 0
	}
	if !legalCoord(board, head.Move(model.Down)) {
		d = 0
	}
	if !legalCoord(board, head.Move(model.Left)) {
		l = 0
	}
	if !legalCoord(board, head.Move(model.Right)) {
		r = 0
	}

	// TODO: Bias towards turning to the side that has more space - the bigger the difference the stronger the bias

	return chooseMove(u, d, l, r)
}

func moveWeights(head model.Coord, weights map[model.Coord]int) (u, d, l, r int) {
	for c, w := range weights {
		dist := dist(head, c)
		dist = math.Pow(dist, 2) // Inverse square to the distance
		if c.X > head.X {
			r += int(float64(w) / dist)
		}
		if c.X < head.X {
			l += int(float64(w) / dist)
		}
		if c.Y > head.Y {
			u += int(float64(w) / dist)
		}
		if c.Y < head.Y {
			d += int(float64(w) / dist)
		}
	}
	return
}

func chooseMove(u, d, l, r int) model.Move {
	tot := u + d + l + r
	choose := rand.Intn(tot+1)
	if choose < u {
		return model.Up
	} else if choose <= (u + d) {
		return model.Down
	} else if choose <= (u + d + l) {
		return model.Left
	} else if choose <= (u + d + l + r) {
		return model.Right
	}
	// Ssss we die
	return model.Up
}

func sniffedCoords(board model.Board) []model.Coord {
	head := board.Snakes[0].Head
	neck := board.Snakes[0].Body[1]
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

func weightMyCoord(board model.Board, coord model.Coord) int {
	if !inBounds(board, coord) {
		return illegal
	}

	// Check is food
	if state[coord] == model.Food {
		return foodWeight(board)
	} else if state[coord] == model.Snake {
		return snakeWeight(board, coord)
	}

	return 10
}

func foodWeight(board model.Board) int {
	if board.Snakes[0].Health > 40 {
		return 4
	} else if board.Snakes[0].Health > 25 {
		return 15
	}
	return 30
}

func snakeWeight(board model.Board, coord model.Coord) int {
	if isMyTail(board, coord) {
		return myTail
	}
	return illegal
}

func isMyTail(board model.Board, coord model.Coord) bool {
	// At the start the tail overlaps some segments, just check that it's not doing that too pls
	return coord == board.Snakes[0].Body[board.Snakes[0].Length-1] && coord != board.Snakes[0].Body[board.Snakes[0].Length-2]
}
