package console

import (
	"fmt"
	"io/fs"
	"reflect"

	"github.com/quarkcms/quark-go/internal/console/commands"
)

type Kernel struct{}

// 注册服务
var Commands = []interface{}{
	(&commands.Command{}).Init(),
}

// 执行命令
func (p *Kernel) Run(assets fs.FS) {

	var command string
	fmt.Scanln(&command)

	for _, v := range Commands {

		// 命令标识
		signature := reflect.
			ValueOf(v).
			Elem().
			FieldByName("Signature").String()

		// 命令描述
		description := reflect.
			ValueOf(v).
			Elem().
			FieldByName("Description").String()

		if signature == command && signature != "" {
			if description != "" {
				fmt.Println(description)
			}

			v.(interface{ Handle() }).Handle()
		}
	}
}
