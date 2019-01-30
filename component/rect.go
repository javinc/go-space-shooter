package component

import "image/color"

// Rect component.
type Rect struct {
	Color color.RGBA
	W, H  int32
}

// NewRect rect constructor.
func NewRect(c color.RGBA, w, h int32) *Rect {
	return &Rect{c, w, h}
}

// Name component implementation.
func (c *Rect) Name() string {
	return "rect"
}
