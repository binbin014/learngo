package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Json(c *gin.Context,code int,message string,data ...interface{}){
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	c.AbortWithStatusJSON(http.StatusOK,Response{
		Code:    code,
		Message: message,
		Data:    responseData,
	})
	return
}
func JsonAbort(c *gin.Context,code int,message string,data ...interface{}){
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	c.AbortWithStatusJSON(code,Response{
		Code:    code,
		Message: message,
		Data:    responseData,
	})
	return
}

