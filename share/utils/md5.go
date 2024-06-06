package utils

import (
	"crypto/md5"
	"fmt"
)

func EncryptString(str string) (string, error) {
	return EncryptBytes([]byte(str))
}
func EncryptBytes(bytes []byte) (string, error) {
	h := md5.New()
	if _, err := h.Write(bytes); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
