package db

import (
	"fmt"

	"github.com/quarkcms/quark-go/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var SqlConn *gorm.DB

// 初始化
func init() {

	username := config.Get("database.mysql.username")
	password := config.Get("database.mysql.password")
	host := config.Get("database.mysql.host")
	port := config.Get("database.mysql.port")
	database := config.Get("database.mysql.database")
	charset := config.Get("database.mysql.charset")

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=" + charset + "&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}

	SqlConn = conn
}

// 数据库链接句柄
func Conn() *gorm.DB {
	return SqlConn
}
