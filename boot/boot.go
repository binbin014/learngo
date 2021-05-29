package boot

import (
	"learngin/library/driver"
	"learngin/library/logger/zaplog"
)

func Initialize(path ...string) {
	Viper.Initialize(path...)
	driver.MysqlDriver.Initialize()
	zaplog.InitLogger()
}
