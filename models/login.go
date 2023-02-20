package models

import "hello_gin_api/utils"

type LoginRes struct {
	User
	AccessToken string      `json:"access_token"`
	Payloads    interface{} `json:"payloads"`
}

func Login(u *User) utils.Result {
	var result = utils.Result{}
	user := GetUserByLogin(u.UserLogin)
	if user.Valid() {
		b := utils.EqualsUserPass(u.UserPass, user.UserPass)
		if b {
			data := make(map[string]interface{})
			data["data"] = user

			payloads := make(map[string]interface{})
			payloads["uid"] = user.ID
			tokenString, _ := utils.GenerateToken(payloads)
			data["access_token"] = tokenString
			rc, extras := utils.ParseToken(tokenString)
			if rc.Success() {
				data["token_payloads"] = extras
			}

			result.ResultCode = utils.SUCCESS
			result.Data = data
		} else {
			result.ResultCode = utils.ERR_LOGIN_PASSWORD
		}
	} else {
		result.ResultCode = utils.ERR_USER_NOT_EXISTS
	}
	return result
}
