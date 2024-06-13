package utils

import (
	"golang.org/x/crypto/bcrypt"
	shareUtils "worframe/share/utils"
)

func EncryptPassword(password, salt string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func SaltSpawn() string {
	return shareUtils.GenerateSubId(6)
}
func ComparePassword(inputPwd, salt, target string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(target), []byte(inputPwd+salt))
	if err != nil {
		return false
	}
	return true
}
