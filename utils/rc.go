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
	ERR_PARAMS = ResultCode{1, "Invalid params"}

	ERR_LOGIN_PASSWORD = ResultCode{10, "Invalid user_pass"}
	ERR_REGISTER       = ResultCode{11, "Register failed"}

	ERR_UNKNOWN = ResultCode{99, "Unknown error"}

	ERR_404 = ResultCode{404, "Not found"}

	// token
	ERR_TOKEN_GENERATE  = ResultCode{1001, "Couldn't generate a token"}
	ERR_TOKEN_MALFORMED = ResultCode{1002, "That's not even a token"}
	ERR_TOKEN_EXPIRED   = ResultCode{1003, "Token is either expired or not active yet"}
	ERR_TOKEN_UNKNOWN   = ResultCode{1004, "Couldn't handle this token"}

	// user
	ERR_USER_NOT_EXISTS       = ResultCode{1101, "User not exists"}
	ERR_USER_EXISTS           = ResultCode{1102, "User exists"}
	ERR_USER_INVALID_USERNAME = ResultCode{1103, "Invalid user_login"}
	ERR_USER_INVALID_PASSWORD = ResultCode{1104, "Invalid user_pass"}

	// upload
	ERR_UPLOAD_MKDIR           = ResultCode{1201, "Couldn't create uploads dir"}
	ERR_UPLOAD_FILE_NOT_EXISTS = ResultCode{1102, "File not exists"}
)
