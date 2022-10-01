package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jianyuezhexue/MagicAdmin/magic"
	serviceSystem "github.com/jianyuezhexue/MagicAdmin/service/system"
)

type Authority struct{}

var AuthorityAPI = new(Authority)

// 获取角色列表
func (a *Authority) TreeList(c *gin.Context) {
	// 接收参数

	// 逻辑处理
	res, err := serviceSystem.AuthorityApp.TreeList()
	if err != nil {
		magic.Fail(c, http.StatusBadGateway, err.Error(), res)
		return
	}

	magic.Success(c, "创建字典目录成功", res)
}
