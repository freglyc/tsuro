package tsuro

import (
	"math/rand"
)

// Token definition - the player tokens placed on the board
type Token struct {
	Row   int   `json:"row"`   // row location of tile that token lies on
	Col   int   `json:"col"`   // column location of tile that token lies on
	Notch Notch `json:"notch"` // where the token lies on a tile
}

func RandomToken(size int) Token {
	notch := Notch(rand.Intn(8)) // which notch to lie on
	side := rand.Intn(size)      // which side to lie on
	var row = 0
	var col = 0
	switch notch {
	case A, B:
		row = 0
		col = side
	case C, D:
		row = side
		col = size - 1
	case E, F:
		row = size - 1
		col = side
	case G, H:
		row = side
		col = 0
	}
	return Token{
		Row:   row,
		Col:   col,
		Notch: notch,
	}
}
