package main

import (
	"github.com/javinc/ecs"
	"github.com/javinc/ecs/component"
	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/image/colornames"
)

// returns player composition.
func newPlayer() *ecs.Entity {
	input := component.NewInput()
	input.Map[sdl.SCANCODE_LEFT] = component.InputMoveLeft
	input.Map[sdl.SCANCODE_RIGHT] = component.InputMoveRight

	const size = 20
	e := ecs.NewEntity("player")
	e.AddComponents(
		input,
		component.NewRect(colornames.Red, size, size, true),
		// Place player at the bottom-mid of the screen.
		component.NewPosition((screenWidth-size)/2, screenHeight-size),
		component.NewVelocity(0.5),
	)
	return e
}

// returns enemy composition.
func newEnemy() *ecs.Entity {
	const size = 60
	e := ecs.NewEntity("enemy")
	e.AddComponents(
		component.NewRect(colornames.White, size, size, true),
		// Placing enemy at the top-mid of the screen.
		component.NewPosition((screenWidth-size)/2, 0),
	)
	return e
}

// returns bullet composition.
func newBullet() *ecs.Entity {
	input := component.NewInput()
	input.Map[sdl.SCANCODE_SPACE] = component.InputShootBullet

	const size = 10
	e := ecs.NewEntity("bullet")
	e.AddComponents(
		input,
		component.NewRect(colornames.Orange, size, size, false),
		component.NewPosition((screenWidth-size)/2, screenHeight-size),
		component.NewVelocity(1),
		component.NewProjectile(),
	)
	return e
}
