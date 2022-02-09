package db

import (
	"github.com/quarkcms/quark-go/pkg/framework/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Model结构体
type Model struct{}

var conn *gorm.DB

func init() {
	username := config.Get("database.mysql.username")
	password := config.Get("database.mysql.password")
	host := config.Get("database.mysql.host")
	port := config.Get("database.mysql.port")
	database := config.Get("database.mysql.database")
	charset := config.Get("database.mysql.charset")

	if username != "" && host != "" && port != "" && database != "" && charset != "" {
		dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=" + charset + "&parseTime=True&loc=Local"
		conn, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}
}

// 数据库实例
func (model *Model) DB() *gorm.DB {

	return conn
}
