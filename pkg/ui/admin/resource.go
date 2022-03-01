package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/pkg/ui/component/table"
	"gorm.io/gorm"
)

// 结构体
type Resource struct {
	Layout
	Title    string
	SubTitle string
	PerPage  int
	Model    *gorm.DB
}

// 获取模型
func (p *Resource) NewModel(c *fiber.Ctx, resourceInstance interface{}) *gorm.DB {
	return resourceInstance.(*Resource).Model
}

// 列表页组件渲染
func (p *Resource) IndexComponentRender(c *fiber.Ctx, resourceInstance interface{}) interface{} {

	table := &table.Component{}
	component := table.Init().SetTitle("test")

	// indexFields := resourceInstance.(interface {
	// 	IndexFields(c *fiber.Ctx, resourceInstance interface{}) interface{}
	// }).IndexFields(c, resourceInstance)

	return component
}
