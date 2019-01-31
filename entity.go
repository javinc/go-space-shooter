package ecs

// Entity represents a composition of components.
type Entity struct {
	Name string
	cm   *ComponentManager
}

// NewEntity entity constructor.
func NewEntity() *Entity {
	return &Entity{
		cm: &ComponentManager{},
	}
}

// AddComponent add new component to entity.
func (e *Entity) AddComponent(c Component) {
	e.cm.Add(c)
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
func (em *EntityManager) Filter(name string) []*Entity {
	hit := []*Entity{}
	for _, e := range em.All() {
		if e.Name == name {
			hit = append(hit, e)
		}
	}

	return hit
}
