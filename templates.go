package goform

import (
	"text/template"
)

const defaultTemplates = `
{{define "CollectionEdit"}}

{{end}}

{{define "TextField"}}
<div class="control-group">
	<label class="control-label">{{.Label}}</label>
	<div class="controls">
	<input type="text" class="{{.Class}}" placeholder="{{.Placeholder}}" name="{{.Name}}" value="{{.Value}}" >
	</div>
</div>
{{end}}

{{define "Select"}}
<div class="control-group">
	<label class="control-label">{{.Label}}</label>
	<div class="controls">
	<select name="{{.Name}}">
		{{range .Options}}
		<option value="{{.Value}}" {{.Checked}}>{{.Label}}</option>{{end}}
	</select>
	</div>
</div>
{{end}}


{{define "RichTextEditor"}}

{{end}}

{{define "RadioButtons"}}
<div class="control-group">
	<label class="control-label">{{.Label}}</label>
	<div class="controls">
		{{ $name := .Name}}
		{{range .Options}}
		<label class="radio inline">
			<input type="radio" name="{{$name}}" value="{{.Value}}" {{.Checked}}>{{.Label}}
		</label>
		{{end}}
	</div>
</div>
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
