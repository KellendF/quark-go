package models

import (
	"github.com/quarkcms/quark-go/pkg/framework/db"
)

// 字段
type ActionLog struct {
	db.Model
	Id       int
	ObjectId int
	Url      string
	Remark   string
	Ip       string
	Type     string
	Status   bool
}
