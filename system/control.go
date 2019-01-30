package system

import (
	"github.com/javinc/ecs"
	"github.com/javinc/ecs/component"
	"github.com/veandco/go-sdl2/sdl"
)

// Control system.
type Control struct {
	w, h int32
}

// NewControl Control system constructor.
func NewControl(w, h int32) *Control {
	return &Control{w, h}
}

// Process Control system implements System interface.
func (s *Control) Process(ee []*ecs.Entity) {
	kk := sdl.GetKeyboardState()

	for _, e := range ee {
		if !e.Requires("input", "velocity", "rect", "position") {
			continue
		}

		rect := e.Get("rect").(*component.Rect)
		pos := e.Get("position").(*component.Position)
		vel := e.Get("velocity").(*component.Velocity)

		if kk[sdl.SCANCODE_LEFT] == 1 && pos.X > 0 {
			pos.X -= vel.Speed
		} else if kk[sdl.SCANCODE_RIGHT] == 1 && pos.X < float64(s.w-rect.W) {
			pos.X += vel.Speed
		}
	}
}
