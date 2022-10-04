package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model"
	"github.com/jianyuezhexue/MagicAdmin/model/system"
	serviceSystem "github.com/jianyuezhexue/MagicAdmin/service/system"
)

type AuthorityCtr struct{}

var AuthorityAPI = new(AuthorityCtr)

// 获取角色列表
func (a *AuthorityCtr) TreeList(c *gin.Context) {
	// 接收参数

	// 逻辑处理
	res, err := serviceSystem.AuthorityApp.TreeList()
	if err != nil {
		magic.Fail(c, http.StatusBadGateway, err.Error(), res)
		return
	}

	magic.Success(c, "创建字典目录成功", res)
}

// 创建角色
func (a *AuthorityCtr) Create(c *gin.Context) {
	// 接收参数
	var param system.Authority
	err := c.ShouldBind(&param)
	if err != nil {
		magic.Fail(c, http.StatusBadRequest, err.Error(), param)
		return
	}

	// 逻辑处理
	res, err := serviceSystem.AuthorityApp.Create(param)
	if err != nil {
		magic.Fail(c, http.StatusBadGateway, err.Error(), res)
		return
	}

	magic.Success(c, "创建字典目录成功", res)
}

// 修改角色
func (a AuthorityCtr) Update(c *gin.Context) {
	// 接收参数
	var param system.Authority
	err := c.ShouldBind(&param)
	if err != nil {
		magic.Fail(c, http.StatusBadRequest, err.Error(), param)
		return
	}

	// 逻辑处理
	res, err := serviceSystem.AuthorityApp.Update(param)
	if err != nil {
		magic.Fail(c, http.StatusBadGateway, err.Error(), res)
		return
	}

	magic.Success(c, "创建字典目录成功", res)
}

// 删除角色
func (a AuthorityCtr) Delete(c *gin.Context) {
	var id model.GetById
	err := c.ShouldBindUri(&id)
	if err != nil {
		magic.Fail(c, http.StatusBadRequest, err.Error(), id)
		return
	}

	// 逻辑处理
	res, err := serviceSystem.AuthorityApp.Delete(id)
	if err != nil {
		magic.Fail(c, http.StatusBadGateway, err.Error(), res)
		return
	}

	magic.Success(c, "创建字典目录成功", res)
}