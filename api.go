package main

import (
	"encoding/json"
	"fmt"
	"github.com/wesrobin/battlesnakes/engine"
	"github.com/wesrobin/battlesnakes/model"
	"log"
	"net/http"
)

// HandleIndex is called when your Battlesnake is created and refreshed
// by play.battlesnake.com. BattlesnakeInfoResponse contains information about
// your Battlesnake, including what it should look like on the game board.
func HandleIndex(w http.ResponseWriter, r *http.Request) {
	response := model.BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "Yung Snek V0", // TODO: Your Battlesnake username
		Color:      "#888888",      // TODO: Personalize
		Head:       "default",      // TODO: Personalize
		Tail:       "default",      // TODO: Personalize
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatal(err)
	}
}

// HandleStart is called at the start of each game your Battlesnake is playing.
// The GameRequest object contains information about the game that's about to start.
// TODO: Use this function to decide how your Battlesnake is going to look on the board.
func HandleStart(w http.ResponseWriter, r *http.Request) {
	request := model.GameRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Fatal(err)
	}

	engine.CreateState(request.Game.ID, request.Board, request.You)

	// Nothing to respond with here
	fmt.Print("\nSTART\n")
	fmt.Printf("Head: (%d,%d)\n\n", request.Board.Snakes[0].Head.X, request.Board.Snakes[0].Head.Y)
}

// HandleMove is called for each turn of each game.
// Valid responses are "up", "down", "left", or "right".
// TODO: Use the information in the GameRequest object to determine your next move.
func HandleMove(w http.ResponseWriter, r *http.Request) {
	request := model.GameRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Fatal(err)
	}

	engine.UpdateState(request.Game.ID, request.Board, request.You)

	response := model.MoveResponse{
		Move: engine.GetMove(request.Game.ID, request.Board),
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatal(err)
	}
}

// HandleEnd is called when a game your Battlesnake was playing has ended.
// It's purely for informational purposes, no response required.
func HandleEnd(w http.ResponseWriter, r *http.Request) {
	request := model.GameRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Fatal(err)
	}

	engine.DeleteState(request.Game.ID)

	// Nothing to respond with here
	fmt.Print("END\n")
}
