package models

import (
	"github.com/quarkcms/quark-go/pkg/framework/db"
)

// 字段
type ActionLog struct {
	db.Model
	Id       int    `json:"id"`
	ObjectId int    `json:"object_id"`
	Url      string `json:"url"`
	Remark   string `json:"remark"`
	Ip       string `json:"ip"`
	Type     string `json:"type"`
	Status   bool   `json:"status"`
}
