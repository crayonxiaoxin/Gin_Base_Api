package models

import (
	"hello_gin_api/utils"
	"regexp"
)

// 角色
type Role struct {
	utils.BaseModel
	RoleValue    string `gorm:"unique"`
	RoleName     string
	Capabilities []Capability `gorm:"many2many:role_capabilities;"`
}

// 能力（权限）
type Capability struct {
	utils.BaseModel
	CapValue string `gorm:"unique"`
	CapName  string
}

var (
	// 用戶
	USER_READ   = Capability{CapValue: "list_users", CapName: "查看用戶列表"}
	USER_ADD    = Capability{CapValue: "add_users", CapName: "添加用戶"}
	USER_EDIT   = Capability{CapValue: "edit_users", CapName: "修改用戶"}
	USER_DELETE = Capability{CapValue: "delete_users", CapName: "刪除用戶"}

	// 角色
	ROLE_READ   = Capability{CapValue: "list_roles", CapName: "查看角色列表"}
	ROLE_ADD    = Capability{CapValue: "add_roles", CapName: "添加角色"}
	ROLE_EDIT   = Capability{CapValue: "edit_roles", CapName: "修改角色"}
	ROLE_DELETE = Capability{CapValue: "delete_roles", CapName: "刪除角色"}

	// 權限
	CAP_READ   = Capability{CapValue: "list_caps", CapName: "查看權限列表"}
	CAP_ADD    = Capability{CapValue: "add_caps", CapName: "添加權限"}
	CAP_EDIT   = Capability{CapValue: "edit_caps", CapName: "修改權限"}
	CAP_DELETE = Capability{CapValue: "delete_caps", CapName: "刪除權限"}

	// 文章
	POST_READ   = Capability{CapValue: "list_posts", CapName: "查看文章列表"}
	POST_ADD    = Capability{CapValue: "add_posts", CapName: "添加文章"}
	POST_EDIT   = Capability{CapValue: "edit_posts", CapName: "修改文章"}
	POST_DELETE = Capability{CapValue: "delete_posts", CapName: "刪除文章"}

	// 媒體
	MEDIA_READ   = Capability{CapValue: "list_media", CapName: "查看媒體列表"}
	MEDIA_ADD    = Capability{CapValue: "add_media", CapName: "添加媒體"}
	MEDIA_EDIT   = Capability{CapValue: "edit_media", CapName: "修改媒體"}
	MEDIA_DELETE = Capability{CapValue: "delete_media", CapName: "刪除媒體"}

	ROLE_ADMIN      = "admin"
	ROLE_SUBSCRIBER = "subscriber"
)

func init() {
	// 创建或更新数据表结构
	utils.DB.AutoMigrate(&Capability{}, &Role{})

	// 1. 预设角色：管理员、普通用户
	// 2. 预设管理员
	// TODO: 预设能力：不可刪除，可改名稱；用戶自身擁有對自身相關的 CRUD 權限
	preset_caps := []Capability{
		// 用戶
		USER_READ,
		USER_ADD,
		USER_EDIT,
		USER_DELETE,
		// 角色
		ROLE_READ,
		ROLE_ADD,
		ROLE_EDIT,
		ROLE_DELETE,
		// 權限
		CAP_READ,
		CAP_ADD,
		CAP_EDIT,
		CAP_DELETE,
		// 文章
		POST_READ,
		POST_ADD,
		POST_EDIT,
		POST_DELETE,
		// 媒體
		MEDIA_READ,
		MEDIA_ADD,
		MEDIA_EDIT,
		MEDIA_DELETE,
	}
	preset_caps_subscriber := []Capability{
		POST_READ,
		POST_ADD,
	}

	// 预设角色
	preset_role_admin := Role{RoleValue: ROLE_ADMIN, RoleName: "管理員"}
	preset_role_subscriber := Role{RoleValue: ROLE_SUBSCRIBER, RoleName: "普通用戶"}
	preset_roles := []Role{
		preset_role_admin,
		preset_role_subscriber,
	}
	for _, r := range preset_roles {
		AddRole(&r) // 先添加预设角色，因为 AddCapability 会自动将权限添加到 admin
	}
	for _, c := range preset_caps { // 预设能力添加到数据库
		AddCapability(&c) // 这里会自动将权限添加到 admin
	}
	// 为普通用户预设权限
	AddCapabilities2Role(preset_role_subscriber, preset_caps_subscriber)

	// 添加初始管理员
	role_admin := GetRoleByValue(ROLE_ADMIN)
	AddUser(&User{UserLogin: "eftech", UserPass: "@ne2Nine", RoleId: role_admin.ID})
}

func (role *Role) Valid() bool {
	return role.ID > 0
}

