package models

import (
	"github.com/quarkcms/quark-go/pkg/framework/db"
)

// 字段
type Picture struct {
	db.Model
	Id                int    `json:"id"`
	ObjType           string `json:"obj_type"`
	ObjId             int    `json:"obj_id"`
	PictureCategoryId int    `json:"picture_category_id"`
	Sort              int    `json:"sort"`
	Name              string `json:"name"`
	Size              string `json:"size"`
	Width             string `json:"width"`
	Height            string `json:"height"`
	Ext               string `json:"ext"`
	Path              string `json:"path"`
	Md5               string `json:"md5"`
	Status            bool   `json:"status"`
}
