package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"learngin/library/config"
)
var (
	Viper *viper.Viper
 	DB *gorm.DB
	Config *config.Config
	Logger *zap.Logger
)
