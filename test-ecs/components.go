package main

import "image/color"

type component interface {
	exec() error
}

type render struct {
	color color.RGBA
	w, h  int32
}

func newRender(c color.RGBA, w, h int32) *render {
	return &render{c, w, h}
}

func (r *render) exec() error {
	return nil
}

type position struct {
	x, y float64
}

func newPosition(x, y float64) *position {
	return &position{x, y}
}

func (r *position) exec() error {
	return nil
}

type input struct {
}

func newInput() *input {
	return &input{}
}

func (i *input) exec() error {
	return nil
}
