package ecs

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// Engine represents objects and world.
type Engine struct {
	title        string
	screenWidth  int32
	screenHeight int32

	Renderer *sdl.Renderer

	EntityManager
	ComponentManager
	SystemManager

	closeFn func()
}

// New engine constructor.
func New(title string, w, h int32) *Engine {
	e := new(Engine)
	e.title = title
	e.screenWidth = w
	e.screenHeight = h
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
		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			switch e.(type) {
			case *sdl.QuitEvent:
				fmt.Println("Quit")
				return nil
			}
		}

		// Run all systems.
		g.SystemManager.Process(g.EntityManager.All())
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

	g.Renderer, err = sdl.CreateRenderer(win, -1, sdl.RENDERER_ACCELERATED)
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
