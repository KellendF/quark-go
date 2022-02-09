package models

import (
	"github.com/quarkcms/quark-go/pkg/framework/db"
)

// 模型
type MenuModel struct {
	db.Model
}

// 字段
type Menu struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	GuardName string `json:"guard_name"`
	Icon      string `json:"icon"`
	Type      string `json:"type"`
	Pid       string `json:"pid"`
	Sort      string `json:"sort"`
	Path      string `json:"path"`
	Show      string `json:"show"`
	Status    string `json:"status"`
}

// 获取菜单
func (model *MenuModel) GetMenus(username string) *Menu {
	menu := &Menu{}
	model.DB().Where("status = ?", 1).Find(&menu)

	return menu
}
