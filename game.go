package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 450
	screenHeight = 600
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("could not init sdl:", err)
		return
	}
	defer sdl.Quit()

	w, err := sdl.CreateWindow(
		"shoot-em-up",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("could not create window:", err)
		return
	}
	defer w.Destroy()

	r, err := sdl.CreateRenderer(w, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("could not create renderer:", err)
		return
	}
	defer r.Destroy()

	player := newPlayer()
	enemy := newEnemy()

	for {
		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			switch e.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				return
			}
		}

		drawBackground(r)
		r.Clear()

		enemy.draw(r)

		player.draw(r)
		player.update()

		r.Present()
	}
}

func drawBackground(r *sdl.Renderer) {
	r.SetDrawColor(0, 0, 0, 0)
}
