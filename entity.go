package ecs

// Entity represents a composition of components.
type Entity struct {
	Name string
	cm   *ComponentManager
}

// NewEntity entity constructor.
func NewEntity(name string) *Entity {
	return &Entity{
		Name: name,
		cm:   &ComponentManager{},
	}
}

// AddComponents add new component to entity.
func (e *Entity) AddComponents(cc ...Component) {
	for _, c := range cc {
		e.cm.Add(c)
	}
}

// ComponentManager return entity's ComponentManager.
func (e *Entity) ComponentManager() *ComponentManager {
	return e.cm
}

// EntityManager manages entities for the system.
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

// Get return single entity base on name.
func (em *EntityManager) Get(name string) *Entity {
	for _, e := range em.All() {
		if e.Name == name {
			return e
		}
	}

	return nil
}

// Filter return array of entity filter by name.
func (em *EntityManager) Filter(names ...string) []*Entity {
	hit := []*Entity{}
	for _, e := range em.All() {
		for _, n := range names {
			if e.Name == n {
				hit = append(hit, e)
			}
		}
	}

	return hit
}
