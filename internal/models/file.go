package models

import (
	"strconv"
	"time"

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
	Status         int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// 插入数据并返回ID
func (model *File) InsertGetId(data map[string]interface{}) int {
	size := strconv.FormatInt(data["size"].(int64), 10)
	picture := Picture{
		ObjType: data["obj_type"].(string),
		ObjId:   data["obj_id"].(int),
		Name:    data["name"].(string),
		Size:    size,
		Md5:     data["md5"].(string),
		Path:    data["path"].(string),
		Ext:     data["ext"].(string),
		Status:  1,
	}
	model.DB().Create(&picture)

	return picture.Id
}
