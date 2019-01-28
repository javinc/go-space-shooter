package main

import (
	"image/color"

	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/image/colornames"
)

// Player entity
type Player struct {
	color color.RGBA
	size  int32
	x, y  float64
}

const (
	playerSize  = 20
	playerSpeed = 0.4
)

func newPlayer() Player {
	return Player{
		color: colornames.Red,
		size:  playerSize,
		// Placing player at the bottom-mid
		x: (screenWidth - playerSize) / 2,
		y: screenHeight - playerSize,
	}
}

func (p *Player) draw(r *sdl.Renderer) error {
	rect := &sdl.Rect{
		X: int32(p.x),
		Y: int32(p.y),
		W: p.size,
		H: p.size,
	}

	if err := setDrawColorByColorname(r, p.color); err != nil {
		return err
	}
	if err := r.FillRect(rect); err != nil {
		return err
	}

	return nil
}

func (p *Player) update() error {
	// Control movement
	keys := sdl.GetKeyboardState()
	switch {
	case keys[sdl.SCANCODE_LEFT] == 1 && p.x > 0:
		p.x -= playerSpeed
	case keys[sdl.SCANCODE_RIGHT] == 1 && p.x < (screenWidth-playerSize):
		p.x += playerSpeed
	}

	return nil
}
