package ecs

import "fmt"

// ComponentName identifier.
type ComponentName string

// Component holds data of one behavior.
type Component interface {
	Name() string
}

// ComponentManager manages components.
type ComponentManager struct {
	// Component name regsitry.
	names []string
	cc    []Component
}

// Add appends new component.
func (cm *ComponentManager) Add(c Component) {
	n := c.Name()
	if cm.hasName(n) {
		panic(fmt.Sprintf("duplicate component [%s]", n))
	}

	cm.names = append(cm.names, n)
	cm.cc = append(cm.cc, c)
}

// Requires checks component existence.
func (cm *ComponentManager) Requires(names ...string) bool {
	hit := 0
	for _, n := range names {
		if cm.hasName(n) {
			hit++
			continue
		}
	}

	return hit == len(names)
}

// Get returns component base on name.
func (cm *ComponentManager) Get(name string) Component {
	for _, c := range cm.cc {
		if c.Name() == name {
			return c
		}
	}

	return nil
}

func (cm *ComponentManager) hasName(name string) bool {
	for _, n := range cm.names {
		if n == name {
			return true
		}
	}

	return false
}
