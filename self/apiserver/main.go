package main

import (
  "github.com/jw803/webapp/config"
	"github.com/jw803/webapp/middleware"
	"github.com/jw803/webapp/handler"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	load()

	gin.SetMode(config.Val.Mode)
	r := gin.Default() // 預設會幫你掛上兩個middleware 1.access log 2.panic自動recovery
	r.Use(middleware.CROSS())

	r.GET("/ping", handler.Ping)

	r.Run(":" + config.Val.Port)

	log.Infof("serve port: %v \n", config.Val.Port)
}

func load() {
	config.Init()
}