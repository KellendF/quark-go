package models

import (
	"github.com/quarkcms/quark-go/pkg/framework/db"
)

// 权限
type Permission struct {
	db.Model
	Id        int    `json:"id"`
	Name      string `json:"name"`
	GuardName string `json:"guard_name"`
}
