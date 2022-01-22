package dashboard

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/pkg/component/layout"
	"github.com/quarkcms/quark-go/pkg/component/page"
	"github.com/quarkcms/quark-go/pkg/component/pagecontainer"
)

// 渲染页面组件
func PageComponentRender(c *fiber.Ctx, dashboard DashboardInterface, content interface{}) interface{} {
	component := &page.Component{}

	layoutComponent := LayoutComponentRender(c, dashboard, content)

	return component.
		SetTitle(dashboard.GetTitle()).
		SetBody(layoutComponent).
		JsonSerialize()
}

// 渲染页面布局组件
func LayoutComponentRender(c *fiber.Ctx, dashboard DashboardInterface, content interface{}) interface{} {

	component := &layout.Component{}
	return component.SetTitle(dashboard.GetTitle()).
		SetBody(PageContainerComponentRender(content)).
		JsonSerialize()
}

// 渲染页面容器组件
func PageContainerComponentRender(content interface{}) interface{} {
	component := &pagecontainer.Component{}
	return component.SetBody(content).JsonSerialize()
}

// 渲染列表页组件
func (p *Dashboard) DashboardComponentRender() interface{} {
	return "xxxx"
}

// 渲染列表页组件
func (p *Dashboard) Render(c *fiber.Ctx, dashboard DashboardInterface, content interface{}) interface{} {

	// 初始化资源
	dashboard.HandleInit(dashboard)

	return PageComponentRender(c, dashboard, content)
}
