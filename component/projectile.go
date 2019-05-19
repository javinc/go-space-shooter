package component

// Projectile component.
type Projectile struct {
	FireRate int
}

// NewProjectile projectile constructor.
func NewProjectile() *Projectile {
	return &Projectile{
		FireRate: 12,
	}
}

// Name component implementation.
func (c *Projectile) Name() string {
	return "projectile"
}
