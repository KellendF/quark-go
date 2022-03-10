package commands

import (
	"fmt"
)

type Install struct {
	Command
}

// 初始化
func (p *Install) Init() *Install {
	p.Signature = "install"
	p.Description = "The Command Description"

	return p
}

// 执行命令
func (p *Install) Handle() {
	fmt.Println("Execute Command")
}
