package tsuro

// Hand definition - a hand of tiles
type Hand struct {
	Tiles []Tile `json:"tiles"`
}

func NewHand() Hand {
	return Hand{
		Tiles: []Tile{},
	}
}

func (hand *Hand) Add(tile Tile) {
	hand.Tiles = append(hand.Tiles, tile)
}

func (hand *Hand) Remove(idx int) {
	if idx < len(hand.Tiles) {
		copy(hand.Tiles[idx:], hand.Tiles[idx+1:])
		hand.Tiles[len(hand.Tiles)-1] = Tile{}
		hand.Tiles = hand.Tiles[:len(hand.Tiles)-1]
	}
}
