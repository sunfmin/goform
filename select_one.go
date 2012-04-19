package goform

import (
	"bytes"
	. "github.com/paulbellamy/mango"
	"strings"
)

type SelectOne struct {
	Name            string
	CollectionMaker CollectionMaker
	context         *OptionsRenderContext
}

func (so *SelectOne) Collection(cm CollectionMaker) (r *SelectOne) {
	r = so
	r.CollectionMaker = cm
	return
}

func (b *Builder) SelectOne(name string) (r *SelectOne) {
	r = &SelectOne{Name: name, context: &OptionsRenderContext{}}
	b.Add(name, r)
	return
}

func (f *SelectOne) Html(fo FormObject, env Env) string {
	b := bytes.NewBuffer([]byte{})
	f.context.FormObject = fo
	f.context.Name = fo.FormName() + "[" + f.Name + "]"
	opts := f.CollectionMaker(fo, env)
	for i := 0; i < opts.Len(); i++ {
		f.context.Options = append(f.context.Options, &optionvalue{Label: opts.Label(i), Value: opts.Value(i)})
	}

	if f.context.Label == "" {
		f.context.Label = strings.Title(f.Name)
	}

	err := Template.ExecuteTemplate(b, "SelectOne", f.context)
	if err != nil {
		panic(err)
	}
	return b.String()
}

func (so *SelectOne) Label(label string) (rso *SelectOne) {
	so.context.Label = label
	rso = so
	return
}
