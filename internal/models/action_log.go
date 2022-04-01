package models

import (
	"time"

	"github.com/quarkcms/quark-go/pkg/framework/db"
)

// 字段
type ActionLog struct {
	db.Model
	Id        int
	ObjectId  int
	Username  string `gorm:"<-:false"`
	Url       string
	Remark    string
	Ip        string
	Type      string
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// 插入数据
func (model *ActionLog) Insert(data map[string]interface{}) {

	log := ActionLog{
		ObjectId: data["obj_id"].(int),
		Url:      data["url"].(string),
		Ip:       data["ip"].(string),
		Type:     data["type"].(string),
		Status:   1,
	}

	model.DB().Create(&log)
}
