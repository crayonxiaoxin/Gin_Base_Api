package models

import (
	"hello_gin_api/utils"
	"regexp"
)

// 角色
type Role struct {
	utils.BaseModel
	RoleValue    string
	RoleName     string
	Capabilities []Capability `gorm:"many2many:role_capabilities;"`
}

// 能力（权限）
type Capability struct {
	utils.BaseModel
	CapValue string
	CapName  string
}

func init() {
	// 创建或更新数据表结构
	utils.DB.AutoMigrate(&Capability{}, &Role{})
}

func (role *Role) Valid() bool {
	return role.ID > 0
}

func (cap *Capability) Valid() bool {
	return cap.ID > 0
}

// 获取角色列表
func GetRoles(page int, pageSize int) (roles []Role, count int64) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize
	utils.DB.Model(&Role{}).Preload("Capabilities").Order("id asc").Count(&count).Limit(pageSize).Offset(offset).Find(&roles)
	return
}

// 通过 id 获取 role
func GetRole(id int) (role *Role) {
	utils.DB.First(&role, id)
	return
}

// 通过 value 获取角色
func GetRoleByValue(value string) (role *Role) {
	utils.DB.Find(&role, "role_value = ?", value)
	return
}

// 添加角色
func AddRole(role *Role) utils.Result {
	result := utils.Result{}
	if len(role.RoleValue) == 0 {
		result.ResultCode = utils.ERR_PARAMS
	} else {
		matched, _ := regexp.MatchString("^[0-9a-zA-Z_]{1,}$", role.RoleValue)
		if matched {
			getrole := GetRoleByValue(role.RoleValue)
			if getrole.Valid() {
				result.ResultCode = utils.ERR_ROLE_EXISTS
			} else {
				err := utils.DB.Create(&role).Error
				if err != nil {
					result.ResultCode = utils.ERR_ROLE_ADD
				} else {
					result.ResultCode = utils.SUCCESS
				}
			}

		} else {
			result.ResultCode = utils.ERR_ROLE_REGEX
		}
	}

	return result
}

// 删除角色
func DeleteRole(id int) utils.Result {
	var result = utils.Result{}
	if id > 0 {
		role := GetRole(id)
		if role.Valid() {
			utils.DB.Delete(&role)
			result.ResultCode = utils.SUCCESS
		} else {
			result.ResultCode = utils.ERR_ROLE_NOT_EXISTS
		}
	} else {
		result.ResultCode = utils.ERR_PARAMS
	}
	return result
}

// 获取能力（權限）
func GetCapabilities(page int, pageSize int) (caps []Capability, count int64) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize
	utils.DB.Model(&Capability{}).Order("id desc").Count(&count).Limit(pageSize).Offset(offset).Find(&caps)
	return
}

// 通过 id 获取 能力（權限）
func GetCapability(id int) (cap *Capability) {
	utils.DB.First(&cap, id)
	return
}

// 通过 value 获取能力（權限）
func GetCapabilityByValue(value string) (cap *Capability) {
	utils.DB.Find(&cap, "cap_value = ?", value)
	return
}

// 添加能力（權限）
func AddCapability(cap *Capability) utils.Result {
	result := utils.Result{}
	if len(cap.CapValue) == 0 {
		result.ResultCode = utils.ERR_PARAMS
	} else {
		matched, _ := regexp.MatchString("^[0-9a-zA-Z_]{1,}$", cap.CapValue)
		if matched {
			getcap := GetCapabilityByValue(cap.CapValue)
			if getcap.Valid() {
				result.ResultCode = utils.ERR_CAP_EXISTS
			} else {
				err := utils.DB.Create(&cap).Error
				if err != nil {
					result.ResultCode = utils.ERR_CAP_ADD
				} else {
					result.ResultCode = utils.SUCCESS
				}
			}
		} else {
			result.ResultCode = utils.ERR_CAP_REGEX
		}
	}
	return result
}

// 删除角色
func DeleteCapability(id int) utils.Result {
	var result = utils.Result{}
	if id > 0 {
		cap := GetCapability(id)
		if cap.Valid() {
			utils.DB.Delete(&cap)
			result.ResultCode = utils.SUCCESS
		} else {
			result.ResultCode = utils.ERR_CAP_NOT_EXISTS
		}
	} else {
		result.ResultCode = utils.ERR_PARAMS
	}
	return result
}
