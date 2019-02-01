package ecs

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// Engine represents objects and world.
type Engine struct {
	// Window settings.
	title        string
	screenWidth  int32
	screenHeight int32

	// Managers
	em *EntityManager
	sm *SystemManager

	Renderer *sdl.Renderer

	closeFn func()
}

// New engine constructor.
func New(title string, w, h int32) *Engine {
	e := new(Engine)
	e.title = title
	e.screenWidth = w
	e.screenHeight = h
	e.em = &EntityManager{}
	e.sm = &SystemManager{
		eventCh: make(chan sdl.Event),
	}
	return e
}

// Start engine initialize.
func (g *Engine) Start() error {
	fmt.Println("Start")
	if err := g.setupSDL(); err != nil {
		return err
	}

	return nil
}

// Run engine game loop.
func (g *Engine) Run() error {
	for {
		// Handle event loop listener.
		for evt := sdl.PollEvent(); evt != nil; evt = sdl.PollEvent() {
			switch evt.GetType() {
			case sdl.QUIT:
				fmt.Println("Quit")
				return nil
			}
		}

		// Run all systems.
		g.sm.ProcessAll(g.em)
	}
}

// AddEntity adds new entity.
func (g *Engine) AddEntity(e *Entity) {
	g.em.Add(e)
}

// AddSystems adds new system.
func (g *Engine) AddSystems(ss ...System) {
	for _, s := range ss {
		g.sm.Add(s)
	}
}

// Stop engine destroy things.
func (g *Engine) Stop() {
	fmt.Println("Stop")
	g.closeFn()
}

func (g *Engine) setupSDL() error {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return fmt.Errorf("could not init sdl: %s", err)
	}

	win, err := sdl.CreateWindow(
		g.title,
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		g.screenWidth, g.screenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		return fmt.Errorf("could not create window:: %s", err)
	}

	g.Renderer, err = sdl.CreateRenderer(win, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		return fmt.Errorf("could not create renderer: %s", err)
	}

	g.closeFn = func() {
		sdl.Quit()
		win.Destroy()
		g.Renderer.Destroy()
	}

	return nil
}
