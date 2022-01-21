package administrator

import (
	"github.com/quarkcms/quark-go/pkg/admin"
	"github.com/quarkcms/quark-go/pkg/component/action"
	"github.com/quarkcms/quark-go/pkg/component/field"
)

type Component struct {
	admin.Resource
}

// 初始化
func (p *Component) Init() {

	p.Title = "测试的标题"
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
