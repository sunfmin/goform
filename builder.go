package goform

import (
	"bytes"
	. "github.com/paulbellamy/mango"
	"html/template"
)

type Builder struct {
	FormObject FormObject
	Fields     []*Field
}

type Htmlize interface {
	Html(fo FormObject, env Env) string
}

type FormObject interface {
	FormName() string
}

type RenderContext struct {
	Name        string
	FormObject  FormObject
	Value       interface{}
	Label       string
	Class       string
	Placeholder string
}

type Field struct {
	Name    string
	Htmlize Htmlize
}

func (b *Builder) Add(name string, htmlize Htmlize) {
	b.Fields = append(b.Fields, &Field{name, htmlize})
}

func NewFormBuilder() (b *Builder) {
	b = new(Builder)
	return
}

func (b *Builder) Render(obj FormObject, env Env) (r template.HTML) {
	buf := bytes.NewBufferString("")

	for _, f := range b.Fields {
		buf.WriteString(f.Htmlize.Html(obj, env))
	}
	r = template.HTML(buf.String())
	return
}
