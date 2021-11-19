package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/middleware"
)

// Routers 初始化总路由
func Routers() *gin.Engine {
	var Router = gin.Default()
	//获取路由组实例
	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
		// 系统-用户登录
	}
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth())
	{

	}
	magic.Logger.Info("router register success")
	return Router
}
