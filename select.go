package goform

import (
	"bytes"
	"fmt"
	. "github.com/paulbellamy/mango"
	"strings"
)

type Select struct {
	Name            string
	CollectionMaker CollectionMaker
	context         *OptionsRenderContext
}

func (so *Select) Collection(cm CollectionMaker) (r *Select) {
	r = so
	r.CollectionMaker = cm
	return
}

func (b *Builder) Select(name string) (r *Select) {
	r = &Select{Name: name, context: &OptionsRenderContext{}}
	b.Add(name, r)
	return
}

func (f *Select) Html(fo FormObject, env Env) string {
	b := bytes.NewBuffer([]byte{})
	f.context.FormObject = fo
	f.context.Name = fo.FormName() + "[" + f.Name + "]"
	f.context.Value = ValueOf(fo, f.Name)
	opts := f.CollectionMaker(fo, env)
	for i := 0; i < opts.Len(); i++ {
		ov := &optionvalue{Label: opts.Label(i), Value: opts.Value(i)}
		if ov.Value == fmt.Sprintf("%v", f.context.Value) {
			ov.Checked = "selected"
		}
		f.context.Options = append(f.context.Options, ov)
	}

	if f.context.Label == "" {
		f.context.Label = strings.Title(f.Name)
	}

	err := Template.ExecuteTemplate(b, "Select", f.context)
	if err != nil {
		panic(err)
	}
	return b.String()
}

func (f *Select) Label(label string) (rf *Select) {
	f.context.Label = label
	rf = f
	return
}
