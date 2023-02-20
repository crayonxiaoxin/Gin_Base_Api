package models

import (
	"fmt"
	"hello_gin_api/utils"
)

func init() {
	// 创建或更新数据表结构
	utils.DB.AutoMigrate(&User{})
}

type User struct {
	utils.BaseModel
	Username string `json:"username"`
	Password string `json:"password"`
}

// 判断用户是否有效
func (u *User) Valid() bool {
	return u.ID > 0
}

// 获取所有用户
func GetAllUsers(page int, pageSize int) (users []User, count int64) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize
	utils.DB.Model(User{}).Order("id desc").Count(&count).Limit(pageSize).Offset(offset).Find(&users)
	return
}

// 通过id获取用户
func GetUser(uid int) (u *User) {
	utils.DB.First(&u, uid)
	return
}

// 添加用户
func AddUser(u *User) utils.Result {
	result := utils.Result{}
	if len(u.Username) < 3 { // 用户名长度
		result.ResultCode = utils.ERR_USER_INVALID_USERNAME
	} else if len(u.Password) < 6 { // 密码长度
		result.ResultCode = utils.ERR_USER_INVALID_PASSWORD
	} else {
		getu := GetUserByLogin(u.Username)
		if !getu.Valid() { // 用户名不存在
			encryptPassword, _ := utils.HashedPassword(u.Password) // 密码加密
			u.Password = encryptPassword
			fmt.Println(u)
			fmt.Println(&u)
			utils.DB.Create(&u) // 添加
			result.ResultCode = utils.SUCCESS
			result.Data = *u
		} else {
			result.ResultCode = utils.ERR_USER_EXISTS
		}
	}
	return result
}

// 更新用户
func UpdateUser(u *User) utils.Result {
	var result = utils.Result{}
	if u.ID > 0 { // id 是有效的
		getu := GetUser(int(u.ID))
		if getu.Valid() { // 用户存在
			if u.Username != getu.Username { // 新用户名与旧用户名不同
				if GetUserByLogin(u.Username).Valid() { // 新用户名已存在，不允许添加
					result.ResultCode = utils.ERR_USER_EXISTS
					return result
				}
			}
			if len(u.Password) == 0 { // 不填时，保留原密码
				u.Password = getu.Password
			} else {
				samePass := utils.EqualsPassword(u.Password, getu.Password) // 比较密码是否变化
				if samePass {                                               // 不变化，继续使用旧密码
					u.Password = getu.Password
				} else { // 变化，重新生成密码
					encryptPassword, _ := utils.HashedPassword(u.Password) // 密码加密
					u.Password = encryptPassword
				}
			}
			utils.DB.Updates(u) // 更新用户
			result.ResultCode = utils.SUCCESS
			result.Data = GetUser(int(u.ID)) // 返回最新的用户信息
		} else {
			result.ResultCode = utils.ERR_USER_NOT_EXISTS
		}
	} else {
		result.ResultCode = utils.ERR_PARAMS
		result.Data = u
	}
	return result
}

// 通过用户名获取用户
func GetUserByLogin(username string) (u *User) {
	utils.DB.Find(&u, "username = ?", username)
	return
}

// 注册
func Register(u *User) utils.Result {
	return AddUser(u)
}

// 删除
func DeleteUser(uid int) utils.Result {
	var result = utils.Result{}
	if uid > 0 {
		u := GetUser(uid)
		if u.Valid() {
			utils.DB.Delete(&u)
			result.ResultCode = utils.SUCCESS
		} else {
			result.ResultCode = utils.ERR_USER_NOT_EXISTS
		}
	} else {
		result.ResultCode = utils.ERR_PARAMS
	}
	return result
}
