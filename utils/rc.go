package utils

// 状态码结构
type ResultCode struct {
	Rc  int    `json:"rc"`
	Msg string `json:"msg"`
}

func (rc *ResultCode) Success() bool {
	return rc.Rc == 0
}

// 统一返回 json 结构
type Result struct {
	ResultCode
	Data interface{} `json:"data"`
}

// 状态码
var (
	SUCCESS    = ResultCode{0, "Success"}
	ERR_PARAMS = ResultCode{1, "參數錯誤"}

	ERR_LOGIN_PASSWORD = ResultCode{10, "密碼錯誤"}
	ERR_REGISTER       = ResultCode{11, "註冊失敗"}

	ERR_UNKNOWN = ResultCode{99, "未知錯誤"}

	ERR_404 = ResultCode{404, "Not found"}

	// token
	ERR_TOKEN_GENERATE  = ResultCode{1001, "無法生成 Token"}
	ERR_TOKEN_MALFORMED = ResultCode{1002, "Token 錯誤"}
	ERR_TOKEN_EXPIRED   = ResultCode{1003, "Token 已過期"}
	ERR_TOKEN_UNKNOWN   = ResultCode{1004, "未知 Token"}

	// user
	ERR_USER_NOT_EXISTS       = ResultCode{1101, "用戶不存在"}
	ERR_USER_EXISTS           = ResultCode{1102, "用戶已存在"}
	ERR_USER_INVALID_USERNAME = ResultCode{1103, "賬戶錯誤"}
	ERR_USER_INVALID_PASSWORD = ResultCode{1104, "密碼錯誤"}
	ERR_USER_ADD              = ResultCode{1105, "添加用戶失敗"}
	ERR_USER_UPDATE           = ResultCode{1106, "更新用戶失敗"}
	ERR_USER_DELETE           = ResultCode{1107, "刪除用戶失敗"}
	ERR_USER_ADD_ROLE         = ResultCode{1108, "設置用戶角色失敗"}

	ERR_USER_META_NOT_EXISTS = ResultCode{1111, "用戶元數據不存在"}
	ERR_USER_META_ADD        = ResultCode{1112, "添加用戶元數據失敗"}
	ERR_USER_META_UPDATE     = ResultCode{1113, "更新用戶元數據失敗"}
	ERR_USER_META_DELETE     = ResultCode{1114, "刪除用戶元數據失敗"}

	// upload
	ERR_UPLOAD_MKDIR           = ResultCode{1201, "無法創建文件夾"}
	ERR_UPLOAD_FILE_NOT_EXISTS = ResultCode{1202, "文件不存在"}
	ERR_UPLOAD_FILE_ADD        = ResultCode{1203, "添加文件失敗"}
	ERR_UPLOAD_FILE_DELETE     = ResultCode{1204, "刪除文件失敗"}

	// post
	ERR_POST_EMPTY_TITLE = ResultCode{1301, "文章標題不能為空"}
	ERR_POST_NOT_EXISTS  = ResultCode{1302, "文章不存在"}
	ERR_POST_ADD         = ResultCode{1303, "添加文章失敗"}
	ERR_POST_UPDATE      = ResultCode{1304, "更新文章失敗"}
	ERR_POST_DELETE      = ResultCode{1305, "刪除文章失敗"}

	ERR_POST_META_NOT_EXISTS = ResultCode{1311, "文章元數據不存在"}
	ERR_POST_META_ADD        = ResultCode{1312, "添加文章元數據失敗"}
	ERR_POST_META_UPDATE     = ResultCode{1313, "更新文章元數據失敗"}
	ERR_POST_META_DELETE     = ResultCode{1314, "刪除文章元數據失敗"}

	// role
	ERR_ROLE_NOT_EXISTS        = ResultCode{1401, "角色不存在"}
	ERR_ROLE_EXISTS            = ResultCode{1402, "角色已存在"}
	ERR_ROLE_REGEX             = ResultCode{1403, "角色的值必須為數字、字母或下劃線"}
	ERR_ROLE_ADD               = ResultCode{1404, "添加角色失敗"}
	ERR_ROLE_UPDATE            = ResultCode{1405, "更新角色失敗"}
	ERR_ROLE_DELETE            = ResultCode{1406, "刪除角色失敗"}
	ERR_ROLE_CAP_EXIST         = ResultCode{1411, "已擁有該權限"}
	ERR_ROLE_CAP_NOT_EXIST     = ResultCode{1412, "未擁有該權限"}
	ERR_ROLE_CAP_GRANT_FAILED  = ResultCode{1413, "賦予權限失敗"}
	ERR_ROLE_CAP_REMOVE_FAILED = ResultCode{1414, "撤銷權限失敗"}

	// cap
	ERR_CAP_NOT_EXISTS = ResultCode{1501, "權限不存在"}
	ERR_CAP_EXISTS     = ResultCode{1502, "權限已存在"}
	ERR_CAP_REGEX      = ResultCode{1503, "權限的值必須為數字、字母或下劃線"}
	ERR_CAP_ADD        = ResultCode{1504, "添加權限失敗"}
	ERR_CAP_UPDATE     = ResultCode{1505, "更新權限失敗"}
	ERR_CAP_DELETE     = ResultCode{1506, "刪除權限失敗"}
)
