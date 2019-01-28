package main

import (
	"image/color"

	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/image/colornames"
)

// Enemy entity
type Enemy struct {
	color color.RGBA
	size  int32
	x, y  float64
}

const (
	enemySize = 60
)

func newEnemy() Enemy {
	return Enemy{
		color: colornames.White,
		size:  enemySize,
		x:     (screenWidth - enemySize) / 2,
		y:     0,
	}
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
	if err := r.DrawRect(rect); err != nil {
		return err
	}
	if err := r.FillRect(rect); err != nil {
		return err
	}

	return nil
}
