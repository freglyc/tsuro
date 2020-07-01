package tsuro

import "encoding/json"

// Team definition - there be at most eight teams
type Team int

const (
	Red Team = iota
	Blue
	Green
	Yellow
	Orange
	Purple
	Pink
	Turquoise
	Neutral
)

func (team Team) String() string {
	switch team {
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
	case Purple:
		return "Purple"
	case Pink:
		return "Pink"
	case Turquoise:
		return "Turquoise"
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
	case "Purple":
		*team = Purple
	case "Pink":
		*team = Pink
	case "Turquoise":
		*team = Turquoise
	default:
		*team = Neutral
	}
	return nil
}

func (team Team) MarshalJSON() ([]byte, error) {
	return json.Marshal(team.String())
}
