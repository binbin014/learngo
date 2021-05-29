package jwt

import "C"
import (
	"crypto/md5"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"learngin/app/api/models"
	"learngin/library/global"
	"learngin/library/response"
	"time"
)

var Auth = new(auth)

type auth struct{}

func (a *auth) Initialize() *jwt.GinJWTMiddleware {
	authMiddleWare, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:                 global.Viper.GetString("jwt.signing-key"),
		SigningAlgorithm:      "",
		Key:                   []byte(global.Viper.GetString("jwt.signing-key")),
		Timeout:               time.Hour * 2,
		MaxRefresh:            time.Hour * 2,
		Authenticator:         Authenticator,
		Authorizator:          nil,
		PayloadFunc:           PayloadFunc,
		Unauthorized:          nil,
		LoginResponse:         LoginResponse,
		LogoutResponse:        LogoutResponse,
		RefreshResponse:       nil,
		IdentityHandler:       nil,
		IdentityKey:           "",
		TokenLookup:           "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:         "Bearer",
		TimeFunc:              time.Now,
		HTTPStatusMessageFunc: nil,
		PrivKeyFile:           "",
		PubKeyFile:            "",
		SendCookie:            false,
		CookieMaxAge:          0,
		SecureCookie:          false,
		CookieHTTPOnly:        false,
		CookieDomain:          "",
		SendAuthorization:     false,
		DisabledAbort:         false,
		CookieName:            "",
		CookieSameSite:        0,
	})
	if err != nil {
		panic("JWT Error:" + err.Error())
	}
	return authMiddleWare
}

// Authenticator 授权验证的操作
func Authenticator(c *gin.Context) (interface{}, error) {
	var userValues models.User
	if err := c.ShouldBind(&userValues); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	username := userValues.Username
	password := fmt.Sprintf("%x", md5.New().Sum([]byte(userValues.Password)))
	result := global.DB.Where("username=? and password=?", username, password).Find(&userValues)
	if result.Error != nil {
		panic("sql select fail ,err:" + result.Error.Error())
	}
	if result.RowsAffected > 0 {
		c.Set("userinfo", userValues)
		return gin.H{"username": userValues.Username, "password": userValues.Password, "uuid": userValues.Uuid}, nil
	}
	return gin.H{}, nil
}

// PayloadFunc jwt的数据处理
func PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	params := data.(gin.H)
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

// LoginResponse 登录成功返回数据
func LoginResponse(c *gin.Context, _ int, token string, expire time.Time) {
	claims, _ := c.Get("userinfo")
	if claims != nil {
		data, ok := claims.(models.User)
		if ok {
			response.Json(c, 0, "登录成功", gin.H{
				"token":    token,
				"username": data.Username,
				"expire":   expire.Unix(),
			})
			return
		}
		response.Json(c, 1, "登录失败,用户不存在")
		return
	}
	response.Json(c, 1, "登录失败,账号密码不正确")
}

func LogoutResponse(c *gin.Context, _ int ) {
	response.Json(c, 0, "退出成功")
}
