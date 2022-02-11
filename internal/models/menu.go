package models

import (
	"github.com/quarkcms/quark-go/pkg/framework/db"
)

// 字段
type Menu struct {
	db.Model
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
func (model *Menu) List() *Menu {
	result := &Menu{}
	model.DB().Where("status = ?", 1).Find(&result)

	return result
}

// 获取管理员权限菜单
func (model *Menu) PermissionList(adminId float64) *Menu {
	result := &Menu{}

	if adminId == 1 {
		model.DB().Where("status = ?", 1).Find(&result)
	} else {
		model.DB().Where("status = ?", 1).Find(&result)
	}

	return result
}
