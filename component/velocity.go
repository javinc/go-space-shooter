package component

// Velocity component.
type Velocity struct {
	Speed float64
}

// NewVelocity velocity constructor.
func NewVelocity(speed float64) *Velocity {
	return &Velocity{speed}
}

// Name component implementation.
func (c *Velocity) Name() string {
	return "velocity"
}
