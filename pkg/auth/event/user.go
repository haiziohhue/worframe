package event

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"worframe/pkg/auth/utils"
	"worframe/share/model"
)

func ExportUserRole(user model.SysUser) []string {
	var result []string
	result = append(result, user.Dept.DeptName)
	result = append(result, "__user_"+user.UserName)
	for _, item := range user.Role {
		result = append(result, item.RoleName)
	}
	return result
}
func ComparePassword(InputPWD string, user *model.SysUser) bool {
	return utils.ComparePassword(InputPWD, user.Salt, user.Password)
}
func CreateNewPassword(password string) (string, string, error) {
	salt := utils.SaltSpawn()
	encryptPwd := utils.EncryptPassword(password, salt)
	if encryptPwd == "" {
		return "", "", fmt.Errorf("failed to encrypt password")
	}
	return encryptPwd, salt, nil
}
func ValidateUsername(username string) (string, bool) {
	v := validator.New()

	err := v.Var(username, "email")
	if err == nil {
		return "email", true
	}

	err = v.Var(username, "phone")
	if err == nil {
		return "phone", true
	}

	err = v.Var(username, "/^[a-z0-9_-]{3,16}$/")
	if err == nil {
		return "default", true
	}

	return "", false
}
