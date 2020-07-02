package tsuro

import "math/rand"

// Deck definition - the deck of tiles that players draw from
type Deck struct {
	Tiles []Tile `json:"tiles"`
}

func (deck *Deck) Shuffle() {
	for i := 0; i < len(deck.Tiles); i++ {
		r := rand.Intn(len(deck.Tiles))
		if i != r {
			deck.Tiles[r], deck.Tiles[i] = deck.Tiles[i], deck.Tiles[r]
		}
	}
}

func (deck *Deck) Add(tile Tile) {
	deck.Tiles = append(deck.Tiles, tile)
	deck.Shuffle()
}

func (deck *Deck) Draw() Tile {
	size := len(deck.Tiles)
	tile := deck.Tiles[size-1]
	deck.Tiles = deck.Tiles[:size-1]
	return tile
}

func NewDeck() Deck {
	tiles := []Tile{
		{Edges: [][]Notch{{A, B}, {C, D}, {E, F}, {G, H}}},
		{Edges: [][]Notch{{A, H}, {B, G}, {C, D}, {E, F}}},
		{Edges: [][]Notch{{A, H}, {B, C}, {D, G}, {E, F}}},
		{Edges: [][]Notch{{A, H}, {B, C}, {D, E}, {F, G}}},
		{Edges: [][]Notch{{A, G}, {B, H}, {C, D}, {E, F}}},
		{Edges: [][]Notch{{A, B}, {C, H}, {D, G}, {E, F}}},
		{Edges: [][]Notch{{A, B}, {C, G}, {D, H}, {E, F}}},
		{Edges: [][]Notch{{A, G}, {B, C}, {D, H}, {E, F}}},
		{Edges: [][]Notch{{A, B}, {C, G}, {D, E}, {F, H}}},
		{Edges: [][]Notch{{A, G}, {B, C}, {D, E}, {F, H}}},
		{Edges: [][]Notch{{A, C}, {B, G}, {D, E}, {F, H}}},
		{Edges: [][]Notch{{A, C}, {B, G}, {D, H}, {E, F}}},
		{Edges: [][]Notch{{A, C}, {B, H}, {D, G}, {E, F}}},
		{Edges: [][]Notch{{A, D}, {B, H}, {C, G}, {E, F}}},
		{Edges: [][]Notch{{A, D}, {B, G}, {C, H}, {E, F}}},
		{Edges: [][]Notch{{A, D}, {B, C}, {E, H}, {F, G}}},
		{Edges: [][]Notch{{A, D}, {B, C}, {E, G}, {F, H}}},
		{Edges: [][]Notch{{A, E}, {B, C}, {D, G}, {F, H}}},
		{Edges: [][]Notch{{A, E}, {B, C}, {D, H}, {F, G}}},
		{Edges: [][]Notch{{A, F}, {B, H}, {C, D}, {E, G}}},
		{Edges: [][]Notch{{A, F}, {B, G}, {C, H}, {D, E}}},
		{Edges: [][]Notch{{A, F}, {B, C}, {D, H}, {E, G}}},
		{Edges: [][]Notch{{A, F}, {B, D}, {C, H}, {E, G}}},
		{Edges: [][]Notch{{A, F}, {B, D}, {C, G}, {E, H}}},
		{Edges: [][]Notch{{A, E}, {B, D}, {C, G}, {F, H}}},
		{Edges: [][]Notch{{A, C}, {B, D}, {E, G}, {F, H}}},
		{Edges: [][]Notch{{A, F}, {B, E}, {C, H}, {D, G}}},
		{Edges: [][]Notch{{A, F}, {B, E}, {C, G}, {D, H}}},
		{Edges: [][]Notch{{A, E}, {B, F}, {C, G}, {D, H}}},
		{Edges: [][]Notch{{A, D}, {B, F}, {C, G}, {E, H}}},
		{Edges: [][]Notch{{A, D}, {B, F}, {C, H}, {E, G}}},
		{Edges: [][]Notch{{A, C}, {B, F}, {D, H}, {E, G}}},
		{Edges: [][]Notch{{A, D}, {B, G}, {C, E}, {F, H}}},
		{Edges: [][]Notch{{A, G}, {B, D}, {C, E}, {F, H}}},
		{Edges: [][]Notch{{A, D}, {B, G}, {C, F}, {E, H}}},
	}
	deck := Deck{
		Tiles: tiles,
	}
	deck.Shuffle()
	return deck
}
