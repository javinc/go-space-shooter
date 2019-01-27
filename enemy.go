package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// Enemy entity
type Enemy struct {
	size int32
	x, y float64
}

const (
	enemySize = 60
)

func newEnemy() Enemy {
	return Enemy{
		size: enemySize,
		x:    (screenWidth - enemySize) / 2,
		y:    0,
	}
}

func (e *Enemy) draw(r *sdl.Renderer) {
	rect := &sdl.Rect{
		X: int32(e.x),
		Y: int32(e.y),
		W: e.size,
		H: e.size,
	}

	r.SetDrawColor(255, 255, 255, 0)
	err := r.DrawRect(rect)
	if err != nil {
		fmt.Println("could not draw enemy rect:", err)
	}
	err = r.FillRect(rect)
	if err != nil {
		fmt.Println("could not fill enemy rect:", err)
	}
}
