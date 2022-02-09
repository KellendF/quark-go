package resource

import (
	"github.com/quarkcms/quark-go/internal/http/controllers/admin/resources"
	"github.com/quarkcms/quark-go/pkg/ui/resource"
)

// 资源
func Resource() map[string]resource.ResourceInterface {

	return map[string]resource.ResourceInterface{
		"admin": &resources.Admin{},
	}
}
