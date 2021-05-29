package Controllers

import (
	"github.com/gin-gonic/gin"
	"learngin/app/api/services"
	"learngin/library/response"
)

var User = new(user)

type user struct{}

// Register 用户注册
func (u *user) Register(c *gin.Context) {
	if err := services.User.Register(c); err != nil {
		response.Json(c, 1, err.Error())
		return
	}
	response.Json(c, 0, "注册成功")
}

// Profile 用户信息
func (u *user) Profile(c *gin.Context) {
	result, err := services.User.FindOne(c)

	if err != nil {
		response.Json(c, 1, err.Error())
		return
	}
	response.Json(c, 0, "获取用户信息成功", gin.H{
		"username": result.Username,
		"roles":    []string{"admin"},
		"avatar":   "https://ss1.bdstatic.com/70cFuXSh_Q1YnxGkpoWK1HF6hhy/it/u=3355464299,584008140&fm=26&gp=0.jpg",
		"name":     result.Username,
	})
}

