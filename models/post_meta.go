package models

import "hello_gin_api/utils"

// 文章 - 元数据
type PostMeta struct {
	utils.BaseModel
	PostId    uint   `json:"post_id" gorm:"not null;default:0"`
	MetaKey   string `json:"meta_key" gorm:"not null"`
	MetaValue string `json:"meta_value"`
}

func init() {
	utils.DB.AutoMigrate(&PostMeta{})
}

func (meta *PostMeta) Valid() bool {
	return meta.ID > 0
}

// 获取用户的所有元数据
func GetPostMetas(post_id interface{}) map[string]string {
	metaList := []PostMeta{}
	strList := make(map[string]string)
	utils.DB.Model(&PostMeta{}).Where("post_id = ?", post_id).Find(&metaList)
	for _, um := range metaList {
		strList[um.MetaKey] = um.MetaValue
	}
	return strList
}

// 获取整个结果
func GetPostMeta(postmeta *PostMeta) utils.Result {
	result := utils.Result{}
	if postmeta.PostId <= 0 || len(postmeta.MetaKey) < 1 {
		result.ResultCode = utils.ERR_PARAMS
	}
	meta := &PostMeta{}
	utils.DB.Find(&meta, "post_id = ? and meta_key = ?", postmeta.PostId, postmeta.MetaKey)
	if meta.Valid() {
		result.ResultCode = utils.SUCCESS
		result.Data = *meta
	} else {
		result.ResultCode = utils.ERR_404
	}
	return result
}

// 获取整个结果 id
func GetPostMetaByID(meta_id uint) utils.Result {
	result := utils.Result{}
	if meta_id <= 0 {
		result.ResultCode = utils.ERR_PARAMS
	}
	meta := &PostMeta{}
	utils.DB.Find(&meta, "id = ?", meta_id)
	if meta.Valid() {
		result.ResultCode = utils.SUCCESS
		result.Data = *meta
	} else {
		result.ResultCode = utils.ERR_404
	}
	return result
}

// 新增或更新
func UpdatePostMeta(postmeta *PostMeta) utils.Result {
	result := utils.Result{}
	r := GetPostMeta(postmeta)
	if r.Success() { // 存在记录，则更新
		um, ok := r.Data.(PostMeta)
		if ok {
			postmeta.ID = um.ID
		}
		utils.DB.Updates(&postmeta)
	} else { // 否则，新增
		utils.DB.Create(&postmeta)
	}
	result.ResultCode = utils.SUCCESS
	return result
}

// 获取字符串
func GetPostMetaValue(postmeta *PostMeta) string {
	r := GetPostMeta(postmeta)
	if r.Success() {
		meta, ok := r.Data.(PostMeta)
		if ok {
			if len(meta.MetaValue) > 0 {
				return meta.MetaValue
			}
		}
	}
	return ""
}

// 删除
func DeletePostMeta(post_id uint, meta_key string) utils.Result {
	var result = utils.Result{}
	if post_id > 0 {
		meta := GetPostMeta(&PostMeta{PostId: post_id, MetaKey: meta_key})
		if meta.Success() {
			pm := meta.Data.(PostMeta)
			utils.DB.Delete(&pm)
			result.ResultCode = utils.SUCCESS
		} else {
			result.ResultCode = utils.ERR_POST_META_NOT_EXISTS
		}
	} else {
		result.ResultCode = utils.ERR_PARAMS
	}
	return result
}
