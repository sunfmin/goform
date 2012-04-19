package goform

import (
	"fmt"
	. "github.com/paulbellamy/mango"
)

type Shoes struct {
	Name string
	Size int
}

type Ambassador struct {
	Type       string
	CategoryId string
	Profile    string
}

func (am *Ambassador) Shoes() []*Shoes {
	return []*Shoes{
		{
			Name: "Nimbus 18",
			Size: 12,
		},
	}
}

func (am *Ambassador) FormName() string {
	return "Ambassador"
}

type Category struct {
	Id   string `bson:"_id"`
	Name string
}

type CategoryOptions []*Category

func (cs CategoryOptions) Label(i int) string {
	return cs[i].Name
}

func (cs CategoryOptions) Value(i int) string {
	return cs[i].Id
}

func (cs CategoryOptions) Len() int {
	return len(cs)
}

func ExampleFormBuilder() {
	am := &Ambassador{
		Type:       "Running",
		CategoryId: "2",
		Profile:    "Hello, I am a runner",
	}

	f := NewFormBuilder()
	f.TextField("Type")

	f.Select("CategoryId").Collection(func(fo FormObject, env Env) Options {
		var cats []*Category
		cats = append(cats, &Category{"1", "Cat1"})
		cats = append(cats, &Category{"2", "Cat2"})
		return CategoryOptions(cats)
	})
	f.RichTextEditor("Profile").Label("The Ambassador Profile")
	shc := f.CollectionEdit("Shoes")
	shc.TextField("Name").Label("Your Name: ")
	shc.Select("Size").Collection(func(obj FormObject, env Env) Options {
		return IntOptions([]int{12, 14, 16, 18})
	})
	env := Env(make(map[string]interface{}))
	html := f.Render(am, env)

	fmt.Println(html)
	//Output: hello
}
