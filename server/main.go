package main

import (
	"fmt"
	"time"

	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/router"
)

func main() {
	// 初始化所有路由
	router := router.Routers()

	// 启动服务
	address := fmt.Sprintf(":%d", magic.Config.System.Addr)
	s := magic.InitServer(address, router)
	time.Sleep(10 * time.Microsecond)
	magic.Logger.Error(s.ListenAndServe().Error())
}
