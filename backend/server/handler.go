package server

import (
	"github.com/freglyc/tsuro/game"
	"time"
)

type GameHandler struct {
	CreatedAt time.Time   `json:"-"`    // Time created
	Game      *tsuro.Game `json:"game"` // Game
	Hub       *Hub        `json:"-"`    // Hub
	Timer     *time.Timer `json:"-"`    // Turn timer that may or may not be enabled
	End       time.Time   `json:"-"`    // Timer end time
}

func NewHandler(game *tsuro.Game, hub *Hub) *GameHandler {
	return &GameHandler{
		CreatedAt: time.Now(),
		Game:      game,
		Hub:       hub,
		Timer:     nil,
	}
}

func (handler *GameHandler) UpdateTime() {
	if handler.Game.Time > 0 {
		currentTime := int(handler.End.Sub(time.Now()).Seconds())
		if currentTime < 0 {
			handler.Game.Countdown = handler.Game.Time
		} else {
			handler.Game.Countdown = currentTime
		}
	}
}

func (handler *GameHandler) StartTimer() {
	handler.StopTimer()
	handler.Game.Countdown = handler.Game.Time
	handler.Timer = time.NewTimer(time.Duration(handler.Game.Time+2) * time.Second)
	handler.End = time.Now().Add(time.Duration(handler.Game.Time) * time.Second)
	go func() {
		<-handler.Timer.C
		// On timer end, place the first tile in hand onto the board
		var space []int
		player := handler.Game.Teams[handler.Game.GetPlayer(handler.Game.Turn)]
		if player.Plays == 0 {
			space = []int{player.Token.Row, player.Token.Col}
		} else {
			space = handler.Game.GetSpace(handler.Game.Turn)
		}
		handler.Hub.broadcast <- ClientMessage{
			Client: nil,
			Message: Message{
				GameID: handler.Game.GameID,
				Kind:   "place",
				Team:   handler.Game.Turn,
				Idx:    0,
				Row:    space[0],
				Col:    space[1],
			}}
		handler.StartTimer()
	}()
}

// Stops the game timer
func (handler *GameHandler) StopTimer() {
	if handler.Timer != nil {
		handler.Timer.Stop()
	}
}
