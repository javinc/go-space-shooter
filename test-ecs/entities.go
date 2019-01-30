package main

import (
	"reflect"

	"golang.org/x/image/colornames"
)

const (
	tagPlayer = "player"
	tagEnemy  = "enemy"
)

type entity struct {
	tag        string
	components []component
}

func (e *entity) addComponent(c component) {
	// Check for duplicate component and replace it.

	e.components = append(e.components, c)
}

func (e *entity) getComponent(c component) component {
	for _, ec := range e.components {
		if reflect.TypeOf(c) == reflect.TypeOf(ec) {
			return ec
		}
	}

	return nil
}

func (e *entity) importEntity(new entity) {
	// Merge new entity components.
}

func newPlayer() *entity {
	e := new(entity)
	e.tag = tagPlayer
	e.addComponent(newRender(colornames.Red, 30, 30))
	e.addComponent(newPosition(100, 100))
	e.addComponent(newInput())
	return e
}

func newEnemy() *entity {
	e := new(entity)
	e.tag = tagEnemy
	e.addComponent(newRender(colornames.White, 60, 60))
	e.addComponent(newPosition(200, 200))
	return e
}

func newBackground(w, h int32) *entity {
	e := new(entity)
	e.addComponent(newRender(colornames.Black, w, h))
	return e
}
