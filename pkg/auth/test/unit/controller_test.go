package unit

import (
	"encoding/json"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeptGetALL(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/dept?page_size=15", nil)
	router.ServeHTTP(w, req)
	var q res
	err := json.Unmarshal(w.Body.Bytes(), &q)
	if err != nil {
		panic(err)
	}
	t.Log(q)
	assert.Equal(t, 200, q.Code)
}
func TestDeptGetOne(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/dept/:id", nil)
	router.ServeHTTP(w, req)
	var q res
	err := json.Unmarshal(w.Body.Bytes(), &q)
	if err != nil {
		panic(err)
	}
	t.Log(q)
	assert.Equal(t, 200, q.Code)
}
func TestDeptPOST(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/dept", nil)
	router.ServeHTTP(w, req)
	var q res
	err := json.Unmarshal(w.Body.Bytes(), &q)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(q)
	assert.Equal(t, 201, q.Code)
}
func TestDeptPUT(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/dept/12", nil)
	router.ServeHTTP(w, req)
	var q res
	err := json.Unmarshal(w.Body.Bytes(), &q)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(q)
	assert.Equal(t, 200, q.Code)
}
func TestDeptDELETE(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/dept/13", nil)
	router.ServeHTTP(w, req)
	var q res
	err := json.Unmarshal(w.Body.Bytes(), &q)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(q)
	assert.Equal(t, 200, q.Code)
}
