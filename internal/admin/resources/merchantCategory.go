package resources

import (
	"github.com/quarkcms/quark-go/pkg/ui/admin"
	"github.com/quarkcms/quark-go/pkg/ui/component/action"
	"github.com/quarkcms/quark-go/pkg/ui/component/field"
)

type MerchantCategory struct {
	admin.Resource
}

// 初始化
func (p *MerchantCategory) Init() interface{} {
	p.Title = "测试的标题"

	return p
}

// 字段
func (p *MerchantCategory) Fields() interface{} {
	field := &field.Component{}

	return []interface{}{
		field.SetTitle("1").JsonSerialize(),
		field.SetTitle("2").JsonSerialize(),
	}
}

// 资源行为
func (p *MerchantCategory) Actions() interface{} {
	action := &action.Component{}

	return []interface{}{
		action.SetTitle("1").JsonSerialize(),
		action.SetTitle("2").JsonSerialize(),
	}
}
