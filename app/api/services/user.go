package services

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"learngin/app/api/models"
	"learngin/library/global"
	"log"
)

var User = new(user)
var UserValues models.User

type user struct{}

// Register 用户注册服务
func (u *user) Register(c *gin.Context) error {
	if err := c.ShouldBind(&UserValues); err != nil {
		return err
	}
	password := fmt.Sprintf("%x", md5.New().Sum([]byte(UserValues.Password)))
	Uuid, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}
	users := models.User{
		Username: UserValues.Username,
		Password: password,
		Uuid:     Uuid.String(),
	}
	result := global.DB.Create(&users)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}
	return nil
}

// FindOne 查询用户信息
func (u *user) FindOne(c *gin.Context) (*models.User, error) {
	userid, ok := c.Get("uuid")
	if !ok {
		return nil, errors.New("无效用户")
	}
	result := global.DB.Where("uuid=?", userid).First(&UserValues)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	usersInfo := &models.User{
		Username: UserValues.Username,
	}
	return usersInfo, nil
}
