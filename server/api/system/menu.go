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

// MyMenu 获取用户动态路由
func MyMenu(c *gin.Context) {
	// 参数获取
	userInfo := magic.TokenInfo(c)
	authorityId, _ := strconv.Atoi(userInfo.AuthorityId)

	// 逻辑处理
	// res, err := service.MyMenu(authorityId)
	res, err := serviceSystem.MyMenu(authorityId)
	if err != nil {
		magic.Fail(c, http.StatusBadGateway, err.Error(), res)
		return
	}
	// 结果返回
	magic.Success(c, "获取我的菜单成功", res)
}

// Menus 分页获取基础menu列表
func Menus(c *gin.Context) {
	// 参数获取
	var pageInfo model.PageInfo
	err := c.ShouldBind(&pageInfo)
	if err != nil {
		magic.Fail(c, http.StatusBadRequest, err.Error(), pageInfo)
		return
	}

	// 逻辑处理
	res, err := serviceSystem.Menus(pageInfo)
	if err != nil {
		magic.Fail(c, http.StatusBadGateway, err.Error(), res)
		return
	}
	// 结果返回
	magic.Success(c, "获取菜单列表成功", res)
}

// CreateMenu 创建菜单
func CreateMenu(c *gin.Context) {
	// 参数校验
	var form system.FormMenu
	err := c.ShouldBind(&form)
	if err != nil {
		magic.Fail(c, http.StatusBadRequest, err.Error(), form)
		return
	}

	// 逻辑处理
	res, err := serviceSystem.CreateMenu(form)
	if err != nil {
		magic.Fail(c, http.StatusBadGateway, err.Error(), res)
		return
	}
	// 结果返回
	magic.Success(c, "创建菜单成功", res)
}

// UpdateMenu 更新菜单
func UpdateMenu(c *gin.Context) {
	// 参数校验

	magic.Success(c, "创建菜单成功", 1)
}

// DelMenu 删除菜单
func DelMenu(c *gin.Context) {
	// 参数校验

	magic.Success(c, "创建菜单成功", 1)
}

// 根据id获取菜单
func FindMenu(c *gin.Context) {
	// 参数校验

	magic.Success(c, "创建菜单成功", 1)
}

// 增加menu和角色关联关系

// 获取指定角色menu
