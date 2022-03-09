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
	p.Signature = "start"
	p.Description = "The service has started"

	return p
}

// 执行命令
func (p *Command) Handle() {
	fmt.Println("doing")
}
