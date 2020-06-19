package tsuro

// Helper functions

// Determines whether or not a player token already exists
func contains(players []Player, token *Token) bool {
	for _, player := range players {
		if player.Token == token {
			return true
		}
	}
	return false
}
