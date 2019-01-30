package main

import (
	"github.com/javinc/ecs"
	"github.com/javinc/ecs/component"
	"github.com/javinc/ecs/system"
	"golang.org/x/image/colornames"
)

const (
	title   = "ECS TEST"
	screenW = 400
	screenH = 600
)

func main() {
	e := ecs.New(title, screenW, screenH)
	e.Start()

	// Register entities.
	e.EntityManager.Add(newPlayer())
	e.EntityManager.Add(newEnemy())

	// Register systems.
	e.SystemManager.Add(system.NewControl(screenW, screenH))
	e.SystemManager.Add(system.NewRender(e.Renderer))

	e.Run()
	e.Stop()
}

// Compose player components.
func newPlayer() *ecs.Entity {
	e := ecs.NewEntity()
	e.Add(component.NewRect(colornames.Red, 30, 30))
	e.Add(component.NewPosition(100, 100))
	e.Add(component.NewInput())
	return e
}

// Compose enemy components.
func newEnemy() *ecs.Entity {
	e := ecs.NewEntity()
	e.Add(component.NewRect(colornames.White, 60, 60))
	e.Add(component.NewPosition(10, 10))
	return e
}
