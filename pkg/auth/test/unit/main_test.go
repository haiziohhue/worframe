package unit

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
	"worframe/pkg/auth/config"
	initialize2 "worframe/pkg/auth/initialize"
	"worframe/pkg/auth/migrate"
	"worframe/pkg/auth/server"
	"worframe/share/core"
	"worframe/share/initialize"
)

var r *gin.Engine

type header struct {
	Key   string
	Value string
}

func performRequest(r http.Handler, method, path string, body []byte, headers ...header) *httptest.ResponseRecorder {
	data := bytes.NewReader(body)
	req := httptest.NewRequest(method, path, data)
	for _, h := range headers {
		req.Header.Add(h.Key, h.Value)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestMain(m *testing.M) {
	core.Cfg = initialize.InitConfig("test")
	config.AuthCfg = initialize2.InitAuthConfig("test")
	core.Logger = initialize.InitZap(core.Cfg)
	core.DB = initialize.InitGorm(core.Cfg)
	core.Logger.Debug("hello world")
	mi := migrate.NewDBMigrate(core.DB)
	err := mi.TestEnvInit()
	if err != nil {
		panic(err)
	}

	core.Redis = initialize.InitRedis(core.Cfg)
	r = gin.New()
	server.AuthInitServer(r)
	m.Run()
}
func TestExample(t *testing.T) {
	fmt.Println("Example test")
}
