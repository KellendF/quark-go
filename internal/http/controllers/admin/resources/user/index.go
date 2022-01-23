package user

import (
	"github.com/quarkcms/quark-go/pkg/component/action"
	"github.com/quarkcms/quark-go/pkg/component/field"
	"github.com/quarkcms/quark-go/pkg/resource"
)

type Component struct {
	resource.Resource
}

// 字段
func (p *Component) Fields() interface{} {

	field := &field.Component{}

	return []interface{}{
		field.SetTitle("1").JsonSerialize(),
		field.SetTitle("2").JsonSerialize(),
	}
}

// 资源行为
func (p *Component) Actions() interface{} {

	action := &action.Component{}

	return []interface{}{
		action.SetTitle("1").JsonSerialize(),
		action.SetTitle("2").JsonSerialize(),
	}
}
