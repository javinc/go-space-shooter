package main

import (
	"image/color"

	"github.com/veandco/go-sdl2/sdl"
	"golang.org/x/image/colornames"
)

// Bullet entity
type Bullet struct {
	position
	color color.RGBA
	size  int32

	active bool
}

const (
	bulletSize  = 10
	bulletSpeed = 1

	bulletPoolMax = 30
)

var (
	bulletPool []*Bullet
)

func newBullet() (b Bullet) {
	b.color = colornames.Orange
	b.size = bulletSize
	return
}

func (b *Bullet) draw(r *sdl.Renderer) error {
	if !b.active {
		return nil
	}

	rect := &sdl.Rect{
		X: int32(b.x),
		Y: int32(b.y),
		W: b.size,
		H: b.size,
	}

	if err := setDrawColorByColorname(r, b.color); err != nil {
		return err
	}
	if err := r.FillRect(rect); err != nil {
		return err
	}

	return nil
}

func (b *Bullet) update() error {
	b.y -= bulletSpeed

	// Bring back the bullet from the pool when off-screen.
	if b.y < 0 || b.x < 0 {
		b.active = false
	}

	return nil
}

func initBulletPool() {
	for i := 0; i < bulletPoolMax; i++ {
		b := newBullet()
		bulletPool = append(bulletPool, &b)
	}
}

func bulletFromPool() (b *Bullet, ok bool) {
	// Search for available bullet.
	for _, b := range bulletPool {
		if b.active == false {
			return b, true
		}
	}

	return nil, false
}
