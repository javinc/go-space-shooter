package main

import (
	"fmt"
	"image/color"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
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

	if err := ttf.Init(); err != nil {
		fmt.Println("could not initialize TTF:", err)
		return
	}
	defer ttf.Quit()

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
	frameAvg := 0
	frameCtr := 0
	frameTicks := 0
	frameTicker := time.NewTicker(time.Second)
	go func() {
		for range frameTicker.C {
			frameTicks++

			avgFps := frameCtr / frameTicks
			frameAvg = avgFps
		}
	}()

	// Capping frame 60 per second.
	const frameCapMs = 1000 / 60

	var capTicks int
	capTicker := time.NewTicker(time.Millisecond)
	go func() {
		for range capTicker.C {
			capTicks++
		}
	}()

	for {
		capTicks = 0

		// Handle event loop listener.
		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			switch e.(type) {
			case *sdl.QuitEvent:
				fmt.Println("Quit")
				return
			}
		}

		drawBackground(r)

		text := fmt.Sprintf("%d", frameAvg)
		if err := drawText(r, text); err != nil {
			fmt.Println(err)
		}

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
			sdl.Delay(uint32(frameCapMs - capTicks))
		}
	}
}

func drawBackground(r *sdl.Renderer) {
	setDrawColorByColorname(r, colornames.Black)
	r.Clear()
}

func setDrawColorByColorname(r *sdl.Renderer, c color.RGBA) error {
	return r.SetDrawColor(c.R, c.G, c.B, c.A)
}

func drawText(r *sdl.Renderer, text string) error {
	f, err := ttf.OpenFont("res/fonts/flappy.ttf", 20)
	if err != nil {
		return fmt.Errorf("could not load font: %v", err)
	}
	defer f.Close()

	c := sdl.Color{R: 0, G: 255, B: 0, A: 255}
	s, err := f.RenderUTF8Solid(text, c)
	if err != nil {
		return fmt.Errorf("could not render title: %v", err)
	}
	defer s.Free()

	t, err := r.CreateTextureFromSurface(s)
	if err != nil {
		return fmt.Errorf("could not create texture: %v", err)
	}
	defer t.Destroy()

	if err := r.Copy(t, nil, &sdl.Rect{
		W: 20,
		H: 20,
		X: screenWidth - 30,
		Y: 10,
	}); err != nil {
		return fmt.Errorf("could not copy texture: %v", err)
	}

	return nil
}

func calcAvgFps(frames, ticks int) int {
	avgFps := frames / ticks
	if avgFps > 2000000 {
		avgFps = 0
	}

	return avgFps
}
