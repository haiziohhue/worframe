package utils

import shareUtils "worframe/share/utils"

func EncryptPassword(username, password string, salt string) string {
	newStr := username + password + salt
	newPassword, err := shareUtils.EncryptString(newStr)
	if err != nil {
		return ""
	}
	return newPassword
}
func SaltSpawn() string {
	return shareUtils.GenerateSubId(6)
}
