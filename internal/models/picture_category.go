package models

import (
	"github.com/quarkcms/quark-go/pkg/framework/db"
)

// 字段
type PictureCategory struct {
	db.Model
	Id          int
	ObjType     string
	ObjId       int
	Title       string
	Sort        int
	Description string
}
