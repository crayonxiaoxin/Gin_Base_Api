package utils

import "golang.org/x/crypto/bcrypt"

// 密码加密
func HashedUserPass(user_pass string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(user_pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// 密码比较
func EqualsUserPass(user_pass string, hashedUserPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedUserPass), []byte(user_pass))
	return err == nil
}
