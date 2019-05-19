package system

import (
	"github.com/javinc/go-space-shooter"
	"github.com/javinc/go-space-shooter/component"
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
func (s *Render) Process(em *ecs.EntityManager) {
	// Draw background.
	s.r.SetDrawColor(0, 0, 0, 0)
	s.r.Clear()

	for _, e := range em.All() {
		cm := e.ComponentManager()
		if !cm.Requires("rect", "position") {
			continue
		}

		rect := cm.Get("rect").(*component.Rect)
		// Render active only.
		if !rect.Active {
			continue
		}

		pos := cm.Get("position").(*component.Position)
		s.draw(rect, pos)
	}

	s.r.Present()
}

func (s *Render) draw(rect *component.Rect, pos *component.Position) {
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
