package actions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/pkg/framework/msg"
	"github.com/quarkcms/quark-go/pkg/ui/admin/actions"
	"gorm.io/gorm"
)

type SyncPermission struct {
	actions.Action
}

// 初始化
func (p *SyncPermission) Init() *SyncPermission {
	// 初始化父结构
	p.ParentInit()

	// 行为名称
	p.Name = "同步权限"

	// 执行成功后刷新的组件
	p.Reload = "table"

	// 是否具有loading，当action 的作用类型为ajax,submit时有效
	p.WithLoading = true

	// 设置展示位置
	p.SetOnlyOnIndex(true)

	// 行为类型
	p.ActionType = "ajax"

	return p
}

// 执行行为句柄
func (p *SyncPermission) Handle(c *fiber.Ctx, model *gorm.DB) error {
	return msg.Error("操作失败，请重试！", "")
}
