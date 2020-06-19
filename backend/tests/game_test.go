package tests

import (
	"github.com/freglyc/tsuro/game"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	rand.Seed(time.Now().UnixNano())
	os.Exit(m.Run())
}

func TestGame(t *testing.T) {
	players := 8
	size := 6

	options := tsuro.Options{
		Players: players,
		Size:    size,
		Time:    -1,
	}
	game := tsuro.NewGame("ID", options)
	if len(game.GetPlayer(game.Turn).Hand.Tiles) != 3 {
		t.Errorf("Failed to set up player hand correctly")
	}

	// Test tile rotation
	notch := game.GetPlayer(game.Turn).Hand.Tiles[0].Edges[0][0]
	game.RotateRight(game.Turn, 0)
	if notch == game.GetPlayer(game.Turn).Hand.Tiles[0].Edges[0][0] {
		t.Errorf("Failed to rotate right correctly")
	}
	game.RotateLeft(game.Turn, 0)
	if notch != game.GetPlayer(game.Turn).Hand.Tiles[0].Edges[0][0] {
		t.Errorf("Failed to rotate left correctly")
	}

	// Test place
	space := game.GetSpace(game.Turn)
	game.Place(space, game.Turn, 0)
	if !game.Board[space[0]][space[1]].Exists() {
		t.Errorf("Failed to place tile")
	}
	if len(game.GetPlayer(game.Turn).Hand.Tiles) != 2 {
		t.Errorf("Failed to remove tile from hand")
	}

	// Test UpdateGameState
	prevTurn := game.Turn
	prevNotch := game.GetPlayer(game.Turn).Token.Notch
	game.UpdateGameState()
	if game.Turn == prevTurn {
		t.Errorf("Failed to update turn")
	}
	if game.GetPlayer(prevTurn).Token.Notch == prevNotch {
		t.Errorf("Failed to update token")
	}
	if len(game.GetPlayer(prevTurn).Hand.Tiles) != 3 {
		t.Errorf("Failed to add tile from hand")
	}

	// Test reset
	game.Reset()
	if game.GetNumTileOnBoard() != 0 {
		t.Errorf("Failed to reset board")
	}
	if len(game.Deck.Tiles) != 35-players*3 {
		t.Errorf("Failed to reset deck")
	}
}

func TestTile(t *testing.T) {
	edges := [][]tsuro.Notch{{tsuro.A, tsuro.B}, {tsuro.C, tsuro.D}, {tsuro.E, tsuro.F}, {tsuro.G, tsuro.H}}
	edgesRight := [][]tsuro.Notch{{tsuro.C, tsuro.D}, {tsuro.E, tsuro.F}, {tsuro.G, tsuro.H}, {tsuro.A, tsuro.B}}
	tile := tsuro.Tile{Edges: edges}
	tile.RotateRight()
	for i := 0; i < len(tile.Edges); i++ {
		for j := 0; j < len(tile.Edges[i]); j++ {
			if tile.Edges[i][j] != edgesRight[i][j] {
				t.Errorf("Failed to rotate right")
			}
		}
	}
	tile.RotateLeft()
	for i := 0; i < len(tile.Edges); i++ {
		for j := 0; j < len(tile.Edges[i]); j++ {
			if tile.Edges[i][j] != edges[i][j] {
				t.Errorf("Failed to rotate left")
			}
		}
	}

}

func TestDeck(t *testing.T) {
	deck := tsuro.NewDeck()
	size := len(deck.Tiles)
	if size != 35 {
		t.Errorf("Original size was %d and not 35", size)
	}
	tile := deck.Draw()
	newSize := len(deck.Tiles)
	if newSize != 34 {
		t.Errorf("Size after drawing was %d and not 34", newSize)
	}
	deck.Add(tile)
	oldSize := len(deck.Tiles)
	if oldSize != 35 {
		t.Errorf("Size after adding was %d and not 35", oldSize)
	}
}
