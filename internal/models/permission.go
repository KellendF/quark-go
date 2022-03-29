package models

import (
	"time"

	"github.com/quarkcms/quark-go/pkg/framework/db"
)

// 权限
type Permission struct {
	db.Model
	Id        int
	MenuId    int
	Name      string
	GuardName string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// 获取列表
func (model *Permission) List() []map[string]interface{} {
	permissions := []Permission{}
	results := []map[string]interface{}{}

	model.DB().Find(&permissions)
	for _, v := range permissions {
		option := map[string]interface{}{
			"label": v.Name,
			"value": v.Id,
		}

		results = append(results, option)
	}

	return results
}
