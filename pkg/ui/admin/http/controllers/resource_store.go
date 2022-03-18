package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/pkg/framework/msg"
	"github.com/quarkcms/quark-go/pkg/ui/admin/http/requests"
)

type ResourceStore struct{}

// 执行行为
func (p *ResourceStore) Handle(c *fiber.Ctx) error {
	result := (&requests.ResourceStore{}).HandleStore(c)

	if result == nil {
		return msg.Success("操作成功！", "", "")
	} else {
		return msg.Error("操作失败！", "")
	}
}