// 判断角色是否拥有该权限
// 只需要传入权限 id 或者 value
func (role *Role) Can(cap *Capability) bool {
	if role.Valid() {
		if cap.Valid() { // 如果有ID
			for _, c := range role.Capabilities {
				if c.ID == cap.ID {
					return true
				}
			}
		} else if len(cap.CapValue) > 0 { // 如果有 value
			for _, c := range role.Capabilities {
				if c.CapValue == cap.CapValue {
					return true
				}
			}
		}
	}
	return false
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
	utils.DB.Preload("Capabilities").First(&role, id)
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
					AddCapability2Role(cap.CapValue, ROLE_ADMIN) // 每个新权限都要为管理员添加
				}
			}
		} else {
			result.ResultCode = utils.ERR_CAP_REGEX
		}
	}
	return result
}

// 删除能力
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

// 为角色添加能力
// 支持傳入 uint、string、struct
func AddCapability2Role(cap interface{}, role interface{}) utils.Result {
	result := utils.Result{}
	cap_obj := GetCapObj(cap)
	role_obj := GetRoleObj(role)
	if cap_obj.Valid() && role_obj.Valid() {
		err := utils.DB.Model(&role_obj).Association("Capabilities").Append(&cap_obj)
		if err != nil {
			result.ResultCode = utils.ERR_ROLE_CAP_GRANT_FAILED
		} else {
			result.ResultCode = utils.SUCCESS
		}
	} else {
		result.ResultCode = utils.ERR_PARAMS
	}
	return result
}

// 为角色添加能力
// 支持傳入 uint、string、struct
func AddCapabilities2Role(role interface{}, cap []Capability) {
	result := utils.Result{}
	role_obj := GetRoleObj(role)
	if role_obj.Valid() {
		for _, c := range cap {
			cap_obj := GetCapObj(c)
			if cap_obj.Valid() {
				err := utils.DB.Model(&role_obj).Association("Capabilities").Append(&cap_obj)
				if err != nil {
					result.ResultCode = utils.ERR_ROLE_CAP_GRANT_FAILED
				} else {
					result.ResultCode = utils.SUCCESS
				}
			}
		}
	}
}

// 为角色刪除某項能力
// 支持傳入 uint、string、struct
func DeleteCapabilityFromRole(cap interface{}, role interface{}) utils.Result {
	result := utils.Result{}
	cap_obj := GetCapObj(cap)
	role_obj := GetRoleObj(role)
	if cap_obj.Valid() && role_obj.Valid() {
		err := utils.DB.Model(&role_obj).Association("Capabilities").Delete(&cap_obj)
		if err != nil {
			result.ResultCode = utils.ERR_ROLE_CAP_REMOVE_FAILED
		} else {
			result.ResultCode = utils.SUCCESS
		}
	} else {
		result.ResultCode = utils.ERR_PARAMS
	}
	return result
}

// 通过传入的参数类型获取能力
// 支持傳入 uint、string、struct
func GetCapObj(cap interface{}) Capability {
	var cap_obj Capability
	cap_id_tmp, ok := cap.(uint) // 支持 uint
	if ok {
		cap_obj = *GetCapability(int(cap_id_tmp))
	} else {
		cap_val_tmp, ok := cap.(string) // 支持 string
		if ok && len(cap_val_tmp) > 0 {
			cap_obj = *GetCapabilityByValue(cap_val_tmp)
		} else {
			cap_obj_tmp, ok := cap.(Capability) // 支持 struct，只要有 ID 或 value 其一即可
			if ok {
				if cap_obj_tmp.Valid() {
					cap_obj = *GetCapability(int(cap_obj_tmp.ID))
				} else if len(cap_obj_tmp.CapValue) > 0 {
					cap_obj = *GetCapabilityByValue(cap_obj_tmp.CapValue)
				}
			}
		}
	}
	return cap_obj
}

// 通过传入的参数类型获取角色
// 支持傳入 uint、string、struct
func GetRoleObj(role interface{}) Role {
	var role_obj Role
	role_id_tmp, ok := role.(uint) // 支持 uint
	if ok {
		role_obj = *GetRole(int(role_id_tmp))
	} else {
		role_val_tmp, ok := role.(string) // 支持 string
		if ok && len(role_val_tmp) > 0 {
			role_obj = *GetRoleByValue(role_val_tmp)
		} else {
			role_obj_tmp, ok := role.(Role) // 支持 struct
			if ok {
				if role_obj_tmp.Valid() {
					role_obj = *GetRole(int(role_obj_tmp.ID))
				} else if len(role_obj_tmp.RoleValue) > 0 {
					role_obj = *GetRoleByValue(role_obj_tmp.RoleValue)
				}
			}
		}
	}
	return role_obj
}
