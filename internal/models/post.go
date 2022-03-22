package models

import (
	"github.com/quarkcms/quark-go/pkg/framework/db"
)

// 字段
type Post struct {
	db.Model
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}
