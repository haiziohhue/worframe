package utils

func EncryptPassword(username, password string, salt string) string {
	newStr := username + password + salt
	newPassword, err := EncryptString(newStr)
	if err != nil {
		return ""
	}
	return newPassword
}
func SaltSpawn() string {
	return GenerateSubId(6)
}
