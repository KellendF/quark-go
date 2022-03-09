package admin

import "github.com/quarkcms/quark-go/pkg/ui/component/form/fields"

type Field struct{}

func (p *Field) Text(params ...string) *fields.Text {
	field := &fields.Text{}

	if len(params) == 2 {
		field.Init().SetName(params[0]).SetLabel(params[1])
	} else {
		field.Init().SetName(params[0]).SetLabel(params[0])
	}

	return field
}

func (p *Field) Radio(params ...string) *fields.Radio {
	field := &fields.Radio{}

	if len(params) == 2 {
		field.Init().SetName(params[0]).SetLabel(params[1])
	} else {
		field.Init().SetName(params[0]).SetLabel(params[0])
	}

	return field
}
