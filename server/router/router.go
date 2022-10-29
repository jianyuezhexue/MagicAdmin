package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jianyuezhexue/MagicAdmin/api/system"
	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/middle"
)

// Routers 初始化总路由
func Routers() *gin.Engine {
	gin.SetMode(gin.DebugMode) // debug模式
	var Router = gin.Default()

	Router.Use(middle.CostTime()).Use(middle.Record())

	// 静态资源地址
	Router.Static("/form-generator", "./resource/page")

	// 公共路由组
	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "是心动啊...")
		})
		// 系统-获取验证码
		PublicGroup.POST("/user/captcha", system.CommonApi.Captcha)
		// 系统-用户登录
		PublicGroup.POST("/user/login", system.UserApi.Login)
	}

	// 私有路由组
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middle.JWTAuth()).Use(middle.Casbin())
	{
		// 系统-用户
		PrivateGroup.GET("/user/info", system.UserApi.UserInfo)           // 用户信息
		PrivateGroup.GET("/users", system.UserApi.List)                   // 用户列表
		PrivateGroup.POST("/user/register", system.UserApi.Register)      // 新建用户
		PrivateGroup.PUT("/user", system.UserApi.Update)                  // 编辑用户
		PrivateGroup.PATCH("/user/auth", system.UserApi.SetUserAuth)      // 设置用户角色
		PrivateGroup.PATCH("/user/status", system.UserApi.SetUserStatus)  // 设置用户状态
		PrivateGroup.PATCH("/user/reSetPwd/:id", system.UserApi.ReSetPwd) // 重置用户密码
		PrivateGroup.PATCH("/user/switchAuth", system.UserApi.SwitchAuth) // 切换用户角色
		PrivateGroup.DELETE("/user/:id", system.UserApi.Delete)           // 删除菜单

		// 系统-菜单
		PrivateGroup.GET("/myMenu", system.MenuApi.MyMenu)          // 我的菜单
		PrivateGroup.GET("/menus", system.MenuApi.Menus)            // 树型菜单
		PrivateGroup.GET("/menu/:id", system.MenuApi.FindMenu)      // 查询菜单
		PrivateGroup.PUT("/menu", system.MenuApi.UpdateMenu)        // 编辑菜单
		PrivateGroup.POST("/menu", system.MenuApi.CreateMenu)       // 新增菜单
		PrivateGroup.DELETE("/menu/:id", system.MenuApi.DeleteMenu) // 删除菜单
		PrivateGroup.GET("/menuTree", system.MenuApi.MenuTree)      // 权限菜单

		// 系统-字典目录
		PrivateGroup.POST("/dictionary", system.DictionaryApi.Create)       // 新增字典目录
		PrivateGroup.GET("/dictionary", system.DictionaryApi.List)          // 分页字典目录
		PrivateGroup.GET("/dictionary/:id", system.DictionaryApi.Item)      // 查询字典目录
		PrivateGroup.PUT("/dictionary", system.DictionaryApi.Update)        // 更新字典目录
		PrivateGroup.DELETE("/dictionary/:id", system.DictionaryApi.Delete) // 删除字典目录

		// 系统-字典详情
		PrivateGroup.POST("/dictionaryDetail", system.DictionaryDetailApi.Create)       // 新增字典目录
		PrivateGroup.GET("/dictionaryDetail", system.DictionaryDetailApi.List)          // 分页字典目录
		PrivateGroup.GET("/dictionaryDetail/:id", system.DictionaryDetailApi.Item)      // 查询字典目录
		PrivateGroup.PUT("/dictionaryDetail", system.DictionaryDetailApi.Update)        // 更新字典目录
		PrivateGroup.DELETE("/dictionaryDetail/:id", system.DictionaryDetailApi.Delete) // 删除字典目录

		// 系统-角色
		PrivateGroup.GET("/authorityTree", system.AuthorityAPI.TreeList)         // 角色树形列表
		PrivateGroup.POST("/authority", system.AuthorityAPI.Create)              // 创建角色
		PrivateGroup.PUT("/authority", system.AuthorityAPI.Update)               // 更新角色
		PrivateGroup.DELETE("/authority/:id", system.AuthorityAPI.Delete)        // 删除角色
		PrivateGroup.PATCH("/authority/menu", system.AuthorityAPI.SetMenuAuth)   // 设置角色菜单权限｜待优化
		PrivateGroup.PATCH("/authority/api", system.AuthorityAPI.SetApiAuth)     // 设置角色API权限｜待优化
		PrivateGroup.PATCH("/authority/extAuth", system.AuthorityAPI.SetExtAuth) // 设置角色拓展权限｜待优化

		// 系统-API
		PrivateGroup.GET("/api", system.ApiAPI.List)          // 分页角色列表
		PrivateGroup.DELETE("/api/:id", system.ApiAPI.Delete) // 删除API

		// 系统-拓展权限
		PrivateGroup.DELETE("/extAuth/:id", system.ExtAuthAPI.Delete) // 删除拓展权限

		// 系统-公共接口
		PrivateGroup.GET("/common/fileList", system.CommonApi.MediaList) // 查询媒体列表
	}

	magic.Logger.Info("router register success")
	return Router
}
