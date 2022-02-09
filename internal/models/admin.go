package models

import (
	"github.com/quarkcms/quark-go/pkg/framework/db"
)

// 模型
type AdminModel struct {
	db.Model
}

// 字段
type Admin struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}

// 通过用户名获取管理员信息
func (model *AdminModel) FindByUsername(username string) *Admin {
	admin := &Admin{}
	model.DB().Where("status = ?", 1).Where("username = ?", username).First(&admin)

	return admin
}
