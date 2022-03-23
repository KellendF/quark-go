package models

import (
	"github.com/quarkcms/quark-go/pkg/framework/db"
	"github.com/quarkcms/quark-go/pkg/ui/admin/utils"
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

// 获取菜单的有序列表
func (model *Menu) OrderedList() map[int]string {
	var menus []interface{}
	model.DB().Where("guard_name = ?", "admin").Order("sort asc,id asc").Find(&menus)
	lists := map[int]string{}

	menuTrees := utils.ListToTree(menus, "id", "pid", "children", 0)
	menuTreeLists := utils.TreeToOrderedList(menuTrees, 0, "name", "children")

	lists[0] = "根节点"
	for _, v := range menuTreeLists {
		getID := v.(map[string]interface{})["id"].(float64)
		lists[int(getID)] = v.((map[string]interface{}))["name"].(string)
	}

	return lists
}

// 获取菜单的tree
func (model *Menu) Tree() []interface{} {
	menus := []Menu{}
	model.DB().Where("status = ?", 1).Select("name", "id", "pid").Find(&menus)
	lists := []interface{}{}

	for _, v := range menus {
		item := map[string]interface{}{
			"key":   v.Id,
			"pid":   v.Pid,
			"title": v.Name,
		}
		lists = append(lists, item)
	}

	return utils.ListToTree(lists, "key", "pid", "children", 0)
}
