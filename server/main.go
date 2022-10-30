package main

import (
	"fmt"
	"time"

	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/router"
	"github.com/jianyuezhexue/MagicAdmin/schedule"
)

func main() {
	// 初始化所有路由
	router := router.Routers()

	// 启动服务
	address := fmt.Sprintf(":%d", magic.Config.System.Addr)
	s := magic.InitServer(address, router)

	// 启动定时任务
	schedule.InitCron()

	// 监听错误
	time.Sleep(10 * time.Microsecond)
	magic.Logger.Error(s.ListenAndServe().Error())
}
