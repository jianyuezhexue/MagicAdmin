package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/router"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func main() {
	// 初始化所有路由
	router := router.Routers()

	// 启动服务
	address := fmt.Sprintf(":%d", magic.Config.System.Addr)
	s := initServer(address, router)
	time.Sleep(10 * time.Microsecond)
	magic.Logger.Info("server run success on ", zap.String("address", address))
	magic.Logger.Error(s.ListenAndServe().Error())
}
