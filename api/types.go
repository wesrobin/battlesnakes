package api

import "github.com/wesrobin/battlesnakes/model"

type BattlesnakeInfoResponse struct {
	APIVersion string `json:"apiversion"`
	Author     string `json:"author"`
	Color      string `json:"color"`
	Head       string `json:"head"`
	Tail       string `json:"tail"`
}

type GameRequest struct {
	Game  model.Game        `json:"game"`
	Turn  int               `json:"turn"`
	Board model.Board       `json:"board"`
	You   model.Battlesnake `json:"you"`
}

type MoveResponse struct {
	Move  string `json:"move"`
	Shout string `json:"shout,omitempty"`
}
