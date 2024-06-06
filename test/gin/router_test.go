package gin

type header struct {
	Key   string
	Value string
}

//
//func performRequest(r http.Handler, method, path string, headers ...header) *httptest.ResponseRecorder {
//	req := httptest.NewRequest(method, path, nil)
//	for _, h := range headers {
//		req.Header.Add(h.Key, h.Value)
//	}
//	w := httptest.NewRecorder()
//	r.ServeHTTP(w, req)
//	return w
//}
//func TestRouterMethod(t *testing.T) {
//	r := gin.New()
//	r.GET("/book/:id", func(c *gin.Context) {
//		var query struct {
//			Author bool `form:"author"`
//			Id     int  `uri:"id" binding:"required"`
//			Page   int  `form:"page"`
//		}
//		//_ = c.ShouldBindQuery(&query)
//		_ = c.ShouldBind(&query)
//		_ = c.ShouldBindUri(&query)
//		c.JSON(http.StatusOK, query)
//	})
//
//	w := performRequest(r, http.MethodGet, "/book/12?author=true&page=10")
//
//	assert.Equal(t, http.StatusOK, w.Code)
//}
