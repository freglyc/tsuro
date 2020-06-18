package tsuro

import (
	"time"
)

// Game definition - the game itself
type Game struct {
	GameID    string    `json:"game_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	GameState
	Options
}

type Options struct {
	Players int `json:"players"` // number of players
	Size    int `json:"size"`    // width and height of the board
	Time    int `json:"time"`    // timer length, -1 means no timer
}

func NewGame(gameID string, options Options) *Game {
	game := &Game{
		GameID:    gameID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		GameState: NewGameState(options),
		Options:   options,
	}
	return game
}

// GETTER FUNCTIONS

// Get a player given a team
func (game *Game) GetPlayer(team Team) Player {
	var player Player
	for _, t := range game.Teams {
		if t.Color == team {
			player = t
			break
		}
	}
	return player
}

// Get the team with the dragon tile, Neutral otherwise
func (game *Game) GetDragonTeam() Team {
	for _, t := range game.Teams {
		if t.Dragon {
			return t.Color
		}
	}
	return Neutral
}

// Get the number of tiles needed to fill every hand of every team still in the game
func (game *Game) GetNumToFillHands() int {
	total := 0
	for _, t := range game.Teams {
		// If still in the game add number of tiles missing to total
		if t.Token.Notch != None {
			if len(t.Hand.Tiles) < 3 {
				total += 3 - len(t.Hand.Tiles)
			}
		}
	}
	return total
}

// Get the space that the player can team can play a tile ASSUMING the makers have been updated
func (game *Game) GetSpace(team Team) []int {
	player := game.GetPlayer(team)
	var space []int
	switch player.Token.Notch {
	case A, B:
		space = []int{player.Token.Row + 1, player.Token.Col}
	case C, D:
		space = []int{player.Token.Row, player.Token.Col + 1}
	case E, F:
		space = []int{player.Token.Row - 1, player.Token.Col}
	case G, H:
		space = []int{player.Token.Row, player.Token.Col - 1}
	default:
		space = []int{}
	}
	return space
}

// Get the next players turn after the input team
func (game *Game) GetNextTurn(team Team) Team {
	for i := 0; i < len(game.GameState.Teams); i++ {
		if team == game.GameState.Teams[i].Color {
			return game.GameState.Teams[(i+1)%game.Options.Players].Color
		}
	}
	return Neutral
}

// GAME LOGIC

// Update the tokens and tile path segments
func (game *Game) UpdateTokens() {}

// Update winner
func (game *Game) UpdateWinner() {
	alive := 0
	for _, player := range game.Teams {
		if player.Token.Notch != None {
			alive += 1
		}
	}
	// TODO
}

// Update team hands
func (game *Game) UpdateHands(turn Team) {
	if turn == Neutral {
		return
	}
	dragon := game.GetDragonTeam()
	var active Player
	// Give the first draw to the dragon player
	if dragon != Neutral {
		active = game.GetPlayer(dragon)
		active.Dragon = false
	} else {
		active = game.GetPlayer(turn)
	}

	if len(game.deck.Tiles) > 0 {
		// Add tile to team hand if needed
		if len(active.Hand.Tiles) < 3 {
			active.Hand.Add(game.deck.Draw())
		}
		// If other players do not have full hands fill them
		if len(game.deck.Tiles) > 0 && game.GetNumToFillHands() > 0 {
			game.UpdateHands(game.GetNextTurn(turn))
		}
	} else if game.Options.Players > 2 {
		for _, player := range game.Teams {
			player.Dragon = false
		}
		active.Dragon = true
	}
}

// Update turn
func (game *Game) UpdateTurn() {
	game.GameState.Turn = game.GetNextTurn(game.Turn)
}

// Updates game state after a team places a tile
func (game *Game) UpdateGameState() {
	game.UpdateTokens() // moves all tokens to correct position
	game.UpdateWinner() // updates the winner if someone has won
	if game.Winner != Neutral {
		return
	}
	game.UpdateHands(game.Turn) // updates all player hands
	game.UpdateTurn()           // updates the turn
}

// PLAYER ACTIONS

// Rotate tile at index idx in team hand
func (game *Game) RotateRight(team Team, idx int) {
	player := game.GetPlayer(team)
	if idx < len(player.Hand.Tiles) {
		player.Hand.Tiles[idx].RotateRight()
	}
}

// Rotate tile at index idx in team hand
func (game *Game) RotateLeft(team Team, idx int) {
	player := game.GetPlayer(team)
	if idx < len(player.Hand.Tiles) {
		player.Hand.Tiles[idx].RotateLeft()
	}
}

// Place a tile on the board in an open position
func (game *Game) Place(space []int, team Team, idx int) {
	check := game.GetSpace(team)
	player := game.GetPlayer(team)
	if game.Turn == team && // team turn
		idx < len(player.Hand.Tiles) && // the index is in the hands bounds
		check[0] == space[0] && check[1] == space[1] && // adjacent to team token
		len(game.Board[space[0]][space[1]].Edges) == 0 { // there in no tile in the space

		game.Board[space[0]][space[1]] = player.Hand.Tiles[idx] // add to board
		player.Hand.Remove(idx)                                 // remove from hand
	}
}

// Resets the game
func (game *Game) Reset() {
	game.GameState = NewGameState(game.Options)
}
