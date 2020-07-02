package tsuro

import (
	"time"
)

// Game definition - the game itself
type Game struct {
	GameID    string    `json:"gameID"`
	CreatedAt time.Time `json:"created"`
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
		GameState: NewGameState(options),
		Options:   options,
	}
	return game
}

// GETTER FUNCTIONS

// Get a player given a team
func (game *Game) GetPlayer(team Team) int {
	for i := 0; i < len(game.Teams); i++ {
		if game.Teams[i].Color == team {
			return i
		}
	}
	return -1
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

func (game *Game) GetNumTileOnBoard() int {
	placed := 0
	for _, row := range game.Board {
		for _, tile := range row {
			// Real tile exists in this location
			if tile.Exists() {
				placed += 1
			}
		}
	}
	return placed
}

// Get the space that the player can team can play a tile ASSUMING the makers have been updated
func (game *Game) GetSpace(team Team) []int {
	player := game.Teams[game.GetPlayer(team)]
	var space []int
	// If first turn of the game then return current space
	if game.GetNumTileOnBoard() <= game.Options.Players &&
		((player.Token.Row == 0 && (player.Token.Notch == A || player.Token.Notch == B)) ||
			(player.Token.Row == game.Options.Size-1 && (player.Token.Notch == E || player.Token.Notch == F)) ||
			(player.Token.Col == 0 && (player.Token.Notch == G || player.Token.Notch == H)) ||
			(player.Token.Col == game.Options.Size-1 && (player.Token.Notch == C || player.Token.Notch == D))) {
		return []int{player.Token.Row, player.Token.Col}
	}
	// Otherwise return next space to go to
	switch player.Token.Notch {
	case A, B:
		space = []int{player.Token.Row - 1, player.Token.Col}
	case C, D:
		space = []int{player.Token.Row, player.Token.Col + 1}
	case E, F:
		space = []int{player.Token.Row + 1, player.Token.Col}
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
func (game *Game) UpdateTokens() {

	notchMap := map[Notch]Notch{A: F, B: E, C: H, D: G, E: B, F: A, G: D, H: C}

	for i := 0; i < len(game.Teams); i++ {
		player := game.Teams[i]
		flag := false
		count := 0
		for {
			if flag {
				break
			}
			space := game.GetSpace(player.Color)
			// if out of bounds return
			if len(space) == 0 || space[0] >= game.Options.Size || space[0] < 0 || space[1] >= game.Options.Size || space[1] < 0 {
				break
			}
			// If a tile was placed where the player token can move then update
			tile := game.Board[space[0]][space[1]]
			if tile.Exists() {
				// Update tile path
				var before Notch
				if (player.Plays == 0 || player.Plays == 1) && player.Color == game.Turn && count == 0 {
					flag = true
					before = player.Token.Notch
				} else {
					before = notchMap[player.Token.Notch]
				}
				after := tile.GetNotch(before)
				if game.Board[space[0]][space[1]].Paths == nil {
					game.Board[space[0]][space[1]].Paths = map[string][][]Notch{}
				}
				if len(game.Board[space[0]][space[1]].Paths[player.Color.String()]) == 0 {
					game.Board[space[0]][space[1]].Paths[player.Color.String()] = [][]Notch{{before, after}}
				} else {
					game.Board[space[0]][space[1]].Paths[player.Color.String()] = append(game.Board[space[0]][space[1]].Paths[player.Color.String()], []Notch{before, after})
				}

				// Update token
				game.Teams[i].Token.Row = space[0]
				game.Teams[i].Token.Col = space[1]
				game.Teams[i].Token.Notch = after

				// If collide with other player both lose
				for i := 0; i < len(game.Teams)-1; i++ {
					p1 := game.Teams[i]
					for j := i + 1; j < len(game.Teams); j++ {
						p2 := game.Teams[j]
						if p1.Token.Row == p2.Token.Row && p1.Token.Col == p2.Token.Col {
							if p1.Token.Notch == p2.Token.Notch || game.Board[p1.Token.Row][p2.Token.Col].GetNotch(p1.Token.Notch) == p2.Token.Notch {
								flag = true
							}
						}
					}
				}
			} else {
				break
			}
			count += 1
		}
	}
}

// Update winner
func (game *Game) UpdateWinner() {
	var alive []Team
	for _, player := range game.Teams {
		// TODO add collision here
		// If lost then update token
		if player.Plays > 0 {
			switch player.Token.Notch {
			case A, B:
				if player.Token.Row == 0 {
					player.Token.lost()
				}
			case C, D:
				if player.Token.Col == game.Options.Size-1 {
					player.Token.lost()
				}
			case E, F:
				if player.Token.Row == game.Options.Size-1 {
					player.Token.lost()
				}
			case G, H:
				if player.Token.Col == 0 {
					player.Token.lost()
				}
			}
		}
		// If player still in the game
		if player.Token.Notch != None {
			alive = append(alive, player.Color)
		}
	}
	if len(alive) == 0 {
		game.Winner = []Team{Neutral}
		return
	}
	if len(alive) == 1 || game.GetNumTileOnBoard() >= 35 {
		game.Winner = alive
		return
	}
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
		active = game.Teams[game.GetPlayer(dragon)]
		active.Dragon = false
	} else {
		active = game.Teams[game.GetPlayer(turn)]
	}

	if len(game.Deck.Tiles) > 0 {
		// Add tile to team hand if needed
		if len(active.Hand.Tiles) < 3 {
			active.Hand.Add(game.Deck.Draw())
		}
		// If other players do not have full hands fill them
		if len(game.Deck.Tiles) > 0 && game.GetNumToFillHands() > 0 {
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
	alive := map[Team]bool{}
	for _, player := range game.Teams {
		if player.Token.Notch != None {
			alive[player.Color] = true
		} else {
			alive[player.Color] = false
		}
	}
	count := 0
	for {
		turn := game.GetNextTurn(game.Turn)
		game.GameState.Turn = turn
		if alive[turn] || count > 8 {
			break
		}
		count += 1
	}
}

// Updates game state after a team places a tile
func (game *Game) UpdateGameState() {
	game.UpdateTokens() // moves all tokens to correct position
	game.UpdateWinner() // updates the winner if someone has won
	if len(game.Winner) > 0 {
		return
	}
	game.UpdateHands(game.Turn) // updates all player hands
	game.UpdateTurn()           // updates the turn
}

// TEAM ACTIONS

// Rotate tile at index idx in team hand
func (game *Game) RotateRight(team Team, idx int) {
	player := game.Teams[game.GetPlayer(team)]
	if idx < len(player.Hand.Tiles) {
		player.Hand.Tiles[idx].RotateRight()
	}
}

// Rotate tile at index idx in team hand
func (game *Game) RotateLeft(team Team, idx int) {
	player := game.Teams[game.GetPlayer(team)]
	if idx < len(player.Hand.Tiles) {
		player.Hand.Tiles[idx].RotateLeft()
	}
}

// Place a tile on the board in an open position
func (game *Game) Place(space []int, team Team, idx int) {
	check := game.GetSpace(team)
	player := game.Teams[game.GetPlayer(team)]
	if game.Turn == team && // team turn
		idx < len(player.Hand.Tiles) && // the index is in the hands bounds
		check[0] == space[0] && check[1] == space[1] && // adjacent to team token
		!game.Board[space[0]][space[1]].Exists() { // there in no tile in the space

		game.Board[space[0]][space[1]] = player.Hand.Tiles[idx] // add to board
		player.Hand.Remove(idx)                                 // remove from hand
		game.Started = true

		game.Teams[game.GetPlayer(team)].Plays += 1

		game.UpdateGameState()
	}
}

// Resets the game
func (game *Game) Reset() {
	game.GameState = NewGameState(game.Options)
}
