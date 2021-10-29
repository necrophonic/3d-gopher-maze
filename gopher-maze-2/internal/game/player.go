package game

const (
	N orientation = iota
	S
	E
	W
)

type (
	orientation byte

	// Player represents the current player state
	Player struct {
		Position    Point
		Orientation orientation
	}
)
