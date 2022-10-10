package utils

import "golang.org/x/crypto/bcrypt"

/*
   @Auth: menah3m
   @Desc: bCrypt 加密
*/

//EncodeBCrypt 加密
func EncodeBCrypt(value string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

//VerifyPassword 验证密码
func VerifyPassword(pwd, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}
