package routers

import (
	"github.com/gin-gonic/gin"
	"learngin/app/api/Controllers"
	"learngin/app/api/middleware"
	"learngin/app/api/middleware/jwt"
)

func Routers(e *gin.Engine) {
	e.POST("/api/user/login",jwt.Auth.Initialize().LoginHandler)
	e.POST("/api/user/register",Controllers.User.Register)
	r:=e.Group("/api",middleware.Auth())
	{
		r.POST("/user/profile",Controllers.User.Profile)
		r.POST("/user/loginout",jwt.Auth.Initialize().LogoutHandler)
	}

}
