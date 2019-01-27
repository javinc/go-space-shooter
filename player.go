package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

// Player entity
type Player struct {
	W int32
	H int32
}

func (p *Player) draw(r *sdl.Renderer) {
	const playerSize = 20
	playerRect := &sdl.Rect{
		X: (screenWidth - playerSize) / 2,
		Y: screenHeight - playerSize,
		W: playerSize,
		H: playerSize,
	}

	err := r.FillRect(playerRect)
	if err != nil {
		log.Println("could not draw player rect:", err)
	}
}

func (p *Player) update() {}

func drawPlayer() {

}
