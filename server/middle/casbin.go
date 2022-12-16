package middle

import (
	"github.com/gin-gonic/gin"
	"github.com/jianyuezhexue/MagicAdmin/magic"
	serviceSystem "github.com/jianyuezhexue/MagicAdmin/service/system"
)

// API鉴权中间件
func Casbin() gin.HandlerFunc {
	return func(c *gin.Context) {
		e := serviceSystem.CasbinApp.Casbin()
		userInfo := magic.TokenInfo(c)
		// 角色
		sub := userInfo.AuthorityId
		// 方法
		obj := c.Request.Method
		// 路由
		act := c.Request.URL.Path

		check, _ := e.Enforce(sub, obj, act)
		magic.Print(check)
		// if !check {
		// 	magic.HttpFail(c, 1001, "权限不足", "权限不足")
		// 	c.Abort()
		// 	return
		// }

		c.Next()
	}
}
