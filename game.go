package main

import (
	"time"

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
	enemyRect := &sdl.Rect{
		X: (screenWidth - enemyDim) / 2,
		Y: 0,
		W: enemyDim,
		H: enemyDim,
	}
	surface.FillRect(enemyRect, 0xffffffff)

	// Player dimension size
	const playerSize = 20
	playerRect := &sdl.Rect{
		X: (screenWidth - playerSize) / 2,
		Y: screenHeight - playerSize,
		W: playerSize,
		H: playerSize,
	}
	surface.FillRect(playerRect, 0xffff0000)

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

		keys := sdl.GetKeyboardState()
		if keys[sdl.SCANCODE_LEFT] == 1 {
			playerRect.X++
			println("LEFT!!!!", playerRect.X)
			time.Sleep(time.Millisecond)
			window.UpdateSurface()

		} else if keys[sdl.SCANCODE_RIGHT] == 1 {
			playerRect.X--
			println("RIGHT!!!!", playerRect.X)
			time.Sleep(time.Millisecond)
			window.UpdateSurface()
		}
	}
}
