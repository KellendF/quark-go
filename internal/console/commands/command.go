package commands

import (
	"fmt"
)

type Command struct {
	Signature   string
	Description string
}

// 初始化
func (p *Command) Init() *Command {
	p.Signature = "artisan"
	p.Description = "The Command Description"

	return p
}

// 执行命令
func (p *Command) Handle() {
	fmt.Println("Execute Command")
}
