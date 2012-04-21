package goform

import (
	"bytes"
	"fmt"
	. "github.com/paulbellamy/mango"
	"strings"
)

type RadioButtons struct {
	Name            string
	CollectionMaker CollectionMaker
	context         *OptionsRenderContext
}

func (so *RadioButtons) Collection(cm CollectionMaker) (r *RadioButtons) {
	r = so
	r.CollectionMaker = cm
	return
}

func (b *Builder) RadioButtons(name string) (r *RadioButtons) {
	r = &RadioButtons{Name: name, context: &OptionsRenderContext{}}
	b.Add(name, r)
	return
}

func (f *RadioButtons) Html(fo FormObject, env Env) string {
	b := bytes.NewBuffer([]byte{})
	f.context.FormObject = fo
	f.context.Name = fo.FormName() + "[" + f.Name + "]"
	f.context.Value = ValueOf(fo, f.Name)
	opts := f.CollectionMaker(fo, env)
	for i := 0; i < opts.Len(); i++ {
		ov := &optionvalue{Label: opts.Label(i), Value: opts.Value(i)}
		if ov.Value == fmt.Sprintf("%v", f.context.Value) {
			ov.Checked = "checked"
		}
		f.context.Options = append(f.context.Options, ov)
	}

	if f.context.Label == "" {
		f.context.Label = strings.Title(f.Name)
	}

	err := Template.ExecuteTemplate(b, "RadioButtons", f.context)
	if err != nil {
		panic(err)
	}
	return b.String()
}

func (f *RadioButtons) Label(label string) (rf *RadioButtons) {
	f.context.Label = label
	rf = f
	return
}
