package ecs

// System handles entities component logic.
type System interface {
	Process([]*Entity)
}

// SystemManager manages systems.
type SystemManager struct {
	ss []System
}

// Add appends new system.
func (sm *SystemManager) Add(s System) {
	sm.ss = append(sm.ss, s)
}

// Process executes all systems on entities.
func (sm *SystemManager) Process(ee []*Entity) {
	for _, s := range sm.ss {
		s.Process(ee)
	}
}
