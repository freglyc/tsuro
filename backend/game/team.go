package tsuro

import "encoding/json"

// Team definition - there be at most eight teams
type Team int

const (
	Black Team = iota
	White
	Red
	Blue
	Green
	Yellow
	Orange
	Gray
	Neutral
)

func (team Team) String() string {
	switch team {
	case Black:
		return "Black"
	case White:
		return "White"
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	case Orange:
		return "Orange"
	case Gray:
		return "Gray"
	default:
		return "Neutral"
	}
}

func (team *Team) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "Black":
		*team = Black
	case "White":
		*team = White
	case "Red":
		*team = Red
	case "Blue":
		*team = Blue
	case "Green":
		*team = Green
	case "Yellow":
		*team = Yellow
	case "Orange":
		*team = Orange
	case "Gray":
		*team = Gray
	default:
		*team = Neutral
	}
	return nil
}

func (team Team) MarshalJSON() ([]byte, error) {
	return json.Marshal(team.String())
}
