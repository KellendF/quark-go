package dashboard

import (
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/pkg/ui/admin/utils"
	"github.com/quarkcms/quark-go/pkg/ui/component/card"
	"github.com/quarkcms/quark-go/pkg/ui/component/statistic"
)

// 资源结构体
type Dashboard struct {
	Title    string
	SubTitle string
}

// 资源接口
type DashboardInterface interface {
	Init()
	HandleInit(dashboard DashboardInterface)
	SetTitle(title string)
	GetTitle() string
	SetSubTitle(subTitle string)
	GetSubTitle() string
	Cards(c *fiber.Ctx) []any
	GetCards(c *fiber.Ctx, dashboard DashboardInterface) interface{}
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

// 设置子标题
func (p *Dashboard) SetSubTitle(subTitle string) {
	p.SubTitle = subTitle
}

// 获取子标题
func (p *Dashboard) GetSubTitle() string {
	return p.SubTitle
}

// 卡片列表
func (p *Dashboard) Cards(c *fiber.Ctx) []any {
	return nil
}

// 获取卡片列表
func (p *Dashboard) GetCards(c *fiber.Ctx, dashboard DashboardInterface) interface{} {
	cards := dashboard.Cards(c)
	var result []interface{}
	var colNum float64 = 0

	for _, v := range cards {

		item := (&card.Component{}).Init().SetBody(v.(interface{ Calculate() *statistic.Component }).Calculate())

		// struct转换map
		vMap := utils.StructToMap(v).(map[string]interface{})
		colNum = colNum + vMap["Col"].(float64)

		result = append(result, item)
	}

	return result
}
