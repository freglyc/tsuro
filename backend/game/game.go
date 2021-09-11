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
	Players int  `json:"players"` // number of players
	Size    int  `json:"size"`    // width and height of the board
	Time    int  `json:"time"`    // timer length, -1 means no timer
	Change  bool `json:"change"`  // true means can change, false means cannot
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
		if t.Dragon && t.Token.Notch != None {
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
	index := game.GetPlayer(team)
	if index < 0 {
		return []int{}
	}
	player := game.Teams[index]
	var space []int
	// Return next space to go to
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
	// Update first time placed token
	currentPlayer := game.Teams[game.GetPlayer(game.Turn)]
	if currentPlayer.Plays == 1 && game.GetNumTileOnBoard() <= game.Options.Players {
		tile := game.Board[currentPlayer.Token.Row][currentPlayer.Token.Col]
		before := currentPlayer.Token.Notch
		after := tile.GetNotch(before)
		// Update paths
		if game.Board[currentPlayer.Token.Row][currentPlayer.Token.Col].Paths == nil {
			game.Board[currentPlayer.Token.Row][currentPlayer.Token.Col].Paths = map[string][][]Notch{}
		}
		if len(game.Board[currentPlayer.Token.Row][currentPlayer.Token.Col].Paths[currentPlayer.Color.String()]) == 0 {
			game.Board[currentPlayer.Token.Row][currentPlayer.Token.Col].Paths[currentPlayer.Color.String()] = [][]Notch{{before, after}}
		} else {
			game.Board[currentPlayer.Token.Row][currentPlayer.Token.Col].Paths[currentPlayer.Color.String()] = append(game.Board[currentPlayer.Token.Row][currentPlayer.Token.Col].Paths[currentPlayer.Color.String()], []Notch{before, after})
		}
		game.Teams[game.GetPlayer(game.Turn)].Token.Notch = after
	}
	// Update all normal paths
	for i := 0; i < len(game.Teams); i++ {
		player := game.Teams[i]
		if player.Plays > 0 && player.Token.Notch != None {
			for {
				space := game.GetSpace(player.Color)
				if len(space) == 0 || space[0] >= game.Options.Size || space[0] < 0 || space[1] >= game.Options.Size || space[1] < 0 {
					break
				}
				tile := game.Board[space[0]][space[1]]
				if tile.Exists() {
					// If collided break out of this teams update
					flag := false
					for x := 0; x < len(game.Teams)-1; x++ {
						p1 := game.Teams[x]
						for y := x + 1; y < len(game.Teams); y++ {
							p2 := game.Teams[y]
							// If same tile
							if p1.Token.Row == p2.Token.Row && p1.Token.Col == p2.Token.Col &&
								p1.Token.Row < game.Options.Size-1 && p1.Token.Row >= 0 &&
								p1.Token.Col < game.Options.Size-1 && p1.Token.Col >= 0 {
								collisionTile := game.Board[p1.Token.Row][p1.Token.Col]
								if collisionTile.GetNotch(p1.Token.Notch) == p2.Token.Notch {
									flag = true
									game.Losers = append(game.Losers, p1.Color)
									game.Losers = append(game.Losers, p2.Color)
									p1.Token.lost()
									p2.Token.lost()
								}
							}
						}
					}
					if flag {
						break
					}
					before := notchMap[player.Token.Notch]
					after := tile.GetNotch(before)
					// Update paths
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
				} else {
					break
				}
			}
		}
	}
}

