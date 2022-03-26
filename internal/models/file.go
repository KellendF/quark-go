package models

import (
	"github.com/quarkcms/quark-go/pkg/framework/db"
)

// 字段
type File struct {
	db.Model
	Id             int
	ObjType        string
	ObjId          int
	FileCategoryId int
	Sort           int
	Name           string
	Size           string
	Ext            string
	Path           string
	Md5            string
	Status         bool
}
