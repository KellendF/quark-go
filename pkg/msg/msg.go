package msg

import (
	"github.com/gofiber/fiber/v2"
)

var ctx *fiber.Ctx

const DEFAULT_MSG string = ""
const DEFAULT_URL string = ""
const DEFAULT_DATA string = ""

// 初始化
func Init(c *fiber.Ctx) {
	ctx = c
}

// 生成随机字符串
func Error(msg string, url string) error {
	return ctx.JSON(fiber.Map{
		"status": "error",
		"msg":    msg,
		"url":    url,
	})
}

// 生成字母类型字符串
func Success(msg string, url string, data interface{}) error {
	return ctx.JSON(fiber.Map{
		"status": "success",
		"msg":    msg,
		"url":    url,
		"data":   data,
	})
}
