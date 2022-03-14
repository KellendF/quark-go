package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type ResourceCreate struct{}

// 执行行为
func (p *ResourceCreate) Handle(c *fiber.Ctx) error {
	return c.JSON("")
}
