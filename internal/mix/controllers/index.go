package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/pkg/framework/msg"
)

type Index struct{}

// Index
func (p *Index) Index(c *fiber.Ctx) error {

	return msg.Success(c, "操作成功！", "测试", "")
}
