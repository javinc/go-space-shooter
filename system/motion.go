package system

import (
	"github.com/javinc/ecs"
	"github.com/javinc/ecs/component"
)

// Motion system.
type Motion struct{}

// NewMotion Motion system constructor.
func NewMotion() *Motion {
	return &Motion{}
}

// Process Control system implements System interface.
func (s *Motion) Process(em *ecs.EntityManager) {
	for _, e := range em.All() {
		cm := e.ComponentManager()
		if !cm.Requires("rect", "position", "velocity", "projectile") {
			continue
		}

		rect := cm.Get("rect").(*component.Rect)
		if !rect.Active {
			continue
		}

		pos := cm.Get("position").(*component.Position)
		vel := cm.Get("velocity").(*component.Velocity)
		// proj := cm.Get("position").(*component.Projectile)

		// Shoot up
		pos.Y -= vel.Speed

		// Out of bounds for reuse
		if pos.Y < 0 || pos.X < 0 {
			rect.Active = false
		}
	}
}
