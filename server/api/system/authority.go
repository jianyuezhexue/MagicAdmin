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
		magic.HttpFail(c, http.StatusBadGateway, err.Error(), res)
		return
	}

	magic.Success(c, "获取角色列表成功", res)
}

// 创建角色
func (a *AuthorityCtr) Create(c *gin.Context) {
	// 接收参数
	var param system.Authority
	err := c.ShouldBind(&param)
	if err != nil {
		magic.HttpFail(c, http.StatusBadRequest, err.Error(), param)
		return
	}

	// 逻辑处理
	res, err := serviceSystem.AuthorityApp.Create(param)
	if err != nil {
		magic.HttpFail(c, http.StatusBadGateway, err.Error(), res)
		return
	}

	magic.Success(c, "创建角色成功", res)
}

// 更新角色
func (a AuthorityCtr) Update(c *gin.Context) {
	// 接收参数
	var param system.Authority
	err := c.ShouldBind(&param)
	if err != nil {
		magic.HttpFail(c, http.StatusBadRequest, err.Error(), param)
		return
	}

	// 逻辑处理
	res, err := serviceSystem.AuthorityApp.Update(param)
	if err != nil {
		magic.HttpFail(c, http.StatusBadGateway, err.Error(), res)
		return
	}

	magic.Success(c, "更新角色成功", res)
}

// 删除角色
func (a AuthorityCtr) Delete(c *gin.Context) {
	var id model.GetById
	err := c.ShouldBindUri(&id)
	if err != nil {
		magic.HttpFail(c, http.StatusBadRequest, err.Error(), id)
		return
	}

	// 逻辑处理
	res, err := serviceSystem.AuthorityApp.Delete(id)
	if err != nil {
		magic.HttpFail(c, http.StatusBadGateway, err.Error(), res)
		return
	}

	magic.Success(c, "删除角色成功", res)
}

// 角色设置菜单权限
func (a *AuthorityCtr) SetMenuAuth(c *gin.Context) {
	// 接收参数
	var form system.SetAuth
	err := c.ShouldBind(&form)
	if err != nil {
		magic.HttpFail(c, http.StatusBadRequest, err.Error(), form)
		return
	}

	// 逻辑处理
	res, err := serviceSystem.AuthorityApp.SetMenuAuth(form)
	if err != nil {
		magic.HttpFail(c, http.StatusBadGateway, err.Error(), res)
		return
	}

	magic.Success(c, "角色设置菜单权限成功", res)
}

// 角色设置API权限
func (a *AuthorityCtr) SetApiAuth(c *gin.Context) {
	// 接收参数
	var form system.SetAuth
	err := c.ShouldBind(&form)
	if err != nil {
		magic.HttpFail(c, http.StatusBadRequest, err.Error(), form)
		return
	}

	// 逻辑处理
	res := serviceSystem.AuthorityApp.SetApiAuth(form)
	if res.Code != 0 {
		magic.HttpFail(c, res.Code, res.Msg, res)
		return
	}

	magic.Success(c, res.Msg, res)
}

// 角色设置拓展权限
func (a *AuthorityCtr) SetExtAuth(c *gin.Context) {
	// 接收参数
	var form system.SetAuth
	err := c.ShouldBind(&form)
	if err != nil {
		magic.HttpFail(c, http.StatusBadRequest, err.Error(), form)
		return
	}

	// 逻辑处理
	res, err := serviceSystem.AuthorityApp.SetExtAuth(form)
	if err != nil {
		magic.HttpFail(c, http.StatusBadGateway, err.Error(), res)
		return
	}

	magic.Success(c, "角色设置API权限成功", res)
}
