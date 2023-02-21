package models

import "hello_gin_api/utils"

// 用户 - 元数据
type UserMeta struct {
	utils.BaseModel
	Uid       uint   `json:"uid" gorm:"not null;default:0"`
	MetaKey   string `json:"meta_key" gorm:"not null"`
	MetaValue string `json:"meta_value"`
}

func init() {
	utils.DB.AutoMigrate(&UserMeta{})
}

func (meta *UserMeta) Valid() bool {
	return meta.ID > 0
}

// 获取用户的所有元数据
func GetUserMetas(uid interface{}) map[string]string {
	metaList := []UserMeta{}
	strList := make(map[string]string)
	utils.DB.Model(&UserMeta{}).Where("uid = ?", uid).Find(&metaList)
	for _, um := range metaList {
		strList[um.MetaKey] = um.MetaValue
	}
	return strList
}

// 获取整个结果
func GetUserMeta(usermeta *UserMeta) utils.Result {
	result := utils.Result{}
	if usermeta.Uid <= 0 || len(usermeta.MetaKey) < 1 {
		result.ResultCode = utils.ERR_PARAMS
	}
	meta := &UserMeta{}
	utils.DB.Find(&meta, "uid = ? and meta_key = ?", usermeta.Uid, usermeta.MetaKey)
	if meta.Valid() {
		result.ResultCode = utils.SUCCESS
		result.Data = *meta
	} else {
		result.ResultCode = utils.ERR_404
	}
	return result
}

// 获取整个结果 id
func GetUserMetaByID(meta_id uint) utils.Result {
	result := utils.Result{}
	if meta_id <= 0 {
		result.ResultCode = utils.ERR_PARAMS
	}
	meta := &UserMeta{}
	utils.DB.Find(&meta, "id = ? ", meta_id)
	if meta.Valid() {
		result.ResultCode = utils.SUCCESS
		result.Data = *meta
	} else {
		result.ResultCode = utils.ERR_404
	}
	return result
}

// 新增或更新
func UpdateUserMeta(usermeta *UserMeta) utils.Result {
	result := utils.Result{}
	r := GetUserMeta(usermeta)
	if r.Success() { // 存在记录，则更新
		um, ok := r.Data.(UserMeta)
		if ok {
			usermeta.ID = um.ID
		}
		err := utils.DB.Updates(&usermeta).Error
		if err != nil {
			result.ResultCode = utils.ERR_USER_META_UPDATE
		} else {
			result.ResultCode = utils.SUCCESS
		}
	} else { // 否则，新增
		err := utils.DB.Create(&usermeta).Error
		if err != nil {
			result.ResultCode = utils.ERR_USER_META_ADD
		} else {
			result.ResultCode = utils.SUCCESS
		}
	}
	return result
}

// 获取字符串
func GetUserMetaValue(usermeta *UserMeta) string {
	r := GetUserMeta(usermeta)
	if r.Success() {
		meta, ok := r.Data.(UserMeta)
		if ok {
			if len(meta.MetaValue) > 0 {
				return meta.MetaValue
			}
		}
	}
	return ""
}

// 删除
func DeleteUserMeta(uid uint, meta_key string) utils.Result {
	var result = utils.Result{}
	if uid > 0 {
		meta := GetUserMeta(&UserMeta{Uid: uid, MetaKey: meta_key})
		if meta.Success() {
			um := meta.Data.(UserMeta)
			err := utils.DB.Delete(&um).Error
			if err != nil {
				result.ResultCode = utils.ERR_USER_META_DELETE
			} else {
				result.ResultCode = utils.SUCCESS
			}
		} else {
			result.ResultCode = utils.ERR_USER_META_NOT_EXISTS
		}
	} else {
		result.ResultCode = utils.ERR_PARAMS
	}
	return result
}
