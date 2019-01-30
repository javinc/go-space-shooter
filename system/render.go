package system

import (
	"github.com/javinc/ecs"
	"github.com/javinc/ecs/component"
	"github.com/veandco/go-sdl2/sdl"
)

// Render system.
type Render struct {
	r *sdl.Renderer
}

// NewRender Render system constructor.
func NewRender(r *sdl.Renderer) *Render {
	return &Render{r}
}

// Process Render system implements System interface.
func (s *Render) Process(ee []*ecs.Entity) {
	// Draw background.
	s.r.SetDrawColor(0, 0, 0, 0)
	s.r.Clear()

	for _, e := range ee {
		if !e.Requires("rect", "position") {
			continue
		}

		rect := e.Get("rect").(*component.Rect)
		pos := e.Get("position").(*component.Position)

		s.r.SetDrawColor(
			rect.Color.R,
			rect.Color.G,
			rect.Color.B,
			rect.Color.A,
		)
		s.r.FillRect(&sdl.Rect{
			W: rect.W, H: rect.H,
			X: int32(pos.X), Y: int32(pos.Y),
		})
	}

	s.r.Present()
}
