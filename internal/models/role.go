package models

import (
	"github.com/quarkcms/quark-go/pkg/framework/db"
)

// 角色
type Role struct {
	db.Model
	Id        int    `json:"id"`
	Name      string `json:"name"`
	GuardName string `json:"guard_name"`
}

// 模型角色关联表
type ModelHasRole struct {
	db.Model
	RoleId    int    `json:"role_id"`
	ModelType string `json:"model_type"`
	ModelId   int    `json:"model_id"`
}

// 角色权限关联表
type RoleHasPermission struct {
	db.Model
	PermissionId int    `json:"permission_id"`
	RoleId       string `json:"role_id"`
}
