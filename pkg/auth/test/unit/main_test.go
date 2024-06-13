package unit

import (
	"testing"
	authCore "worframe/pkg/auth/core"
	"worframe/pkg/auth/core/iface"
	"worframe/share/core"
)

var testApp iface.ICore

func TestMain(m *testing.M) {
	shareApp := core.
		NewApp("test").InitZap().InitDb().InitRedis()

	if shareApp.GetErr() != nil {
		panic(shareApp.GetErr())
	}
	testApp = authCore.
		NewAuthCore(shareApp).InitEngine()

	if testApp.GetErr() != nil {
		panic(testApp.GetErr())
	}

	m.Run()
}
