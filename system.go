package ecs

// System handles entities component logic.
type System interface {
	Process(*EntityManager)
}

// SystemManager manages systems for the game.
type SystemManager struct {
	ss []System
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
