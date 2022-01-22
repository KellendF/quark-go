package dashboard

import "github.com/gofiber/fiber/v2"

// 资源结构体
type Dashboard struct {
	Title string
}

// 资源接口
type DashboardInterface interface {
	Init()
	HandleInit(dashboard DashboardInterface)
	SetTitle(title string)
	GetTitle() string
	Cards() interface{}
	GetCards(dashboard DashboardInterface) interface{}
	Render(c *fiber.Ctx, dashboard DashboardInterface, content interface{}) interface{}
	DashboardComponentRender(c *fiber.Ctx, dashboard DashboardInterface) interface{}
}

// 初始化资源
func (p *Dashboard) Init() {}

// 执行资源初始化
func (p *Dashboard) HandleInit(dashboard DashboardInterface) {
	dashboard.Init()
}

// 设置标题
func (p *Dashboard) SetTitle(title string) {
	p.Title = title
}

// 获取标题
func (p *Dashboard) GetTitle() string {
	return p.Title
}

// 卡片列表
func (p *Dashboard) Cards() interface{} {
	return nil
}

// 获取卡片列表
func (p *Dashboard) GetCards(dashboard DashboardInterface) interface{} {
	return dashboard.Cards()
}
