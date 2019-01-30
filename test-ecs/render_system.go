package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type renderSystem struct {
	render *sdl.Renderer
}

func newRenderSystem(r *sdl.Renderer) *renderSystem {
	return &renderSystem{r}
}

func (s *renderSystem) process(ee []*entity) {
	// Draw background.
	s.render.SetDrawColor(0, 0, 0, 0)
	s.render.Clear()

	for _, e := range ee {
		// Skip entity that does not have required component.
		rect, ok := e.getComponent(new(render)).(*render)
		if !ok {
			return
		}
		post, ok := e.getComponent(new(position)).(*position)
		if !ok {
			return
		}

		s.render.SetDrawColor(
			rect.color.R,
			rect.color.G,
			rect.color.B,
			rect.color.A,
		)
		s.render.FillRect(&sdl.Rect{
			W: rect.w, H: rect.h,
			X: int32(post.x), Y: int32(post.y),
		})
	}

	s.render.Present()
}
