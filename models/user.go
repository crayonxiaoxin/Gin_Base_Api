package models

import (
	"hello_gin_api/utils"
)

func init() {
	// 创建或更新数据表结构
	utils.DB.AutoMigrate(&Role{}, &User{})
}

// 用户
type User struct {
	utils.BaseModel
	UserLogin string `json:"user_login" gorm:"unique;not null"`
	UserPass  string `json:"user_pass" gorm:"not null"`
	RoleId    uint   `json:"role_id" gorm:"not null"`
	Role      Role   `json:"role"`
}

// 判断用户是否有效
func (u *User) Valid() bool {
	return u.ID > 0
}

// 判断用户是否拥有该权限
// 只需要传入权限 id 或者 value
func (u *User) Can(cap *Capability) bool {
	if u.Valid() {
		if u.Role.Valid() { // 如果 role 存在
			return u.Role.Can(cap)
		} else if u.RoleId > 0 { // 如果 roleId 存在
			role := GetRole(int(u.RoleId))
			return role.Can(cap)
		}
	}
	return false
}

type UserListOptions struct {
	utils.ListOptions
	RoleId int
}

// 获取用户列表
func GetUsers(options *UserListOptions) (users []User, count int64) {
	options.Prepare()
	// utils.DB.Model(User{}).Preload("Role").Preload("Role.Capabilities").Order("id desc").Count(&count).Limit(pageSize).Offset(offset).Find(&users)
	tx := utils.DB.Model(User{}).Preload("Role")
	if options.Keyword != "" {
		tx = tx.Where("user_login like ?", options.EscKeyword())
	}
	if options.RoleId > 0 {
		tx = tx.Where("role_id = ?", options.RoleId)
	}
	tx.Order(options.Order).Count(&count).Limit(options.PageSize).Offset(options.Offset()).Find(&users)
	return
}

// 通过id获取用户
func GetUser(uid uint) (u *User) {
	utils.DB.Model(User{}).Preload("Role").First(&u, uid)
	return
}

// 添加用户
func AddUser(u *User) utils.Result {
	result := utils.Result{}
	if len(u.UserLogin) < 3 { // 用户名长度
		result.ResultCode = utils.ERR_USER_INVALID_USERNAME
	} else if len(u.UserPass) < 6 { // 密码长度
		result.ResultCode = utils.ERR_USER_INVALID_PASSWORD
	} else {
		getu := GetUserByLogin(u.UserLogin)
		if !getu.Valid() { // 用户名不存在
			encryptUserPass, _ := utils.HashedUserPass(u.UserPass) // 密码加密
			u.UserPass = encryptUserPass
			err := utils.DB.Create(&u).Error // 添加
			if err != nil {
				result.ResultCode = utils.ERR_USER_ADD
			} else {
				result.ResultCode = utils.SUCCESS
				result.Data = *u
			}

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
		getu := GetUser(u.ID)
		if getu.Valid() { // 用户存在
			if len(u.UserLogin) > 0 && u.UserLogin != getu.UserLogin { // 新用户名与旧用户名不同
				if GetUserByLogin(u.UserLogin).Valid() { // 新用户名已存在，不允许添加
					result.ResultCode = utils.ERR_USER_EXISTS
					return result
				}
			}
			if len(u.UserPass) == 0 { // 不填时，保留原密码
				u.UserPass = getu.UserPass
			} else {
				samePass := utils.EqualsUserPass(u.UserPass, getu.UserPass) // 比较密码是否变化
				if samePass {                                               // 不变化，继续使用旧密码
					u.UserPass = getu.UserPass
				} else { // 变化，重新生成密码
					encryptUserPass, _ := utils.HashedUserPass(u.UserPass) // 密码加密
					u.UserPass = encryptUserPass
				}
			}
			// 如果角色存在
			role_obj := GetRoleObj(u.RoleId)
			if role_obj.Valid() {
				u.RoleId = role_obj.ID
			} else {
				u.RoleId = getu.RoleId
			}
			err := utils.DB.Updates(&u).Error // 更新用户
			if err != nil {
				result.ResultCode = utils.ERR_USER_UPDATE
			} else {
				result.ResultCode = utils.SUCCESS
				result.Data = GetUser(u.ID) // 返回最新的用户信息
			}
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
func GetUserByLogin(user_login string) (u *User) {
	utils.DB.Find(&u, "user_login = ?", user_login)
	return
}

// 注册
func Register(u *User) utils.Result {
	return AddUser(u)
}

// 删除
func DeleteUser(uid uint) utils.Result {
	var result = utils.Result{}
	if uid > 0 {
		u := GetUser(uid)
		if u.Valid() {
			err := utils.DB.Delete(&u).Error
			if err != nil {
				result.ResultCode = utils.ERR_USER_DELETE
			} else {
				result.ResultCode = utils.SUCCESS
			}
		} else {
			result.ResultCode = utils.ERR_USER_NOT_EXISTS
		}
	} else {
		result.ResultCode = utils.ERR_PARAMS
	}
	return result
}

// 設置用戶角色
func AddRole2User(role_id uint, uid uint) utils.Result {
	var result = utils.Result{}
	if role_id == 0 || uid == 0 {
		result.ResultCode = utils.ERR_PARAMS
	} else {
		user_obj := GetUser(uid)
		if user_obj.Valid() {
			role_obj := GetRoleObj(role_id)
			if role_obj.Valid() {
				err := utils.DB.Model(&User{}).Update("RoleId", role_obj.ID).Error
				if err != nil {
					result.ResultCode = utils.ERR_USER_ADD_ROLE
				} else {
					result.ResultCode = utils.SUCCESS
				}
			} else {
				result.ResultCode = utils.ERR_ROLE_NOT_EXISTS
			}
		} else {
			result.ResultCode = utils.ERR_USER_NOT_EXISTS
		}
	}
	return result
}
