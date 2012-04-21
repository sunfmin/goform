package goform

import (
	"fmt"
	. "github.com/paulbellamy/mango"
)

type Options interface {
	Label(i int) string
	Value(i int) string
	Len() int
}

type IntOptions []int

func (cs IntOptions) Label(i int) string {
	return fmt.Sprintf("%d", cs[i])
}

func (cs IntOptions) Value(i int) string {
	return fmt.Sprintf("%d", cs[i])
}

func (cs IntOptions) Len() int {
	return len(cs)
}

type StringOptions [][]string

func (cs StringOptions) Label(i int) string {
	return cs[i][1]
}

func (cs StringOptions) Value(i int) string {
	return cs[i][0]
}

func (cs StringOptions) Len() int {
	return len(cs)
}

type optionvalue struct {
	Label   string
	Value   string
	Checked string
}

type OptionsRenderContext struct {
	RenderContext
	Options []*optionvalue
}

type CollectionMaker func(fo FormObject, env Env) Options
