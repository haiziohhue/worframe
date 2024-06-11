package unit

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	authCore "worframe/pkg/auth/core"
	"worframe/share/core"
)

var testApp *authCore.AuthCore

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
	shareApp := core.
		NewApp("test").InitZap().InitDb().InitRedis()

	if shareApp.Error != nil {
		panic(shareApp.Error)
	}
	testApp = authCore.
		NewAuthCore(shareApp).InitEngine()

	if testApp.Error != nil {
		panic(testApp.Error)
	}

	m.Run()
}
func TestExample(t *testing.T) {
	fmt.Println("Example test")
}
