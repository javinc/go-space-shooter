package main

import (
	"fmt"
	"image/color"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/image/colornames"
)

const (
	windowTitle  = "Shoot-em-Up"
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
		windowTitle,
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

	initBulletPool()

	// Display average fps
	frameCtr := 0
	frameTicks := 0
	frameTicker := time.NewTicker(time.Second)
	go func() {
		for range frameTicker.C {
			frameTicks++
			fmt.Println("fps:", calcAvgFps(frameCtr, frameTicks))
		}
	}()

	// Capping frame 60 per second.
	const frameCapMs = 1000 / 60

	for {
		capTicks := 0
		capTicker := time.NewTicker(time.Millisecond)
		go func() {
			for range capTicker.C {
				capTicks++
			}
		}()

		// Handle event loop listener.
		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			switch e.(type) {
			case *sdl.QuitEvent:
				fmt.Println("Quit")
				return
			}
		}

		drawBackground(r)

		if err := enemy.draw(r); err != nil {
			fmt.Println("could not draw enemy:", err)
			return
		}

		player.update()
		if err := player.draw(r); err != nil {
			fmt.Println("could not draw player:", err)
			return
		}

		for _, b := range bulletPool {
			b.update(&enemy)
			b.draw(r)
		}

		r.Present()

		frameCtr++

		if capTicks < frameCapMs {
			d := frameCapMs - capTicks
			sdl.Delay(uint32(d))
		}
		capTicker.Stop()
	}
}

func drawBackground(r *sdl.Renderer) {
	setDrawColorByColorname(r, colornames.Black)
	r.Clear()
}

func setDrawColorByColorname(r *sdl.Renderer, c color.RGBA) error {
	return r.SetDrawColor(c.R, c.G, c.B, c.A)
}

func calcAvgFps(frames, ticks int) int {
	avgFps := frames / ticks
	if avgFps > 2000000 {
		avgFps = 0
	}

	return avgFps
}
