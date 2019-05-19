package component

// Position component.
type Position struct {
	X, Y float64
}

// NewPosition position constructor.
func NewPosition(x, y float64) *Position {
	return &Position{x, y}
}

// Name component implementation.
func (c *Position) Name() string {
	return "position"
}
