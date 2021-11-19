package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/middleware"
)

// Routers 初始化总路由
func Routers() *gin.Engine {
	var Router = gin.Default()

	// 公共路由组
	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
		// 系统-用户登录
	}

	// 私有路由组
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth())
	{
		// 系统-查询所有菜单
	}

	magic.Logger.Info("router register success")
	return Router
}
