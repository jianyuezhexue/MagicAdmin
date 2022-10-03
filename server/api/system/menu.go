package system

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model"
	"github.com/jianyuezhexue/MagicAdmin/model/system"
	serviceSystem "github.com/jianyuezhexue/MagicAdmin/service/system"
)

type MenuCtr struct{}

var MenuApi = new(MenuCtr)

// 树形菜单
func (m *MenuCtr) MenuTree(c *gin.Context) {
	// 逻辑处理
	res, err := serviceSystem.MenuApp.MenuTree()
	if err != nil {
		magic.Fail(c, http.StatusBadGateway, err.Error(), res)
		return
	}

	// 结果返回
	magic.Success(c, "获取树形菜单成功", res)
}

// 我的菜单
func (m *MenuCtr) MyMenu(c *gin.Context) {
	// 参数获取
	userInfo := magic.TokenInfo(c)
	authorityId, _ := strconv.Atoi(userInfo.AuthorityId)

	// 逻辑处理
	res, err := serviceSystem.MenuApp.MyMenu(authorityId)
	if err != nil {
		magic.Fail(c, http.StatusBadGateway, err.Error(), res)
		return
	}
	// 结果返回
	magic.Success(c, "获取我的菜单成功", res)
}

// 分页获取基础menu列表
func (m *MenuCtr) Menus(c *gin.Context) {
	// 参数获取
	var pageInfo model.PageInfo
	err := c.ShouldBind(&pageInfo)
	if err != nil {
		magic.Fail(c, http.StatusBadRequest, err.Error(), pageInfo)
		return
	}

	// 逻辑处理
	res, err := serviceSystem.MenuApp.Menus(pageInfo)
	if err != nil {
		magic.Fail(c, http.StatusBadGateway, err.Error(), res)
		return
	}

	// 结果返回
	magic.Success(c, "获取菜单列表成功", res)
}

// 根据id获取菜单
func (m *MenuCtr) FindMenu(c *gin.Context) {
	// 参数校验
	var id model.GetById
	err := c.ShouldBindUri(&id)
	if err != nil {
		magic.Fail(c, http.StatusBadRequest, err.Error(), id)
		return
	}

	// 逻辑处理
	res, err := serviceSystem.MenuApp.FindMenu(id)
	if err != nil {
		magic.Fail(c, 1202, err.Error(), res)
		return
	}

	// 返回结果
	magic.Success(c, "根据id查询菜单成功", res)
}

// 更新菜单
func (m *MenuCtr) UpdateMenu(c *gin.Context) {
	// 参数校验
	var menu system.Menu
	err := c.ShouldBind(&menu)
	if err != nil {
		magic.Fail(c, 1204, err.Error(), menu)
		return
	}

	// 校验路由自己名字不能重复

	// 逻辑处理
	res, err := serviceSystem.MenuApp.UpdateMenu(menu)
	if err != nil {
		magic.Fail(c, 1206, err.Error(), res)
		return
	}

	// 返回结果
	magic.Success(c, "更新菜单成功", res)
}

// 创建菜单
func (m *MenuCtr) CreateMenu(c *gin.Context) {
	// 参数校验
	var param system.Menu
	err := c.ShouldBind(&param)
	if err != nil {
		magic.Fail(c, http.StatusBadRequest, err.Error(), param)
		return
	}

	// 校验路由自己名字不能重复

	// 逻辑处理
	res, err := serviceSystem.MenuApp.CreateMenu(param)
	if err != nil {
		magic.Fail(c, http.StatusBadGateway, err.Error(), res)
		return
	}
	// 结果返回
	magic.Success(c, "创建菜单成功", res)
}

// 删除菜单
func (m *MenuCtr) DeleteMenu(c *gin.Context) {
	// 参数校验
	var id model.GetById
	err := c.ShouldBindUri(&id)
	if err != nil {
		magic.Fail(c, http.StatusBadRequest, err.Error(), id)
		return
	}

	// 逻辑处理
	err = serviceSystem.MenuApp.DeleteMenu(id)
	if err != nil {
		magic.Fail(c, 1202, err.Error(), err)
		return
	}
	magic.Success(c, "删除菜单成功", 1)
}
