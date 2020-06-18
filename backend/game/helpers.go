package tsuro

// Helper functions

// Determines whether or not a player token already exists
func contains(l []Player, t Token) bool {
	for _, a := range l {
		if a.Token == t {
			return true
		}
	}
	return false
}
