package admin

import "github.com/quarkcms/quark-go/pkg/ui/component/form/fields"

type Field struct{}

func (p *Field) Text(params ...string) *fields.Text {
	text := &fields.Text{}

	if len(params) == 2 {
		text.Init().SetName(params[0]).SetLabel(params[1])
	} else {
		text.Init().SetName(params[0]).SetLabel(params[0])
	}

	return text
}
