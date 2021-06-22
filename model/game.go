package model

import (
	"crypto/sha256"
	"fmt"
)

type Game struct {
	ID      string `json:"id"`
	Timeout int32  `json:"timeout"`
}

type Board struct {
	Height int           `json:"height"`
	Width  int           `json:"width"`
	Food   []Coord       `json:"food"`
	Snakes []Battlesnake `json:"snakes"`
}

func (b Board) Hash() string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", b)))

	return fmt.Sprintf("%x", h.Sum(nil))
}

type Coord struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (c Coord) String() string {
	return fmt.Sprintf("{%d,%d}", c.X, c.Y)
}

func (c Coord) Move(move Move) Coord {
	switch move {
	case Up:
		return Coord{X: c.X, Y: c.Y + 1}
	case Down:
		return Coord{X: c.X, Y: c.Y - 1}
	case Left:
		return Coord{X: c.X - 1, Y: c.Y}
	case Right:
		return Coord{X: c.X + 1, Y: c.Y}
	}
	panic("Ssss")
}

type Battlesnake struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Health int32   `json:"health"`
	Body   []Coord `json:"body"`
	Head   Coord   `json:"head"`
	Length int32   `json:"length"`
	Shout  string  `json:"shout"`
}

type Move int

var (
	Up    Move = 0
	Down  Move = 1
	Left  Move = 2
	Right Move = 3
)

var PossibleMoves = map[Move]string{
	Up:    "up",
	Down:  "down",
	Left:  "left",
	Right: "right",
}

type GameObject int

var (
	Nothing GameObject = 0
	Body    GameObject = 1
	Food    GameObject = 2
	Tail    GameObject = 3
	Head    GameObject = 4
)
