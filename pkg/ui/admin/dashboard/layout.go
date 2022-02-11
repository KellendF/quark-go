package dashboard

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/internal/models"
	"github.com/quarkcms/quark-go/pkg/ui/component/layout"
	"github.com/quarkcms/quark-go/pkg/ui/component/page"
	"github.com/quarkcms/quark-go/pkg/ui/component/pagecontainer"
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

	menu := models.Menu{}

	getMenus := menu.List()

	fmt.Println(getMenus)

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
func (p *Dashboard) DashboardComponentRender(c *fiber.Ctx, dashboard DashboardInterface) interface{} {

	return p.GetCards(dashboard)
}

// 渲染列表页组件
func (p *Dashboard) Render(c *fiber.Ctx, dashboard DashboardInterface, content interface{}) interface{} {

	// 初始化资源
	p.HandleInit(dashboard)

	return PageComponentRender(c, dashboard, content)
}
