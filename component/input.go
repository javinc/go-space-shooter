package component

// Input component.
type Input struct {
}

// NewInput input constructor.
func NewInput() *Input {
	return &Input{}
}

// Name component implementation.
func (c *Input) Name() string {
	return "input"
}
