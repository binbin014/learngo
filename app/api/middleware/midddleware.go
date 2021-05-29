package middleware

import (
	ginJwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"learngin/app/api/middleware/jwt"
	"learngin/library/response"
	"net/http"
)

// Auth jwt中间件
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		_jwt, err := jwt.Auth.Initialize().ParseToken(c)
		if err!=nil{
			response.JsonAbort(c,http.StatusBadRequest,err.Error())
		}
		if _jwt != nil {
			var claims = ginJwt.ExtractClaimsFromToken(_jwt)
			c.Set("claims", claims)
			c.Set("uuid", claims["uuid"])
			c.Next()
			return
		}
		response.JsonAbort(c,http.StatusBadRequest,"access_token_invalid")
	}
}

// Cors 跨域的中间件
func Cors() gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AddAllowMethods("OPTIONS")
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
	corsConfig.AllowAllOrigins = true
	return cors.New(corsConfig)
}
