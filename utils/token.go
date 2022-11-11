package utils

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

var signKey = []byte("have_a_good_day")

// 生成 token
func GenerateToken(payloads ...map[string]interface{}) (string, error) {
	data := jwt.MapClaims{
		// 签发者
		"iss": "eftech.site",
		// 签发时间
		"iat": time.Now().Unix(),
		// 过期时间： 默认30 天
		"exp": time.Now().Unix() + 30*24*60*60,
	}
	// 自定义参数
	for _, payload := range payloads {
		for k, v := range payload {
			data[k] = v
		}
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	return token.SignedString(signKey)
}

// 验证并解析 token
func ParseToken(tokenString string) (ResultCode, jwt.MapClaims) {
	token, err := parse(tokenString)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return ERR_TOKEN_MALFORMED, nil
		} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
			return ERR_TOKEN_EXPIRED, nil
		} else {
			ERR_TOKEN_UNKNOWN.Msg = ERR_TOKEN_UNKNOWN.Msg + ": " + err.Error()
			return ERR_TOKEN_UNKNOWN, nil
		}
	} else {
		if token.Valid {
			mc := token.Claims.(jwt.MapClaims)
			return SUCCESS, mc
		} else {
			return ERR_UNKNOWN, nil
		}
	}
}

// token 负载信息
func TokenPayloads(tokenString string) jwt.MapClaims {
	token, err := parse(tokenString)
	if err != nil {
		return nil
	}
	if token.Valid {
		return token.Claims.(jwt.MapClaims)
	}
	return nil
}

// 解析 token
func parse(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return signKey, nil
	})
}
