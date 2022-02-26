package admin

import "github.com/gofiber/fiber/v2"

// 资源结构体
type Resource struct {
	Layout
	Title    string
	SubTitle string
}

// 列表页组件渲染
func (p *Resource) IndexComponentRender(c *fiber.Ctx, componentInterface interface{}) interface{} {
	return "xxxx"
}
