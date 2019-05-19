package main

import (
	"github.com/javinc/go-space-shooter"
	"github.com/javinc/go-space-shooter/system"
)

const (
	title        = "Shoot-em-Up"
	screenWidth  = 450
	screenHeight = 600
)

var engine *ecs.Engine

func init() {
	engine = ecs.New(title, screenWidth, screenHeight)
}

func main() {
	engine.Start()

	// Register entities.
	engine.AddEntity(newEnemy())
	engine.AddEntity(newPlayer())
	newEntityPool(newBullet, 1000)

	// Register systems.
	engine.AddSystems(
		system.NewControl(screenWidth, screenHeight),
		system.NewMotion(),
		system.NewRender(engine.Renderer),
	)

	engine.Run()
	engine.Stop()
}

func newEntityPool(fn func() *ecs.Entity, count int) {
	for i := 0; i < count; i++ {
		engine.AddEntity(fn())
	}
}
