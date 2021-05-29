package driver

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"learngin/library/global"
)

var MysqlDriver = new(mysqlDb)

type mysqlDb struct{}

func (m *mysqlDb) Initialize() {
	db, err := gorm.Open(mysql.Open(global.Viper.GetString("database.link")), &gorm.Config{})
	if err != nil {
		panic("database link err " + err.Error())
	}

	global.DB = db

}
