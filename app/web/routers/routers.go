package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routers(e *gin.Engine)  {
	e.GET("/post", func(context *gin.Context) {
			context.JSON(http.StatusOK,gin.H{
				"message":"post",
			})
	})
	e.GET("/comment", func(context *gin.Context) {
			context.JSON(http.StatusOK,gin.H{
				"message":"comment",
			})
	})
}