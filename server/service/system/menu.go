package system

import (
	"errors"

	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model"
	"github.com/jianyuezhexue/MagicAdmin/model/system"
	"gorm.io/gorm"
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

// 根据id查询菜单
func FindMenu(id model.GetById) (res system.Menu, err error) {
	// 查询数据
	var menu system.Menu
	err = magic.Orm.Where("id = ?", id.ID).Find(&menu).Error
	if err != nil {
		return res, err
	}

	return menu, err
}

// 更新菜单
func UpdateMenu(menu system.Menu) (res system.Menu, err error) {
	// 查询数据
	var oldData system.Menu
	err = magic.Orm.Where("id = ?", menu.Id).Find(&oldData).Error
	if err != nil {
		return res, err
	}

	// 校验是否存在
	if err == gorm.ErrRecordNotFound {
		return res, errors.New("您编辑的菜单不存在！")
	}

	// 更新数据
	err = magic.Orm.Updates(menu).Error
	if err != nil {
		return res, err
	}

	return menu, err
}

// CreateMenu 创建菜单
func CreateMenu(menu system.Menu) (res system.Menu, err error) {
	// 查重名称
	err = magic.Orm.Where("name = ?", menu.Name).Find(&system.Menu{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return res, errors.New("存在重复name,请修改name")
	}

	// 保存数据
	err = magic.Orm.Create(&menu).Error
	return menu, err
}
