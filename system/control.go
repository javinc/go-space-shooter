package system

import (
	"github.com/javinc/ecs"
	"github.com/javinc/ecs/component"
	"github.com/veandco/go-sdl2/sdl"
)

// Control system.
type Control struct {
	w, h int32
}

// NewControl Control system constructor.
func NewControl(w, h int32) *Control {
	return &Control{w, h}
}

// Process Control system implements System interface.
func (s *Control) Process(em *ecs.EntityManager) {
	kk := sdl.GetKeyboardState()

	// Only for player and bullet entities
	player := em.Get("player").ComponentManager()
	s.movement(kk, player)

	for _, e := range em.Filter("bullet") {
		bullet := e.ComponentManager()
		s.shoot(kk, player, bullet)
	}
}

func (s *Control) movement(kk []uint8, e *ecs.ComponentManager) {
	if !e.Requires("input", "velocity", "rect", "position") {
		return
	}

	rect := e.Get("rect").(*component.Rect)
	pos := e.Get("position").(*component.Position)
	vel := e.Get("velocity").(*component.Velocity)

	if kk[sdl.SCANCODE_LEFT] == 1 && pos.X > 0 {
		pos.X -= vel.Speed
	} else if kk[sdl.SCANCODE_RIGHT] == 1 && pos.X < float64(s.w-rect.W) {
		pos.X += vel.Speed
	}
}

func (s *Control) shoot(kk []uint8, player, bullet *ecs.ComponentManager) {
	if !bullet.Requires("projectile", "velocity", "rect", "position") {
		return
	}

	bulletRect := bullet.Get("rect").(*component.Rect)
	bulletPos := bullet.Get("position").(*component.Position)
	bulletVel := bullet.Get("velocity").(*component.Velocity)

	// Projecting to top.
	if bulletRect.Active {
		bulletPos.Y -= bulletVel.Speed
	}

	// Out of bounds.
	if bulletPos.X < 0 || bulletPos.Y < 0 {
		bulletRect.Active = false
	}

	// Hit shoot.
	if kk[sdl.SCANCODE_SPACE] == 1 {
		playerRect := player.Get("rect").(*component.Rect)
		playerPos := player.Get("position").(*component.Position)

		bulletRect.Active = true
		bulletPos.X = playerPos.X + float64(bulletRect.W/2)
		bulletPos.Y = playerPos.Y - float64(playerRect.H/2)
	}
}
