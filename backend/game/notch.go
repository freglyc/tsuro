package tsuro

import "encoding/json"

// Notch definition - defines the eight notches in a tile
type Notch int

const (
	A Notch = iota
	B
	C
	D
	E
	F
	G
	H
	None
)

func (notch Notch) String() string {
	switch notch {
	case A:
		return "A"
	case B:
		return "B"
	case C:
		return "C"
	case D:
		return "D"
	case E:
		return "E"
	case F:
		return "F"
	case G:
		return "G"
	case H:
		return "H"
	default:
		return "None"
	}
}

func (notch *Notch) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "A":
		*notch = A
	case "B":
		*notch = B
	case "C":
		*notch = C
	case "D":
		*notch = D
	case "E":
		*notch = E
	case "F":
		*notch = F
	case "G":
		*notch = G
	case "H":
		*notch = H
	default:
		*notch = None
	}
	return nil
}

func (notch Notch) MarshalJSON() ([]byte, error) {
	return json.Marshal(notch.String())
}
