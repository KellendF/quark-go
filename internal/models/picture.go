package models

import (
	"github.com/quarkcms/quark-go/pkg/framework/db"
)

// 字段
type Picture struct {
	db.Model
	Id                int
	ObjType           string
	ObjId             int
	PictureCategoryId int
	Sort              int
	Name              string
	Size              string
	Width             string
	Height            string
	Ext               string
	Path              string
	Md5               string
	Status            int
}
