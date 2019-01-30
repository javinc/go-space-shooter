package main

import "github.com/veandco/go-sdl2/sdl"

type controlSystem struct {
	w, h int32
}

func newControlSystem(w, h int32) *controlSystem {
	return &controlSystem{w, h}
}

func (s *controlSystem) process(ee []*entity) {
	kk := sdl.GetKeyboardState()

	for _, e := range ee {
		// Skip entity that does not have required component.
		_, ok := e.getComponent(new(input)).(*input)
		if !ok {
			continue
		}
		rect, ok := e.getComponent(new(render)).(*render)
		if !ok {
			continue
		}
		post, ok := e.getComponent(&position{}).(*position)
		if !ok {
			continue
		}

		velocity := 0.4
		if kk[sdl.SCANCODE_LEFT] == 1 && post.x > 0 {
			post.x -= velocity
		} else if kk[sdl.SCANCODE_RIGHT] == 1 && post.x < float64(s.w-rect.w) {
			post.x += velocity
		}
	}
}
