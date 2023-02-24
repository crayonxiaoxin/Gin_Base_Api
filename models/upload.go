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

// 媒体
type Media struct {
	utils.BaseModel
	Name       string `json:"name" gorm:"not null"`
	Path       string `json:"path" gorm:"not null"`
	Mime       string `json:"mime" gorm:"not null"`
	Ext        string `json:"ext"`
	Size       int64  `json:"size" gorm:"not null"`
	OriginName string `json:"origin"`
	Uid        uint   `json:"uid" gorm:"not null;default:0"`
}

func init() {
	utils.DB.AutoMigrate(&Media{})
}

// 判断媒体是否有效
func (m *Media) Valid() bool {
	return m.ID > 0
}

type MediaListOptions struct {
	utils.ListOptions
	Uid int
}

// 获取媒体列表
func GetAllMedia(options *MediaListOptions) (media []Media, count int64) {
	options.Prepare()
	tx := utils.DB.Model(Media{})
	if options.Keyword != "" {
		tx = tx.Where("name like ? or mime like ?", options.EscKeyword(), options.EscKeyword())
	}
	if options.Uid > 0 {
		tx = tx.Where("uid = ?", options.Uid)
	}
	tx.Order(options.Order).Count(&count).Limit(options.PageSize).Offset(options.Offset()).Find(&media)
	return
}

// 通过id获取媒体
func GetMedia(id int) (m *Media) {
	utils.DB.First(&m, id)
	return
}

// 删除
func DeleteMedia(id int) utils.Result {
	var result = utils.Result{}
	if id > 0 {
		m := GetMedia(id)
		if m.Valid() {
			err := utils.DB.Delete(&m).Error
			if err != nil {
				result.ResultCode = utils.ERR_UPLOAD_FILE_DELETE
			} else {
				utils.RemoveFile(m.Path)
				result.ResultCode = utils.SUCCESS
			}

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
			Uid:        uint(uid),
		}
		err := utils.DB.Create(&data).Error
		if err != nil {
			result.ResultCode = utils.ERR_UPLOAD_FILE_ADD
		} else {
			result.ResultCode = utils.SUCCESS
			result.Data = data
		}
	}
	return result
}
