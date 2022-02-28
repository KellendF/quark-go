package admin

import (
	"github.com/quarkcms/quark-go/internal/admin/dashboards"
	"github.com/quarkcms/quark-go/internal/admin/resources"
)

// 注册服务
var Providers = []interface{}{
	&dashboards.Index{},
	&resources.Admin{},
}
