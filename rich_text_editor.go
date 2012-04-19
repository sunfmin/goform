package goform

import (
	. "github.com/paulbellamy/mango"
)

type RichTextEditor struct {
	Name       string
	LabelValue string
}

func (b *Builder) RichTextEditor(name string) (r *RichTextEditor) {
	r = &RichTextEditor{Name: name}
	b.Add(name, r)
	return
}

func (rte *RichTextEditor) Label(label string) (r *RichTextEditor) {
	rte.LabelValue = label
	r = rte
	return
}

func (ce *RichTextEditor) Html(obj FormObject, env Env) string {
	return "RichTextEditor"
}
