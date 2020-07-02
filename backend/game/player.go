package tsuro

// Player definition - what the player owns
type Player struct {
	Color  Team   `json:"color"`  // team color
	Token  *Token `json:"token"`  // player's token
	Hand   *Hand  `json:"hand"`   // list of tiles in hand
	Plays  int    `json:"plays"`  // number of tiles played
	Dragon bool   `json:"dragon"` // whether or not has dragon tile
}
