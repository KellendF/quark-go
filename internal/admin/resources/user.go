package resources

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/pkg/ui/admin"
	"github.com/quarkcms/quark-go/pkg/ui/component/action"
	"github.com/quarkcms/quark-go/pkg/ui/component/field"
)

type User struct {
	admin.Resource
}

// 初始化
func (p *User) Init() interface{} {
	p.Title = "测试的标题"

	return p
}

// 字段
func (p *User) Fields(c *fiber.Ctx) interface{} {
	field := &field.Component{}

	return []interface{}{
		field.SetTitle("1").JsonSerialize(),
		field.SetTitle("2").JsonSerialize(),
	}
}

// 资源行为
func (p *User) Actions(c *fiber.Ctx) interface{} {
	action := &action.Component{}

	return []interface{}{
		action.SetTitle("1").JsonSerialize(),
		action.SetTitle("2").JsonSerialize(),
	}
}
