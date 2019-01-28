package main

import (
	"image/color"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/image/colornames"
)

// Player entity
type Player struct {
	position
	color     color.RGBA
	size      int32
	lastFired time.Time
}

const (
	playerSize = 20
	// Player movement speed.
	playerSpeed = 0.4
	// Number of bullets fired per second.
	playerFireRate = 12
)

func newPlayer() (p Player) {
	p.color = colornames.Red
	p.size = playerSize
	// Place player at the bottom-mid of the screen.
	p.x = (screenWidth - playerSize) / 2
	p.y = screenHeight - playerSize

	return
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
	kk := sdl.GetKeyboardState()

	// Control movements.
	p.movement(kk)

	// Shoot bullets.
	p.shoot(kk)

	return nil
}

func (p *Player) movement(kk []uint8) {
	if kk[sdl.SCANCODE_LEFT] == 1 && p.x > 0 {
		p.x -= playerSpeed
	} else if kk[sdl.SCANCODE_RIGHT] == 1 && p.x < (screenWidth-playerSize) {
		p.x += playerSpeed
	}
}

func (p *Player) shoot(kk []uint8) {
	if kk[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(p.lastFired) >= (time.Second / playerFireRate) {
			if b, ok := bulletFromPool(); ok {
				b.active = true
				// Place bullet on top of the player.
				b.x = p.x + float32(b.size/2)
				b.y = p.y - float32(p.size/2)
			}

			p.lastFired = time.Now()
		}
	}
}
