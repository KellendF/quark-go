package actions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/pkg/ui/admin/actions"
)

type CreateLink struct {
	actions.Link
}

// 初始化
func (p *CreateLink) Init() *CreateLink {
	// 初始化行为
	p.ActionInit()
	p.Name = "创建"

	return p
}

// 字段
func (p *CreateLink) GetHref(c *fiber.Ctx) string {
	return ""
}
