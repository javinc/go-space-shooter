package main

import (
	"github.com/javinc/ecs"
	"github.com/javinc/ecs/component"
	"github.com/javinc/ecs/system"
	"golang.org/x/image/colornames"
)

const (
	title        = "ECS TEST"
	screenWidth  = 450
	screenHeight = 600
)

func main() {
	engine := ecs.New(title, screenWidth, screenHeight)
	engine.Start()

	// Register entities.
	engine.EntityManager.Add(newPlayer())
	engine.EntityManager.Add(newEnemy())

	// Register systems.
	engine.SystemManager.Add(system.NewControl(screenWidth, screenHeight))
	engine.SystemManager.Add(system.NewRender(engine.Renderer))

	engine.Run()
	engine.Stop()
}

// returns player composition.
func newPlayer() *ecs.Entity {
	const size = 20
	e := ecs.NewEntity()
	e.Add(component.NewRect(colornames.Red, size, size))
	// Place player at the bottom-mid of the screen.
	e.Add(component.NewPosition((screenWidth-size)/2, screenHeight-size))
	e.Add(component.NewVelocity(0.5))
	e.Add(component.NewInput())
	return e
}

// returns enemy composition.
func newEnemy() *ecs.Entity {
	const size = 60
	e := ecs.NewEntity()
	e.Add(component.NewRect(colornames.White, size, size))
	// Placing enemy at the top-mid of the screen.
	e.Add(component.NewPosition((screenWidth-size)/2, 0))
	return e
}
