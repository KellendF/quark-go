package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Println(err)
	}

	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(err)
	fmt.Println(status)
}
