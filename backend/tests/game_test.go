package tests

import (
	"fmt"
	"github.com/freglyc/tsuro/game"
	"math/rand"
	"testing"
	"time"
)

func TestGame(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	options := tsuro.Options{
		Players: 8,
		Size:    6,
		Time:    -1,
	}
	state := tsuro.NewGameState(options)
	fmt.Println(state.Board[0][0])
	fmt.Println(state.Teams)
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
