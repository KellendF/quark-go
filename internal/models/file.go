package models

import (
	"github.com/quarkcms/quark-go/pkg/framework/db"
)

// 字段
type File struct {
	db.Model
	Id             int    `json:"id"`
	ObjType        string `json:"obj_type"`
	ObjId          int    `json:"obj_id"`
	FileCategoryId int    `json:"file_category_id"`
	Sort           int    `json:"sort"`
	Name           string `json:"name"`
	Size           string `json:"size"`
	Ext            string `json:"ext"`
	Path           string `json:"path"`
	Md5            string `json:"md5"`
	Status         bool   `json:"status"`
}
