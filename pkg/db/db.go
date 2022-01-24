package db

import (
	"fmt"

	"github.com/quarkcms/quark-go/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Model结构体
type Model struct {
	Conn *gorm.DB
}

// 数据库链接句柄
func (model *Model) conn() *gorm.DB {

	if model.Conn != nil {
		return model.Conn
	}

	username := config.Get("database.mysql.username")
	password := config.Get("database.mysql.password")
	host := config.Get("database.mysql.host")
	port := config.Get("database.mysql.port")
	database := config.Get("database.mysql.database")
	charset := config.Get("database.mysql.charset")

	if username != "" && host != "" && port != "" && database != "" && charset != "" {
		dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=" + charset + "&parseTime=True&loc=Local"
		conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			fmt.Println(err)
		}

		model.Conn = conn
	}

	return model.Conn
}

// 数据库实例
func (model *Model) DB() *gorm.DB {

	return model.conn()
}
