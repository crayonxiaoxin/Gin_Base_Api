package models

import (
	"fmt"
	"hello_gin_api/utils"
	"mime"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type Media struct {
	utils.BaseModel
	Name       string `json:"name"`
	Path       string `json:"path"`
	Mime       string `json:"mime"`
	Ext        string `json:"ext"`
	Size       int64  `json:"size"`
	OriginName string `json:"origin"`
	Uid        int64  `json:"uid"`
}

func init() {
	utils.DB.AutoMigrate(&Media{})
}

// 判断媒体是否有效
func (m *Media) Valid() bool {
	return m.ID > 0
}

// 获取所有媒体
func GetAllMedia(page int, pageSize int) (media []Media, count int64) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize
	utils.DB.Model(Media{}).Order("id desc").Count(&count).Limit(pageSize).Offset(offset).Find(&media)
	return
}

// 通过id获取媒体
func GetMedia(id int) (m *Media) {
	utils.DB.First(&m, id)
	return
}

// 添加媒体文件
func AddMedia(m *Media) {
	utils.DB.Create(&m)
}

// 删除
func DeleteMedia(id int) utils.Result {
	var result = utils.Result{}
	if id > 0 {
		m := GetMedia(id)
		if m.Valid() {
			utils.DB.Delete(&m)
			utils.RemoveFile(m.Path)
			result.ResultCode = utils.SUCCESS
		} else {
			result.ResultCode = utils.ERR_UPLOAD_FILE_NOT_EXISTS
		}
	} else {
		result.ResultCode = utils.ERR_PARAMS
	}
	return result
}

// 上传媒体文件
func UploadMedia(fh *multipart.FileHeader, uid int64) utils.Result {
	result := utils.Result{}
	// 文件名
	filename := filepath.Base(fh.Filename)
	// 后缀
	ext := filepath.Ext(filename)
	// 重命名
	newName := strconv.FormatInt(time.Now().UnixNano(), 10) + ext
	// Mime类型
	fileMime := mime.TypeByExtension(ext)
	// 文件大小
	size := fh.Size
	// 创建保存目录
	if !utils.PathExists("uploads") {
		err3 := os.Mkdir("uploads", os.ModePerm)
		if err3 != nil {
			result.ResultCode = utils.ERR_UPLOAD_MKDIR
			return result
		}
	}
	// 保存路径
	dst := fmt.Sprintf("uploads/%v", newName)
	// 保存
	err2 := utils.SaveFile(fh, dst)
	if err2 != nil {
		rc := utils.ERR_PARAMS
		rc.Msg = rc.Msg + ": " + err2.Error()
		result.ResultCode = rc
	} else {
		data := Media{
			Name:       newName,
			Path:       dst,
			Mime:       fileMime,
			Ext:        ext,
			Size:       size,
			OriginName: filename,
			Uid:        uid,
		}
		AddMedia(&data)
		result.ResultCode = utils.SUCCESS
		result.Data = data
	}
	return result
}
