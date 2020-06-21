package server

import (
	"github.com/freglyc/tsuro/game"
	"time"
)

type GameHandler struct {
	Game    *tsuro.Game      `json:"game"` // Game
	Clients map[*Client]bool `json:"-"`    // Holds mapping of clients subscribed to the game
	Timer   *time.Timer      `json:"-"`    // Turn timer that may or may not be enabled
	End     time.Time        `json:"-"`    // Timer end time
}

func NewHandler(game *tsuro.Game) *GameHandler {
	return &GameHandler{
		Clients: make(map[*Client]bool),
		Game:    game,
		Timer:   nil,
	}
}
