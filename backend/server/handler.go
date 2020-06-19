package server

import (
	"github.com/freglyc/tsuro/game"
	"time"
)

type GameHandler struct {
	Clients map[*Client]bool // Holds mapping of clients subscribed to the game
	Game    *tsuro.Game      // Game

	Timer *time.Timer // Turn timer that may or may not be enabled
	End   time.Time   // Timer end time
}

func NewHandler(game *tsuro.Game) *GameHandler {
	return &GameHandler{
		Clients: make(map[*Client]bool),
		Game:    game,

		Timer: nil,
		End:   nil,
	}
}
