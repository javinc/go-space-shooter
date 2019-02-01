package system

import (
	"time"

	"github.com/javinc/ecs"
	"github.com/javinc/ecs/component"
	"github.com/veandco/go-sdl2/sdl"
)

// Control system.
type Control struct {
	w, h     int32
	lastShot time.Time
}

// NewControl Control system constructor.
func NewControl(w, h int32) *Control {
	return &Control{
		w: w,
		h: h,
	}
}

// Process Control system implements System interface.
func (s *Control) Process(em *ecs.EntityManager) {
	kk := sdl.GetKeyboardState()

	// Only for player and bullet entities
	player := em.Get("player").ComponentManager()
	s.movement(kk, player)

	bullet := s.bulletFromPool(em)
	if bullet == nil {
		return
	}

	// Hit shoot.
	if kk[sdl.SCANCODE_SPACE] == 1 {
		// Shoot cooldown.
		if time.Since(s.lastShot) <= (time.Second / 10) {
			return
		}

		playerRect := player.Get("rect").(*component.Rect)
		playerPos := player.Get("position").(*component.Position)

		bulletRect := bullet.Get("rect").(*component.Rect)
		bulletPos := bullet.Get("position").(*component.Position)

		bulletRect.Active = true
		bulletPos.X = playerPos.X + float64(bulletRect.W/2)
		bulletPos.Y = playerPos.Y - float64(playerRect.H/2)

		s.lastShot = time.Now()
	}

	// set bullet to active
	// set bullet position
}
func (s *Control) movement(kk []uint8, cm *ecs.ComponentManager) {
	if !cm.Requires("rect", "position", "velocity", "input") {
		return
	}

	rect := cm.Get("rect").(*component.Rect)
	pos := cm.Get("position").(*component.Position)
	vel := cm.Get("velocity").(*component.Velocity)

	// Supports dynamic mapping.
	if kk[sdl.SCANCODE_LEFT] == 1 && pos.X > 0 {
		pos.X -= vel.Speed
	} else if kk[sdl.SCANCODE_RIGHT] == 1 && pos.X < float64(s.w-rect.W) {
		pos.X += vel.Speed
	}
}

func (s *Control) bulletFromPool(em *ecs.EntityManager) *ecs.ComponentManager {
	pool := em.Filter("bullet")
	for _, e := range pool {
		bullet := e.ComponentManager()
		rect := bullet.Get("rect").(*component.Rect)
		if !rect.Active {
			return bullet
		}
	}

	return nil
}
