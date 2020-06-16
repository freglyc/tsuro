package tsuro

import (
	"encoding/json"
	"time"
)

// Team structures and functions
type Team int

const (
	Black Team = iota
	White
	Red
	Blue
	Green
	Yellow
	Orange
	Gray
	Neutral
)

func (team Team) String() string {
	switch team {
	case Black:
		return "Black"
	case White:
		return "White"
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	case Orange:
		return "Orange"
	case Gray:
		return "Gray"
	default:
		return "Neutral"
	}
}

func (team *Team) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "Black":
		*team = Black
	case "White":
		*team = White
	case "Red":
		*team = Red
	case "Blue":
		*team = Blue
	case "Green":
		*team = Green
	case "Yellow":
		*team = Yellow
	case "Orange":
		*team = Orange
	case "Gray":
		*team = Gray
	default:
		*team = Neutral
	}
	return nil
}

func (team Team) MarshalJSON() ([]byte, error) {
	return json.Marshal(team.String())
}

// Notch structures and functions, represents notches of a tile
type Notch int

const (
	A Notch = iota
	B
	C
	D
	E
	F
	G
	H
	None
)

func (notch Notch) String() string {
	switch notch {
	case A:
		return "A"
	case B:
		return "B"
	case C:
		return "C"
	case D:
		return "D"
	case E:
		return "E"
	case F:
		return "F"
	case G:
		return "G"
	case H:
		return "H"
	default:
		return "None"
	}
}

func (notch *Notch) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "A":
		*notch = A
	case "B":
		*notch = B
	case "C":
		*notch = C
	case "D":
		*notch = D
	case "E":
		*notch = E
	case "F":
		*notch = F
	case "G":
		*notch = G
	case "H":
		*notch = H
	default:
		*notch = None
	}
	return nil
}

func (notch Notch) MarshalJSON() ([]byte, error) {
	return json.Marshal(notch.String())
}

// Game structures and functions
type Options struct {
	Playing int `json:"playing"` // number of playing
	Rows    int `json:"rows"`    // number of rows
	Cols    int `json:"cols"`    // number of columns
	Time    int `json:"time"`    // timer length, -1 means no timer
}

type Tile struct {
	Edges [][]Notch        `json:"edges"` // edges that define a tile
	Paths map[Team][]Notch `json:"paths"` // defines section of team path that runs through the tile
}

// Given an notch, get the resulting notch from moving through the tile.
func (tile *Tile) GetNotch(notch Notch) Notch {
	for i := 0; i < len(tile.Paths); i++ {
		edge := tile.Edges[i]
		if edge[0] == notch {
			return edge[1]
		} else if edge[1] == notch {
			return edge[0]
		}
	}
	return None
}

// Rotates the tile right
func (tile Tile) RotateRight() {
	for i := 0; i < len(tile.Paths); i++ {
		tile.Edges[i][0] = (tile.Edges[i][0] + 2) % 8
		tile.Edges[i][1] = (tile.Edges[i][1] + 2) % 8
	}
}

// Rotates the tile left
func (tile *Tile) RotateLeft() {
	for i := 0; i < len(tile.Paths); i++ {
		tile.Edges[i][0] = (tile.Edges[i][0] + 6) % 8
		tile.Edges[i][1] = (tile.Edges[i][1] + 6) % 8
	}
}

type Token struct {
	Row   int   `json:"row"`   // row location of tile that token lies on
	Col   int   `json:"col"`   // column location of tile that token lies on
	Notch Notch `json:"notch"` // where the token lies on a tile
}

type Player struct {
	Color  Team   `json:"color"`  // team color
	Token  Token  `json:"token"`  // player's token
	Hand   []Tile `json:"hand"`   // list of tiles in hand
	Dragon bool   `json:"dragon"` // whether or not has dragon tile
}

type GameState struct {
	Board   [][]Tile `json:"board"` // game board
	deck    []Tile   // list of tiles in deck
	Players []Player `json:"players"` // list of players in game
	Turn    Team     `json:"turn"`    // current team turn
	Winner  Team     `json:"winner"`  // winning team

	Started   bool `json:"started"`   // whether or not the game has started
	Countdown int  `json:"countdown"` // time left on timer at time of sending to clients
}

// TODO
func NewGameState(options Options) GameState {
	state := GameState{}
	return state
}

type Game struct {
	GameID    string    `json:"game_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	GameState
	Options
}

// TODO
func NewGame(gameID string, options Options) *Game {
	game := &Game{}
	return game
}

func (game *Game) Reset() {
	game.GameState = NewGameState(game.Options)
}

// Gets list of spaces to play tiles
func (game *Game) GetSpaces() [][]int {
	return nil
}

// Place a tile on the board
func (game *Game) Place(row, col, int, tile Tile) {}

// Update board + hand + dragon tile
func (game *Game) Update() {}

// Check for winner
func (game *Game) CheckWinner() {}

// Change turn
func (game *Game) NextTurn() {}
