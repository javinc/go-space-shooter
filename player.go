package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// Player entity
type Player struct {
	size int32
	x, y float64
}

const (
	playerSize  = 20
	playerSpeed = 0.5
)

func newPlayer() Player {
	return Player{
		size: playerSize,
		x:    (screenWidth - playerSize) / 2,
		y:    screenHeight - playerSize,
	}
}

func (p *Player) draw(r *sdl.Renderer) {
	rect := &sdl.Rect{
		X: int32(p.x),
		Y: int32(p.y),
		W: p.size,
		H: p.size,
	}

	r.SetDrawColor(100, 0, 0, 0)
	err := r.DrawRect(rect)
	if err != nil {
		fmt.Println("could not draw player rect:", err)
	}
	err = r.FillRect(rect)
	if err != nil {
		fmt.Println("could not fill player rect:", err)
	}
}

func (p *Player) update() {
	keys := sdl.GetKeyboardState()
	if keys[sdl.SCANCODE_LEFT] == 1 {
		p.x -= playerSpeed
		println("LEFT!!!!", p.x)

	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		p.x += playerSpeed
		println("RIGHT!!!!", p.x)
	}
}
