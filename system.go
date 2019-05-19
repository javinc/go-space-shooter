package ecs

import (
	"github.com/veandco/go-sdl2/sdl"
)

// System handles entities component logic.
type System interface {
	Process(*EntityManager)
}

// SystemManager manages systems for the game.
type SystemManager struct {
	eventCh chan sdl.Event
	ss      []System
}

// Add appends new system.
func (sm *SystemManager) Add(s System) {
	sm.ss = append(sm.ss, s)
}

// ProcessAll executes all systems on entities.
func (sm *SystemManager) ProcessAll(em *EntityManager) {
	for _, s := range sm.ss {
		s.Process(em)
	}
}
