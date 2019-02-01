package component

const (
	// InputMoveLeft left key
	InputMoveLeft = "move-left"
	// InputMoveRight right key
	InputMoveRight = "move-right"
	// InputShootBullet shoot key
	InputShootBullet = "shoot-bullet"
)

// Input component.
type Input struct {
	Map map[uint8]string
}

// NewInput input constructor.
func NewInput() *Input {
	return &Input{
		Map: map[uint8]string{},
	}
}

// Name component implementation.
func (c *Input) Name() string {
	return "input"
}
