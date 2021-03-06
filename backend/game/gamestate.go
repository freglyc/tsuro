package tsuro

import (
	"math/rand"
)

// GameState definition - the state of the game at a given point in time
type GameState struct {
	Board  [][]Tile `json:"board"`  // game board
	Deck   Deck     `json:"deck"`   // list of tiles in deck
	Teams  []Player `json:"teams"`  // list of players in game
	Turn   Team     `json:"turn"`   // current team turn
	Losers []Team   `json:"losers"` // losers in order of lose
	Winner []Team   `json:"winner"` // winning team

	Started   bool `json:"started"`   // whether or not the game has started
	Countdown int  `json:"countdown"` // time left on timer at time of sending to clients
}

func NewGameState(options Options) GameState {
	colors := [8]Team{Red, Yellow, Blue, Green, Orange, Purple, Pink, Turquoise}
	// Init board
	var board = make([][]Tile, options.Size)
	for i := 0; i < options.Size; i++ {
		board[i] = make([]Tile, options.Size)
	}
	// Init deck
	deck := NewDeck()
	// Init teams
	var teams = make([]Player, options.Players)
	for i := 0; i < options.Players; i++ {
		teams[i].Color = colors[i]
		teams[i].Hand = NewHand()
		teams[i].Hand.Add(deck.Draw())
		teams[i].Hand.Add(deck.Draw())
		teams[i].Hand.Add(deck.Draw())

		// get a random token not in use
		token := RandomToken(options.Size)
		for contains(teams, token) {
			token = RandomToken(options.Size)
		}
		teams[i].Token = token
		teams[i].Dragon = false
		teams[i].Plays = 0
	}
	state := GameState{
		Board:  board,
		Deck:   deck,
		Teams:  teams,
		Turn:   teams[rand.Intn(options.Players)].Color,
		Losers: []Team{},
		Winner: []Team{},

		Started:   false,
		Countdown: options.Time,
	}
	return state
}
