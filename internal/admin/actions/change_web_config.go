package actions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/pkg/framework/msg"
	"github.com/quarkcms/quark-go/pkg/ui/admin/actions"
	"gorm.io/gorm"
)

type ChangeWebConfig struct {
	actions.Action
}

// 执行行为句柄
func (p *ChangeWebConfig) Handle(c *fiber.Ctx, model *gorm.DB) error {

	if true {
		return msg.Success("操作成功！", "", "")
	} else {
		return msg.Error("操作失败，请重试！", "")
	}
}
