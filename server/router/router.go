package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jianyuezhexue/MagicAdmin/api/system"
	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/middle"
)

// Routers 初始化总路由
func Routers() *gin.Engine {
	var Router = gin.Default()

	// 公共路由组
	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "是心动啊...")
		})
		PublicGroup.GET("/configInfo", system.ConfigInfo)
		// 系统-用户登录
		PublicGroup.POST("/user/register", system.Register)
		PublicGroup.POST("/user/login", system.Login)
	}

	// 私有路由组
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middle.JWTAuth())
	{
		// 系统-菜单
		PrivateGroup.GET("/myMenu", system.MyMenu)    // 权限菜单
		PrivateGroup.GET("/menus", system.Menus)      // 树型菜单
		PrivateGroup.POST("/menu", system.CreateMenu) // 新增菜单
		PrivateGroup.PUT("/menu", system.Menus)       // 编辑菜单
		PrivateGroup.DELETE("/menu", system.Menus)    // 删除菜单
	}

	magic.Logger.Info("router register success")
	return Router
}
