package db

import (
	"github.com/quarkcms/quark-go/pkg/framework/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 结构体
type Model struct{}

var conn *gorm.DB

func init() {
	username := config.Get("database.mysql.username").(string)
	password := config.Get("database.mysql.password").(string)
	host := config.Get("database.mysql.host").(string)
	port := config.Get("database.mysql.port").(string)
	database := config.Get("database.mysql.database").(string)
	charset := config.Get("database.mysql.charset").(string)

	if username != "" && host != "" && port != "" && database != "" && charset != "" {
		dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=" + charset + "&parseTime=True&loc=Local"
		conn, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}
}

// 实例
func (p *Model) DB() *gorm.DB {

	return conn
}

// 模型
func (p *Model) Model(model interface{}) *gorm.DB {

	return conn.Model(model)
}
