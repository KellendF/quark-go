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
	Username  string
	Url       string
	Remark    string
	Ip        string
	Type      string
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
