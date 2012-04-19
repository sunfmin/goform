package goform

import (
	"html/template"
)

const defaultTemplates = `
{{define "CollectionEdit"}}

{{end}}

{{define "TextField"}}
	{{if .Label}}
	<label class="control-label">{{.Label}}</label>
	{{end}}
	<div class="controls">
	<input type="text" class="{{.Class}}" placeholder="{{.Placeholder}}" name="{{.Name}}" value="{{.Value}}" >
	</div>
{{end}}

{{define "Select"}}
	{{if .Label}}
	<label class="control-label">{{.Label}}</label>
	{{end}}
	<div class="controls">
	<select name="{{.Name}}">
		{{range .Options}}
		<option value="{{.Value}}" checked="{{.Selected}}">{{.Label}}</option>{{end}}
	</select>
	</div>
{{end}}


{{define "RichTextEditor"}}

{{end}}
`

var Template *template.Template

func init() {
	var err error
	Template, err = template.New("goform").Parse(defaultTemplates)
	if err != nil {
		panic(err)
	}
}
