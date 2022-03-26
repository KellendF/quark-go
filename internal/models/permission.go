package models

import (
	"github.com/quarkcms/quark-go/pkg/framework/db"
)

// 权限
type Permission struct {
	db.Model
	Id        int
	MenuId    int
	Name      string
	GuardName string
}

// 获取列表
func (model *Permission) List() map[interface{}]interface{} {
	permissions := []Permission{}
	results := map[interface{}]interface{}{}

	model.DB().Find(&permissions)
	for _, v := range permissions {
		results[v.Id] = v.Name
	}

	return results
}
