package models

import (
	"github.com/quarkcms/quark-go/pkg/framework/db"
)

// 字段
type Menu struct {
	db.Model
	Key        string  `json:"key"`
	Id         float64 `json:"id"`
	Name       string  `json:"name"`
	GuardName  string  `json:"guard_name"`
	Icon       string  `json:"icon"`
	Type       string  `json:"type"`
	Pid        float64 `json:"pid"`
	Sort       float64 `json:"sort"`
	Path       string  `json:"path"`
	Show       int     `json:"show"`
	Status     bool    `json:"status"`
	Locale     string  `json:"locale"`
	HideInMenu bool    `json:"hideInMenu"`
}

// 获取菜单
func (model *Menu) List() *Menu {
	result := &Menu{}
	model.DB().Where("status = ?", 1).Find(&result)

	return result
}
