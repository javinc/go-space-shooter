package main

import (
	"image/color"

	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/image/colornames"
)

// Enemy entity
type Enemy struct {
	position
	color color.RGBA
	size  int32
}

const (
	enemySize = 60
)

func newEnemy() (e Enemy) {
	e.color = colornames.White
	e.size = enemySize
	// Placing enemy at the top-mid of the screen.
	e.x = (screenWidth - enemySize) / 2
	e.y = 0
	return
}

func (e *Enemy) draw(r *sdl.Renderer) error {
	rect := &sdl.Rect{
		X: int32(e.x),
		Y: int32(e.y),
		W: e.size,
		H: e.size,
	}

	if err := setDrawColorByColorname(r, e.color); err != nil {
		return err
	}
	if err := r.FillRect(rect); err != nil {
		return err
	}

	return nil
}
