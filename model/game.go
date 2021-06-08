package model

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

type Coord struct {
	X int `json:"x"`
	Y int `json:"y"`
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
	Up Move = 0
	Down Move = 1
	Left Move = 2
	Right Move = 3
)

func (m Move) String() string {
	return []string{"up", "down", "left", "right"}[m]
}

var PossibleMoves = []Move{Up, Down, Left, Right}