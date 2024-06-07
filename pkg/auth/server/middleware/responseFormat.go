package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"worframe/share/constant"
)

type ResponseData struct {
	Code gin.ErrorType `json:"code"`
	Msg  string        `json:"msg"`
	Data interface{}   `json:"data"`
}

func ErrorResponse(c *gin.Context, code gin.ErrorType, message string) {
	c.JSON(c.Writer.Status(), ResponseData{
		Code: code,
		Msg:  message,
		Data: nil,
	})
}

func SuccessResponse(c *gin.Context, code gin.ErrorType, data interface{}) {
	if data == nil {
		c.JSON(c.Writer.Status(), ResponseData{
			Code: code,
			Msg:  "success",
			Data: nil,
		})
	} else {
		c.JSON(c.Writer.Status(), ResponseData{
			Code: code,
			Msg:  "success",
			Data: data,
		})
	}
}

func Response() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if c.Writer.Status() == 0 {
			c.Writer.WriteHeader(http.StatusOK)
		}
		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			ErrorResponse(c, err.Type, err.Error())
			return
		}
		if c.Writer.Status() >= http.StatusOK && c.Writer.Status() < http.StatusMultipleChoices {
			data, _ := c.Get("response_data")
			SuccessResponse(c, constant.SUCCESS, data)
			return
		}
	}
}
