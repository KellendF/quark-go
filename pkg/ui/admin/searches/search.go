package searches

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Search struct {
	Column    string `json:"column"`
	Name      string `json:"name"`
	Component string `json:"component"`
	Operator  string `json:"operator"`
	Api       string `json:"api"`
}

// 初始化
func (p *Search) ParentInit() interface{} {
	p.Component = "input"

	return p
}

/**
 * 获取字段名
 *
 * @return string
 */
func (p *Search) GetColumn() string {
	return p.Column
}

/**
 * 获取名称
 *
 * @return string
 */
func (p *Search) GetName() string {
	return p.Name
}

/**
 * 获取组件名称
 *
 * @return string
 */
func (p *Search) GetComponent() string {
	return p.Component
}

/**
 * 获取接口
 *
 * @return string
 */
func (p *Search) GetApi() string {
	return p.Api
}

/**
 * 获取操作符
 *
 * @return string
 */
func (p *Search) GetOperator() string {
	return p.Operator
}

/**
 * 默认值
 *
 * @var string
 */
func (p *Search) GetDefault() interface{} {
	return true
}

/**
 * 执行查询
 *
 * @param  Request  request
 * @param  Builder  query
 * @param  mixed  value
 * @return Builder
 */
func (p *Search) Apply(c *fiber.Ctx, query *gorm.DB, value interface{}) *gorm.DB {
	return query
}

/**
 * 属性
 *
 * @param  Request  request
 * @return array
 */
func (p *Search) Options(c *fiber.Ctx) interface{} {
	return nil
}
