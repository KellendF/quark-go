package resource

import (
	"github.com/quarkcms/quark-go/internal/http/controllers/admin/resources/administrator"
	"github.com/quarkcms/quark-go/pkg/resource"
)

// 资源
func Resource() map[string]resource.ResourceInterface {

	return map[string]resource.ResourceInterface{
		"administrator": &administrator.Component{},
	}
}
