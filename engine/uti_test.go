package engine

import (
	"github.com/wesrobin/battlesnakes/model"
	"gotest.tools/assert"
	"testing"
)

func TestStep(t *testing.T) {
	testCases := []struct {
		name string
		board model.Board
		move model.Move
		expected model.Board
	} {
		{
			name: "Up",
			board: startBoard,
			move: model.Up,
			expected: model.Board{
				Height: 3,
				Width: 3,
				Snakes: []model.Battlesnake{
					{
						Health: 99,
						Length: 2,
						Body: []model.Coord{
							{
								X: 1,
								Y: 2,
							},
							{
								X: 1,
								Y: 1,
							},
						},
						Head: model.Coord{
							X: 1,
							Y: 2,
						},
					},
				},

			},
		},
		{
			name: "Down",
			board: startBoard,
			move: model.Down,
			expected: model.Board{
				Height: 3,
				Width: 3,
				Snakes: []model.Battlesnake{
					{
						Health: 99,
						Length: 2,
						Body: []model.Coord{
							{
								X: 1,
								Y: 0,
							},
							{
								X: 1,
								Y: 1,
							},
						},
						Head: model.Coord{
							X: 1,
							Y: 0,
						},
					},
				},

			},
		},
		{
			name: "Left",
			board: startBoard,
			move: model.Left,
			expected: model.Board{
				Height: 3,
				Width: 3,
				Snakes: []model.Battlesnake{
					{
						Health: 99,
						Length: 2,
						Body: []model.Coord{
							{
								X: 0,
								Y: 1,
							},
							{
								X: 1,
								Y: 1,
							},
						},
						Head: model.Coord{
							X: 0,
							Y: 1,
						},
					},
				},

			},
		},
		{
			name: "Right",
			board: startBoard,
			move: model.Right,
			expected: model.Board{
				Height: 3,
				Width: 3,
				Snakes: []model.Battlesnake{
					{
						Health: 99,
						Length: 2,
						Body: []model.Coord{
							{
								X: 2,
								Y: 1,
							},
							{
								X: 1,
								Y: 1,
							},
						},
						Head: model.Coord{
							X: 2,
							Y: 1,
						},
					},
				},

			},
		},
		{
			name: "Right, ate food",
			board: model.Board{
				Height: 3,
				Width: 3,
				Food: []model.Coord{
					{
						X: 2,
						Y: 1,
					},
				},
				Snakes: []model.Battlesnake{
					{
						Health: 100,
						Length: 2,
						Body: []model.Coord{
							{
								X: 1,
								Y: 1,
							},
							{
								X: 0,
								Y: 1,
							},
						},
						Head: model.Coord{
							X: 1,
							Y: 1,
						},
					},
				},
			},
			move: model.Right,
			expected: model.Board{
				Height: 3,
				Width: 3,
				Food: []model.Coord{},
				Snakes: []model.Battlesnake{
					{
						Health: 100,
						Length: 3,
						Body: []model.Coord{
							{
								X: 2,
								Y: 1,
							},
							{
								X: 1,
								Y: 1,
							},
							{
								X: 0,
								Y: 1,
							},
						},
						Head: model.Coord{
							X: 2,
							Y: 1,
						},
					},
				},

			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := step(tc.board, tc.move)
			assert.DeepEqual(t, tc.expected, actual)
		})
	}
}

var startBoard = model.Board{
	Height: 3,
	Width: 3,
	Snakes: []model.Battlesnake{
		{
			Health: 100,
			Length: 2,
			Body: []model.Coord{
				{
					X: 1,
					Y: 1,
				},
				{
					X: 0,
					Y: 1,
				},
			},
			Head: model.Coord{
				X: 1,
				Y: 1,
			},
		},
	},
}
