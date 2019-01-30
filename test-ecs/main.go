package main

import "fmt"

const (
	screenW = 400
	screenH = 600
)

var w *world

func main() {
	w = newWorld("Testing ECS", screenW, screenH)

	// Register entities.
	// w.addEntity(newBackground(screenW, screenH))
	w.addEntity(newPlayer())
	w.addEntity(newEnemy())

	// Setup SDL.
	r, err := w.setupSDL()
	if err != nil {
		fmt.Println("could not set SDL:", err)
		return
	}

	// Register systems.
	w.addSystem(newRenderSystem(r))
	w.addSystem(newControlSystem(screenW, screenH))

	w.run()
}
