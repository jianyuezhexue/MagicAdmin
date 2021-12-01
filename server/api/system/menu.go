package system

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model"
	service "github.com/jianyuezhexue/MagicAdmin/service/system"
)

// 获取用户动态路由
func MyMenu(c *gin.Context) {
	// 参数获取
	userInfo := magic.TokenInfo(c)
	authorityId, _ := strconv.Atoi(userInfo.AuthorityId)

	// 逻辑处理
	res, err := service.MyMenu(authorityId)
	if err != nil {
		magic.Fail(c, http.StatusBadGateway, err.Error(), res)
		return
	}
	// 结果返回
	magic.Success(c, "获取我的菜单成功", res)
}

// 分页获取基础menu列表
func Menus(c *gin.Context) {
	// 参数获取
	var pageInfo model.PageInfo
	err := c.ShouldBind(&pageInfo)
	if err != nil {
		magic.Fail(c, http.StatusBadRequest, err.Error(), pageInfo)
		return
	}

	// 逻辑处理
	res, err := service.Menus(pageInfo)
	if err != nil {
		magic.Fail(c, http.StatusBadGateway, err.Error(), res)
		return
	}
	// 结果返回
	magic.Success(c, "获取菜单列表成功", res)
}

// // 增加menu和角色关联关系
// func AddMenuAuthority(c *gin.Context) {
// 	var authorityMenu systemReq.AddMenuAuthorityInfo
// 	_ = c.ShouldBindJSON(&authorityMenu)
// 	if err := utils.Verify(authorityMenu, utils.AuthorityIdVerify); err != nil {
// 		response.FailWithMessage(err.Error(), c)
// 		return
// 	}
// 	if err := menuService.AddMenuAuthority(authorityMenu.Menus, authorityMenu.AuthorityId); err != nil {
// 		magic.Logger.Error("添加失败!", zap.Any("err", err))
// 		response.FailWithMessage("添加失败", c)
// 	} else {
// 		response.OkWithMessage("添加成功", c)
// 	}
// }

// // 获取指定角色menu
// func GetMenuAuthority(c *gin.Context) {
// 	var param request.GetAuthorityId
// 	_ = c.ShouldBindJSON(&param)
// 	if err := utils.Verify(param, utils.AuthorityIdVerify); err != nil {
// 		response.FailWithMessage(err.Error(), c)
// 		return
// 	}
// 	if err, menus := menuService.GetMenuAuthority(&param); err != nil {
// 		magic.Logger.Error("获取失败!", zap.Any("err", err))
// 		response.FailWithDetailed(systemRes.SysMenusResponse{Menus: menus}, "获取失败", c)
// 	} else {
// 		response.OkWithDetailed(gin.H{"menus": menus}, "获取成功", c)
// 	}
// }

// // 新增菜单
// func AddBaseMenu(c *gin.Context) {
// 	var menu system.SysBaseMenu
// 	_ = c.ShouldBindJSON(&menu)
// 	if err := utils.Verify(menu, utils.MenuVerify); err != nil {
// 		response.FailWithMessage(err.Error(), c)
// 		return
// 	}
// 	if err := utils.Verify(menu.Meta, utils.MenuMetaVerify); err != nil {
// 		response.FailWithMessage(err.Error(), c)
// 		return
// 	}
// 	if err := menuService.AddBaseMenu(menu); err != nil {
// 		magic.Logger.Error("添加失败!", zap.Any("err", err))

// 		response.FailWithMessage("添加失败", c)
// 	} else {
// 		response.OkWithMessage("添加成功", c)
// 	}
// }

// // 删除菜单
// func DeleteBaseMenu(c *gin.Context) {
// 	var menu request.GetById
// 	_ = c.ShouldBindJSON(&menu)
// 	if err := utils.Verify(menu, utils.IdVerify); err != nil {
// 		response.FailWithMessage(err.Error(), c)
// 		return
// 	}
// 	if err := baseMenuService.DeleteBaseMenu(menu.ID); err != nil {
// 		magic.Logger.Error("删除失败!", zap.Any("err", err))
// 		response.FailWithMessage("删除失败", c)
// 	} else {
// 		response.OkWithMessage("删除成功", c)
// 	}
// }

// // 更新菜单
// func UpdateBaseMenu(c *gin.Context) {
// 	var menu system.SysBaseMenu
// 	_ = c.ShouldBindJSON(&menu)
// 	if err := utils.Verify(menu, utils.MenuVerify); err != nil {
// 		response.FailWithMessage(err.Error(), c)
// 		return
// 	}
// 	if err := utils.Verify(menu.Meta, utils.MenuMetaVerify); err != nil {
// 		response.FailWithMessage(err.Error(), c)
// 		return
// 	}
// 	if err := baseMenuService.UpdateBaseMenu(menu); err != nil {
// 		magic.Logger.Error("更新失败!", zap.Any("err", err))
// 		response.FailWithMessage("更新失败", c)
// 	} else {
// 		response.OkWithMessage("更新成功", c)
// 	}
// }

// // 根据id获取菜单
// func GetBaseMenuById(c *gin.Context) {
// 	var idInfo request.GetById
// 	_ = c.ShouldBindJSON(&idInfo)
// 	if err := utils.Verify(idInfo, utils.IdVerify); err != nil {
// 		response.FailWithMessage(err.Error(), c)
// 		return
// 	}
// 	if err, menu := baseMenuService.GetBaseMenuById(idInfo.ID); err != nil {
// 		magic.Logger.Error("获取失败!", zap.Any("err", err))
// 		response.FailWithMessage("获取失败", c)
// 	} else {
// 		response.OkWithDetailed(systemRes.SysBaseMenuResponse{Menu: menu}, "获取成功", c)
// 	}
// }
