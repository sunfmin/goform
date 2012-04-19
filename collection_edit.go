package goform

import (
	. "github.com/paulbellamy/mango"
)

type CollectionEdit struct {
	Name string
	B    *Builder
}

func (b *Builder) CollectionEdit(name string) (r *Builder) {
	ib := &Builder{}
	b.Add(name, &CollectionEdit{Name: name, B: ib})
	return ib
}

func (ce *CollectionEdit) Html(obj FormObject, env Env) string {
	return "hello"
}
