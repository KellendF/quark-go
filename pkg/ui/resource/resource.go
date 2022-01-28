package resource

import "github.com/gofiber/fiber/v2"

// 资源结构体
type Resource struct {
	Title string
}

// 资源接口
type ResourceInterface interface {
	Init()
	HandleInit(resource ResourceInterface)
	SetTitle(title string)
	GetTitle() string
	Fields() interface{}
	Actions() interface{}
	Render(c *fiber.Ctx, resource ResourceInterface, content interface{}) interface{}
	IndexComponentRender() interface{}
}

// 初始化资源
func (p *Resource) Init() {}

// 执行资源初始化
func (p *Resource) HandleInit(resource ResourceInterface) {
	resource.Init()
}

// 设置标题
func (p *Resource) SetTitle(title string) {
	p.Title = title
}

// 获取标题
func (p *Resource) GetTitle() string {
	return p.Title
}
