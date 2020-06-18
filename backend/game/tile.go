package tsuro

// Tile definition - the tiles added to the board
type Tile struct {
	Edges [][]Notch        `json:"edges"` // edges that define a tile
	Paths map[Team][]Notch `json:"paths"` // defines section of team path that runs through the tile
}

// Given an notch, get the resulting notch from moving through the tile.
func (tile *Tile) GetNotch(notch Notch) Notch {
	for i := 0; i < len(tile.Edges); i++ {
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
func (tile *Tile) RotateRight() {
	for i := 0; i < len(tile.Edges); i++ {
		tile.Edges[i][0] = (tile.Edges[i][0] + 2) % 8
		tile.Edges[i][1] = (tile.Edges[i][1] + 2) % 8
	}
}

// Rotates the tile left
func (tile *Tile) RotateLeft() {
	for i := 0; i < len(tile.Edges); i++ {
		tile.Edges[i][0] = (tile.Edges[i][0] + 6) % 8
		tile.Edges[i][1] = (tile.Edges[i][1] + 6) % 8
	}
}
