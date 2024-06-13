package unit

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"worframe/pkg/auth/utils"
)

func TestAuthUtils(t *testing.T) {
	t.Run("password", testPasswordUtils)
}
func testPasswordUtils(t *testing.T) {
	var (
		testPWD    = "123456"
		testSalt   = utils.SaltSpawn()
		encryptPWD string
	)
	assert.NotEmpty(t, utils.SaltSpawn())
	encryptPWD = utils.EncryptPassword(testPWD, testSalt)
	assert.NotEmpty(t, encryptPWD)
	testCompare := utils.ComparePassword(testPWD, testSalt, encryptPWD)
	assert.True(t, testCompare)
}
