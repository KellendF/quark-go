package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/pkg/component/layout"
	"github.com/quarkcms/quark-go/pkg/component/page"
	"github.com/quarkcms/quark-go/pkg/component/pagecontainer"
)

// 渲染页面组件
func PageComponentRender(c *fiber.Ctx, resource ResourceInterface, content interface{}) interface{} {
	component := &page.Component{}

	layoutComponent := LayoutComponentRender(c, resource, content)

	return component.
		SetTitle(resource.GetTitle()).
		SetBody(layoutComponent).
		JsonSerialize()
}

// 渲染页面布局组件
func LayoutComponentRender(c *fiber.Ctx, resource ResourceInterface, content interface{}) interface{} {

	component := &layout.Component{}
	return component.SetTitle(resource.GetTitle()).
		SetBody(PageContainerComponentRender(content)).
		JsonSerialize()
}

// 渲染页面容器组件
func PageContainerComponentRender(content interface{}) interface{} {
	component := &pagecontainer.Component{}
	return component.SetBody(content).JsonSerialize()
}

// 渲染列表页组件
func (p *Resource) IndexComponentRender() interface{} {
	return "xxx"
}

// 渲染列表页组件
func (p *Resource) Render(c *fiber.Ctx, resource ResourceInterface, content interface{}) interface{} {

	// 初始化资源
	resource.HandleInit(resource)

	return PageComponentRender(c, resource, content)
}
