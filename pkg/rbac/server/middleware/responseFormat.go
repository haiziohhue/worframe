package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(code, ResponseData{
		Code: code,
		Msg:  message,
		Data: nil,
	})
}

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	if data == nil {
		c.JSON(code, ResponseData{
			Code: code,
			Msg:  "success",
			Data: nil,
		})
	} else {
		c.JSON(code, ResponseData{
			Code: code,
			Msg:  "success",
			Data: data,
		})
	}
}

func Response() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		if c.Writer.Status() == 0 {
			c.Writer.WriteHeader(http.StatusOK)
		}
		if c.Writer.Status() >= http.StatusOK && c.Writer.Status() < http.StatusMultipleChoices {
			data, _ := c.Get("response_data")
			SuccessResponse(c, c.Writer.Status(), data)
			return
		}
	}
}
