package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type world struct {
	title        string
	screenWidth  int32
	screenHeight int32

	entities []*entity
	systems  []system

	destroyFn func()
}

type system interface {
	process([]*entity)
}

func newWorld(t string, w, h int32) *world {
	return &world{
		title:        t,
		screenWidth:  w,
		screenHeight: h,
	}
}

func (w *world) addSystem(s system) {
	w.systems = append(w.systems, s)
}

func (w *world) addEntity(e *entity) {
	w.entities = append(w.entities, e)
}

func (w *world) run() error {
	defer w.destroy()
	for {
		// Handle event loop listener.
		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			switch e.(type) {
			case *sdl.QuitEvent:
				fmt.Println("Quit")
				return nil
			}
		}

		// Run all systems.
		for _, s := range w.systems {
			s.process(w.entities)
		}
	}
}

func (w *world) setupSDL() (*sdl.Renderer, error) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return nil, fmt.Errorf("could not init sdl: %s", err)
	}

	win, err := sdl.CreateWindow(
		w.title,
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		w.screenWidth, w.screenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		return nil, fmt.Errorf("could not create window:: %s", err)
	}

	r, err := sdl.CreateRenderer(win, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return nil, fmt.Errorf("could not create renderer: %s", err)
	}

	w.destroyFn = func() {
		sdl.Quit()
		win.Destroy()
		r.Destroy()
	}

	return r, nil
}

func (w *world) destroy() {
	w.destroyFn()
}