// Update winner
func (game *Game) UpdateWinner() {
	var alive []Team
	var dead []Team
	for _, player := range game.Teams {
		// If lost then update token
		if player.Plays > 0 {
			switch player.Token.Notch {
			case A, B:
				if player.Token.Row == 0 {
					game.Losers = append(game.Losers, player.Color)
					player.Token.lost()
				}
			case C, D:
				if player.Token.Col == game.Options.Size-1 {
					game.Losers = append(game.Losers, player.Color)
					player.Token.lost()
				}
			case E, F:
				if player.Token.Row == game.Options.Size-1 {
					game.Losers = append(game.Losers, player.Color)
					player.Token.lost()
				}
			case G, H:
				if player.Token.Col == 0 {
					game.Losers = append(game.Losers, player.Color)
					player.Token.lost()
				}
			}
		}
		// If player still in the game
		if player.Token.Notch != None {
			alive = append(alive, player.Color)
		} else {
			dead = append(dead, player.Color)
		}
	}
	// If dragon and lost then update dragon
	dragon := game.GetDragonTeam()
	if dragon != Neutral && game.Teams[game.GetPlayer(dragon)].Token.Notch == None {
		var nextTurn = game.GetNextTurn(game.Teams[game.GetPlayer(dragon)].Color)
		playerIdx := game.GetPlayer(nextTurn)
		for game.Teams[playerIdx].Token.Notch == None {
			nextTurn = game.GetNextTurn(nextTurn)
			playerIdx = game.GetPlayer(nextTurn)
		}
		game.Teams[game.GetPlayer(dragon)].Dragon = false
		game.Teams[playerIdx].Dragon = true
	}
	// Add back tiles of dead players to hand
	for i := 0; i < len(dead); i++ {
		p := game.Teams[game.GetPlayer(dead[i])]
		for i := 0; i < len(p.Hand.Tiles); i++ {
			game.Deck.Add(p.Hand.Tiles[i])
		}
		game.Teams[game.GetPlayer(dead[i])].Hand.RemoveAll()
	}
	// Update winner
	if len(alive) == 0 {
		game.Winner = []Team{game.Losers[len(game.Losers)-1], game.Losers[len(game.Losers)-2]}
		return
	}
	if len(alive) == 1 || game.GetNumTileOnBoard() >= 35 {
		game.Winner = alive
		return
	}
}

// Update team hands
func (game *Game) UpdateHands(turn Team) {
	if len(game.Winner) > 0 {
		return
	}
	dragon := game.GetDragonTeam()
	if dragon != Neutral {
		playerIdx := game.GetPlayer(dragon)
		if len(game.Deck.Tiles) > 0 {
			game.Teams[playerIdx].Dragon = false
			game.Teams[playerIdx].Hand.Add(game.Deck.Draw())
			game.UpdateHands(game.GetNextTurn(dragon))
		}
	} else {
		var nextTurn = turn
		playerIdx := game.GetPlayer(nextTurn)
		for game.Teams[playerIdx].Token.Notch == None {
			nextTurn = game.GetNextTurn(nextTurn)
			playerIdx = game.GetPlayer(nextTurn)
		}
		if len(game.Deck.Tiles) > 0 && len(game.Teams[playerIdx].Hand.Tiles) < 3 {
			game.Teams[playerIdx].Hand.Add(game.Deck.Draw())
			game.UpdateHands(game.GetNextTurn(nextTurn))
		} else {
			game.Teams[playerIdx].Dragon = true
		}
	}
}

// Update turn
func (game *Game) UpdateTurn() {
	if len(game.Winner) > 0 {
		return
	}
	var nextTurn = game.GetNextTurn(game.GameState.Turn)
	playerIdx := game.GetPlayer(nextTurn)
	for game.Teams[playerIdx].Token.Notch == None {
		nextTurn = game.GetNextTurn(nextTurn)
		playerIdx = game.GetPlayer(nextTurn)
	}
	game.GameState.Turn = nextTurn
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
	index := game.GetPlayer(team)
	if index < 0 {
		return
	}
	player := game.Teams[index]
	if game.Turn == team && // team turn
		idx < len(player.Hand.Tiles) && // the index is in the hands bounds
		((player.Plays == 0 && space[0] == player.Token.Row && space[1] == player.Token.Col) ||
			(player.Plays != 0 && check[0] == space[0] && check[1] == space[1])) && // adjacent to team token
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
