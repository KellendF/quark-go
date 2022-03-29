package controllers

import (
	"crypto/md5"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gofiber/fiber/v2"
	"github.com/quarkcms/quark-go/internal/models"
	"github.com/quarkcms/quark-go/pkg/framework/db"
	"github.com/quarkcms/quark-go/pkg/framework/msg"
	"github.com/quarkcms/quark-go/pkg/framework/rand"
	"github.com/quarkcms/quark-go/pkg/ui/admin/utils"
)

type Picture struct{}

// 编辑器图片选择
func (p *Picture) GetLists(c *fiber.Ctx) {

}

// 上传图片
func (p *Picture) Upload(c *fiber.Ctx) error {
	var result error

	if utils.WebConfig("OSS_OPEN") == "1" {
		result = p.OssUpload(c)
	} else {
		result = p.LocalUpload(c)
	}

	return result
}

// 通过base64字符串上传图片
func (p *Picture) UploadFromBase64(c *fiber.Ctx) {

}

// 图片上传到本地
func (p *Picture) LocalUpload(c *fiber.Ctx) error {
	file, _ := c.FormFile("file")
	limitW := c.Query("limitW")
	limitH := c.Query("limitH")

	limitType := []string{
		"image/jpg",
		"image/jpeg",
		"image/png",
		"image/gif",
	}

	typeAllowed := false

	for _, v := range file.Header["Content-Type"] {
		for _, limit := range limitType {
			if v == limit {
				typeAllowed = true
			}
		}
	}

	f, err := file.Open()

	if err != nil {
		return msg.Error(err.Error(), "")
	}

	defer func() {
		e := f.Close()
		if err == nil {
			err = e
		}
	}()

	// 限制格式
	if typeAllowed == false {
		return msg.Error("只能上传jpg,jpeg,png,gif格式图片!", "")
	}

	imageConfig, _, err := image.DecodeConfig(f)

	if err != nil {
		return msg.Error(err.Error(), "")
	}

	// 限制宽高
	if limitW != "" && limitH != "" {
		w, err := strconv.Atoi(limitW)
		if err != nil {
			return msg.Error(err.Error(), "")
		}

		h, err := strconv.Atoi(limitH)
		if err != nil {
			return msg.Error(err.Error(), "")
		}

		if imageConfig.Width != w || imageConfig.Height != h {
			return msg.Error("请上传 "+limitW+"*"+limitH+" 尺寸的图片", "")
		}
	}

	filePath := "./storage/app/public/pictures/" + time.Now().Format("20060102") + "/"
	fileName := file.Filename
	fileSize := file.Size

	fileNames := strings.Split(fileName, ".")
	if len(fileNames) <= 1 {
		return msg.Error("无法获取文件扩展名！", "")
	}

	fileExt := fileNames[len(fileNames)-1]
	fileNewName := rand.MakeAlphanumeric(40) + "." + fileExt

	// 文件md5值
	body, err := ioutil.ReadAll(f)
	if err != nil {
		return msg.Error(err.Error(), "")
	}
	fileMd5 := fmt.Sprintf("%x", md5.Sum(body))

	picture := map[string]interface{}{}
	(&db.Model{}).Model(&models.Picture{}).Where("md5", fileMd5).Where("name", fileName).First(&picture)

	result := map[string]interface{}{}

	if len(picture) > 0 {
		result = map[string]interface{}{
			"id":   picture["id"],
			"name": picture["name"],
			"url":  c.BaseURL() + strings.Replace(picture["path"].(string), "./storage/app/public", "/storage", -1),
			"size": picture["size"],
		}

		return msg.Success("上传成功！", "", result)
	}

	// 不存在路径，则创建
	if utils.PathExist(filePath) == false {
		err := os.MkdirAll(filePath, 0666)
		if err != nil {
			return msg.Error(err.Error(), "")
		}
	}

	// 保存文件
	err = c.SaveFile(file, filePath+fileNewName)
	if err != nil {
		return msg.Error(err.Error(), "")
	}

	id := (&models.Picture{}).InsertGetId(map[string]interface{}{
		"obj_type": "ADMINID",
		"obj_id":   utils.Admin(c, "id"),
		"name":     fileName,
		"size":     fileSize,
		"md5":      fileMd5,
		"path":     filePath + fileNewName,
		"width":    imageConfig.Width,
		"height":   imageConfig.Height,
		"ext":      fileExt,
	})

	if id == 0 {
		return msg.Error("上传失败！", "")
	}

	result = map[string]interface{}{
		"id":   id,
		"name": fileName,
		"url":  c.BaseURL() + strings.Replace(filePath+fileNewName, "./storage/app/public", "/storage", -1),
		"size": fileSize,
	}

	return msg.Success("上传成功！", "", result)
}

// 图片上传到阿里云OSS
func (p *Picture) OssUpload(c *fiber.Ctx) error {

	accessKeyId := utils.WebConfig("OSS_ACCESS_KEY_ID")
	accessKeySecret := utils.WebConfig("OSS_ACCESS_KEY_SECRET")
	endpoint := utils.WebConfig("OSS_ENDPOINT")
	ossBucket := utils.WebConfig("OSS_BUCKET")

	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		// HandleError(err)
	}

	bucket, err := client.Bucket(ossBucket)
	if err != nil {
		// HandleError(err)
	}

	err = bucket.PutObjectFromFile("my-object", "LocalFile")
	if err != nil {
		// HandleError(err)
	}

	return msg.Success("上传成功！", "", "")
}

// 图片下载
func (p *Picture) Download(c *fiber.Ctx) {

}
