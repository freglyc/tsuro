package tsuro

// Helper functions

// Determines whether or not a player token already exists on a tile
func contains(players []Player, token *Token) bool {
	for _, player := range players {
		if player.Token != nil && player.Token.Row == token.Row && player.Token.Col == token.Col {
			return true
		}
	}
	return false
}
