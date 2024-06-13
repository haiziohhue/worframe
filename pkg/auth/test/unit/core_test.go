package unit

import (
	"github.com/go-playground/assert/v2"
	"testing"
	authCore "worframe/pkg/auth/core"
	"worframe/share/core"
)

func TestCore(t *testing.T) {
	shareApp := core.
		NewApp("test").InitZap().InitDb().InitRedis()
	app := authCore.NewAuthCore(shareApp)
	assert.Equal(t, app.GetErr(), nil)
	app = app.InitAuthConf()
	assert.Equal(t, app.GetErr(), nil)
	app = app.InitEngine()
	assert.Equal(t, app.GetErr(), nil)

	assert.NotEqual(t, app.GetEngine(), nil)
	assert.NotEqual(t, app.GetRawCore(), nil)
	assert.NotEqual(t, app.GetDB(), nil)
	assert.NotEqual(t, app.GetRedis(), nil)
	assert.NotEqual(t, app.GetLog(), nil)
	assert.NotEqual(t, app.GetSLog(), nil)
	assert.NotEqual(t, app.GetConf(), nil)
	assert.NotEqual(t, app.GetAuthConf(), nil)

}
