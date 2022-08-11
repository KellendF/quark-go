package msg

import (
	"github.com/gofiber/fiber/v2"
)

const DEFAULT_MSG string = ""
const DEFAULT_URL string = ""
const DEFAULT_DATA string = ""

// 返回错误信息
func Error(c *fiber.Ctx, msg string, url string) error {
	return c.JSON(fiber.Map{
		"component": "message",
		"status":    "error",
		"msg":       msg,
		"url":       url,
	})
}

// 返回正确信息
func Success(c *fiber.Ctx, msg string, url string, data interface{}) error {
	return c.JSON(fiber.Map{
		"component": "message",
		"status":    "success",
		"msg":       msg,
		"url":       url,
		"data":      data,
	})
}
