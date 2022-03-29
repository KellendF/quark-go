package models

import (
	"time"

	"github.com/quarkcms/quark-go/pkg/framework/db"
)

// 角色
type Role struct {
	db.Model
	Id        int
	Name      string
	GuardName string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// 模型角色关联表
type ModelHasRole struct {
	db.Model
	RoleId    int
	ModelType string
	ModelId   int
}

// 角色权限关联表
type RoleHasPermission struct {
	db.Model
	PermissionId int
	RoleId       string
}

// 获取角色列表
func (model *Role) List() map[interface{}]interface{} {
	roles := []Role{}
	results := map[interface{}]interface{}{}

	model.DB().Find(&roles)
	for _, v := range roles {
		results[v.Id] = v.Name
	}

	return results
}
