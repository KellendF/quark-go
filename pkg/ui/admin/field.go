package admin

import "github.com/quarkcms/quark-go/pkg/ui/component/form/fields"

type Field struct{}

func (p *Field) Text(params ...string) interface{} {
	text := &fields.Text{}

	if len(params) == 2 {
		return text.Init().SetName(params[0]).SetLabel(params[1])
	} else {
		return text.Init().SetName(params[0]).SetLabel(params[0])
	}
}
