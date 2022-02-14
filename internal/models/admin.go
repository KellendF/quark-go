package models

import (
	"strings"

	"github.com/go-basic/uuid"
	"github.com/quarkcms/quark-go/pkg/framework/db"
	"github.com/quarkcms/quark-go/pkg/ui/admin/utils"
)

// 字段
type Admin struct {
	db.Model
	Id       int    `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}

// 通过用户名获取管理员信息
func (model *Admin) GetAdminViaUsername(username string) *Admin {
	admin := &Admin{}
	model.DB().Where("status = ?", 1).Where("username = ?", username).First(&admin)

	return admin
}

// 通过角色获取管理员权限
func (model *Admin) GetPermissionsViaRoles(id float64) interface{} {
	var roleIds []float64
	var permissionIds []float64

	// 角色id
	(&ModelHasRole{}).DB().Where("model_id", id).Where("model_type", "admin").Pluck("id", &roleIds)

	if roleIds == nil {
		return nil
	}

	// 角色权限id
	(&RoleHasPermission{}).DB().Where("role_id in (?)", roleIds).Pluck("id", &permissionIds)

	if permissionIds == nil {
		return nil
	}

	// 角色权限列表
	permission := &Permission{}
	permission.DB().Where("id in (?)", &permissionIds).Find(&permission)

	return permission
}

// 获取管理员角色
func (model *Admin) GetRoles(id float64) *ModelHasRole {
	modelHasRole := &ModelHasRole{}
	modelHasRole.DB().Where("model_id", id).Where("model_type", "admin").First(&modelHasRole)

	return modelHasRole
}

// 获取管理员权限菜单
func (model *Admin) GetMenus(adminId float64) interface{} {

	menu := &Menu{}
	var menus []Menu
	var menuKey int

	if adminId == 1 {
		menu.DB().Where("status = ?", 1).Where("guard_name", "admin").Order("sort asc").Find(&menus)
	} else {
		var menuIds []float64
		permissions := model.GetPermissionsViaRoles(adminId)

		if permissions != nil {
			for key, v := range permissions.([]map[string]float64) {
				menuIds[key] = v["menu_id"]
			}
		}

		var pids1 []float64

		// 三级查询列表
		menu.DB().
			Where("status = ?", 1).
			Where("id in (?)", menuIds).
			Where("pid <> ?", 0).
			Order("sort asc").
			Find(&menus)

		for key, v := range menus {
			if v.Pid != 0 {
				pids1[key] = v.Pid
			}
			menuKey = key
		}

		var pids2 []float64
		var menu2 []Menu

		// 二级查询列表
		menu.DB().
			Where("status = ?", 1).
			Where("id in (?)", pids1).
			Where("pid <> ?", 0).
			Order("sort asc").
			Find(&menu2)

		for key, v := range menu2 {
			if v.Pid != 0 {
				pids2[key] = v.Pid
			}

			menuKey = menuKey + key
			menus[menuKey] = v
		}

		var menu3 []Menu

		// 一级查询列表
		menu.DB().
			Where("status = ?", 1).
			Where("id in (?)", pids2).
			Where("pid", 0).
			Order("sort asc").
			Find(&menu3)

		for key, v := range menu3 {
			menuKey = menuKey + key
			menus[menuKey] = v
		}
	}

	for key, v := range menus {
		menus[key].Key = uuid.New()
		menus[key].Locale = strings.Replace(v.Path, "/", ".", -1)

		if v.Show == 1 {
			menus[key].HideInMenu = false
		} else {
			menus[key].HideInMenu = true
		}

		if v.Type == "engine" {
			menus[key].Path = "/index?api=" + v.Path
		}
	}

	return utils.ListToTree(utils.StructToMap(menus).([]interface{}), "id", "pid", "routes", 0)
}
