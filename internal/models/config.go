package models

import (
	"github.com/quarkcms/quark-go/pkg/framework/db"
)

// 字段
type Config struct {
	db.Model
	Id        int
	Title     string
	Type      string
	Name      string
	Sort      int
	GroupName string
	Value     string
	Remark    string
	Status    int
}
