package resources

import (
	"github.com/quarkcms/quark-go/pkg/ui/component/action"
	"github.com/quarkcms/quark-go/pkg/ui/component/field"
	"github.com/quarkcms/quark-go/pkg/ui/resource"
)

type Admin struct {
	resource.Resource
}

// 初始化
func (p *Admin) Init() {

	p.Title = "测试的标题"
}

// 字段
func (p *Admin) Fields() interface{} {

	field := &field.Component{}

	return []interface{}{
		field.SetTitle("1").JsonSerialize(),
		field.SetTitle("2").JsonSerialize(),
	}
}

// 资源行为
func (p *Admin) Actions() interface{} {

	action := &action.Component{}

	return []interface{}{
		action.SetTitle("1").JsonSerialize(),
		action.SetTitle("2").JsonSerialize(),
	}
}
