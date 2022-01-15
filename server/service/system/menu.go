package system

import (
	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model"
	"github.com/jianyuezhexue/MagicAdmin/model/system"
)

// TreeTransform 菜单树型转换
func TreeTransform(data []system.Menu) (trees []system.Menu) {
	// 一级找子集
	for index := range data {
		if data[index].ParentId == 0 {
			// todo:性能结合实际业务待优化
			trees = append(trees, FindChild(data[index], data))
		}
	}
	return trees
}

// FindChild 菜单树型转换-子集找子集
func FindChild(node system.Menu, data []system.Menu) system.Menu {
	for index := range data {
		if node.Id == data[index].ParentId {
			node.Children = append(node.Children, FindChild(data[index], data))
		}
	}
	return node
}

// MyMenu 获取动态菜单树
func MyMenu(authorityId int) (menus []system.Menu, err error) {
	// 查权限[防击穿]
	menuIds, err, _ := magic.SingleFlight.Do("ServiceMenu", func() (interface{}, error) {
		return MenuIds(authorityId)
	})
	if err != nil {
		return nil, err
	}

	// 查菜单
	var myMenus []system.Menu
	err = magic.Orm.Where("id IN ?", menuIds).Order("sort").Find(&myMenus).Error
	if err != nil {
		return nil, err
	}

	// 树形转换
	treeMenuus := TreeTransform(myMenus)
	return treeMenuus, err
}

// Menus 分页获取菜单列表
func Menus(pageInfo model.PageInfo) (list model.ResPageData, err error) {
	// 查菜单
	var menus []system.Menu
	err = magic.Orm.Order("sort").Find(&menus).Error
	if err != nil {
		return list, err
	}
	// 树形转换
	treeMenuus := TreeTransform(menus)

	// 数据封装
	list.List = treeMenuus
	list.Page = 1
	list.PageSize = 999
	list.Total = len(menus)
	return list, err
}

// CreateMenu 创建菜单
func CreateMenu(form system.FormMenu) (res system.Menu, err error) {
	// 鉴权

	return res, err
}
