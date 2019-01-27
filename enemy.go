package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// Enemy entity
type Enemy struct {
	W int32
	H int32
}

func drawEnemy(r *sdl.Renderer) {
	const size = 60
	rect := &sdl.Rect{
		X: (screenWidth - size) / 2,
		Y: 0,
		W: size,
		H: size,
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
