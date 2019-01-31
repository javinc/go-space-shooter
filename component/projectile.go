package component

// Projectile component.
type Projectile struct {
	Active bool
}

// NewProjectile projectile constructor.
func NewProjectile() *Projectile {
	return &Projectile{}
}

// Name component implementation.
func (c *Projectile) Name() string {
	return "projectile"
}
