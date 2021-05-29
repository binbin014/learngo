package routers

import (
	"github.com/gin-gonic/gin"
	"learngin/app/api/middleware"
	"learngin/library/logger/zaplog"
)

type Option func(engine *gin.Engine)

var options []Option

func Include(opts ...Option) {
	options = append(options, opts...)
}

func Init() *gin.Engine {
	r := gin.New()
	r.Use(zaplog.GinLogger(),zaplog.GinRecovery(true), middleware.Cors())
	for _, opt := range options {
		opt(r)
	}
	return r
}
