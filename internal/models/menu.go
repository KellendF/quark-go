package models

import (
	"time"

	"github.com/quarkcms/quark-go/pkg/framework/db"
	"github.com/quarkcms/quark-go/pkg/ui/admin/utils"
)

// 字段
type Menu struct {
	db.Model
	Key        string
	Id         int
	Name       string
	GuardName  string
	Icon       string
	Type       string
	Pid        int
	Sort       int
	Path       string
	Show       int
	Status     int
	Locale     string
	HideInMenu int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// 获取菜单的有序列表
func (model *Menu) OrderedList() []map[string]interface{} {
	var menus []map[string]interface{}

	(&db.Model{}).
		Model(&model).
		Where("guard_name = ?", "admin").
		Order("sort asc,id asc").
		Find(&menus)

	lists := []map[string]interface{}{}
	menuTrees := utils.ListToTree(menus, "id", "pid", "children", 0)
	menuTreeLists := utils.TreeToOrderedList(menuTrees, 0, "name", "children")

	lists = append(lists, map[string]interface{}{
		"label": "根节点",
		"value": 0,
	})

	for _, v := range menuTreeLists {
		option := map[string]interface{}{
			"label": v.((map[string]interface{}))["name"],
			"value": v.(map[string]interface{})["id"],
		}

		lists = append(lists, option)
	}

	return lists
}

// 获取菜单的tree
func (model *Menu) Tree() []interface{} {
	menus := []Menu{}
	model.DB().Where("status = ?", 1).Select("name", "id", "pid").Find(&menus)
	lists := []map[string]interface{}{}

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
