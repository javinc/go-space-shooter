package main

import (
	"github.com/javinc/ecs"
	"github.com/javinc/ecs/component"
	"github.com/javinc/ecs/system"
	"golang.org/x/image/colornames"
)

const (
	title        = "Shoot-em-Up"
	screenWidth  = 450
	screenHeight = 600
)

func main() {
	engine := ecs.New(title, screenWidth, screenHeight)
	engine.Start()

	// Register entities.
	engine.AddEntity(newEnemy())
	engine.AddEntity(newPlayer())
	engine.AddEntity(newBullet())

	// Register systems.
	engine.AddSystem(system.NewControl(screenWidth, screenHeight))
	engine.AddSystem(system.NewRender(engine.Renderer))

	engine.Run()
	engine.Stop()
}

// returns player composition.
func newPlayer() *ecs.Entity {
	const size = 20
	e := ecs.NewEntity("player")
	e.AddComponent(component.NewRect(colornames.Red, size, size, true))
	// Place player at the bottom-mid of the screen.
	e.AddComponent(component.NewPosition((screenWidth-size)/2, screenHeight-size))
	e.AddComponent(component.NewVelocity(0.5))
	e.AddComponent(component.NewInput())
	return e
}

// returns enemy composition.
func newEnemy() *ecs.Entity {
	const size = 60
	e := ecs.NewEntity("enemy")
	e.AddComponent(component.NewRect(colornames.White, size, size, true))
	// Placing enemy at the top-mid of the screen.
	e.AddComponent(component.NewPosition((screenWidth-size)/2, 0))
	return e
}

// returns bullet composition.
func newBullet() *ecs.Entity {
	const size = 10
	e := ecs.NewEntity("bullet")
	e.AddComponent(component.NewRect(colornames.Orange, size, size, false))
	e.AddComponent(component.NewPosition((screenWidth-size)/2, screenHeight-size))
	e.AddComponent(component.NewVelocity(1))
	e.AddComponent(component.NewProjectile())
	return e
}
