package ecs

// Entity represents a composition of components.
type Entity struct {
	tag string
	ComponentManager
}

// NewEntity entity constructor.
func NewEntity() *Entity {
	e := new(Entity)
	return e
}

// EntityManager manages entities.
type EntityManager struct {
	ee []*Entity
}

// Add appends new entity.
func (em *EntityManager) Add(e *Entity) {
	em.ee = append(em.ee, e)
}

// All returns all entities.
func (em *EntityManager) All() []*Entity {
	return em.ee
}
