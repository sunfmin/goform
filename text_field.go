package goform

import (
	"bytes"
	. "github.com/paulbellamy/mango"
	"strings"
)

type TextField struct {
	Name    string
	context *RenderContext
}

func (b *Builder) TextField(name string) (r *TextField) {
	r = &TextField{Name: name, context: &RenderContext{}}
	b.Add(name, r)
	return
}

func (tf *TextField) Label(label string) (rtf *TextField) {
	tf.context.Label = label
	rtf = tf
	return
}

func (tf *TextField) Placeholder(placeholder string) (rtf *TextField) {
	tf.context.Placeholder = placeholder
	rtf = tf
	return
}

func (tf *TextField) Html(fo FormObject, env Env) string {
	b := bytes.NewBuffer([]byte{})
	tf.context.FormObject = fo
	tf.context.Name = fo.FormName() + "[" + tf.Name + "]"
	if tf.context.Label == "" {
		tf.context.Label = strings.Title(tf.Name)
	}
	tf.context.Value = ValueOf(fo, tf.Name)

	err := Template.ExecuteTemplate(b, "TextField", tf.context)
	if err != nil {
		panic(err)
	}
	return b.String()
}
