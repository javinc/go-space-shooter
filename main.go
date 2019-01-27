package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 450
	screenHeight = 600
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("shoot-em-up", sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED, screenWidth, screenHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)

	// Enemy dimension size
	const enemyDim = 60
	surface.FillRect(&sdl.Rect{
		X: (screenWidth - enemyDim) / 2,
		Y: 0,
		W: enemyDim,
		H: enemyDim,
	}, 0xffffffff)
	// Player dimension size
	const playerSize = 20
	surface.FillRect(&sdl.Rect{
		X: (screenWidth - playerSize) / 2,
		Y: screenHeight - playerSize,
		W: playerSize,
		H: playerSize,
	}, 0xffff0000)

	window.UpdateSurface()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}
	}
}
