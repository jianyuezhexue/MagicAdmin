package system

import (
	"errors"
	"fmt"

	"github.com/duke-git/lancet/v2/convertor"
	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model"
	"github.com/jianyuezhexue/MagicAdmin/model/system"
	"gorm.io/gorm"
)

// 菜单树型转换
func TreeTransform(data []system.Menu) (trees []system.Menu) {
	// 一级找子集
	for index := range data {
		if data[index].ParentId == 0 {
			trees = append(trees, FindChild(data[index], data))
		}
	}
	return trees
}

// 菜单树型转换-子集找子集
func FindChild(node system.Menu, data []system.Menu) system.Menu {
	for index := range data {
		if node.Id == data[index].ParentId {
			node.Children = append(node.Children, FindChild(data[index], data))
		}
	}
	return node
}

// 获取动态菜单树
func MyMenu(authorityId int) (menus []system.Menu, err error) {
	// 查权限[防击穿]
	menuIds, err, _ := magic.SingleFlight.Do("ServiceMenu", func() (any, error) {
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

// 分页获取菜单列表
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
	magic.Logger.Info(convertor.ToString(menu))

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

// 创建菜单
func CreateMenu(menu system.Menu) (res system.Menu, err error) {
	// 先查数据
	find := system.Menu{}
	err = magic.Orm.Where("name = ?", menu.Name).Find(&find).Error
	fmt.Println("查出来的错误是:", err)
	if err != nil {
		return res, errors.New("数据库跪了")
	}

	// 重名异常
	if menu.Name == find.Name {
		return res, errors.New("存在重复name,请修改name")
	}

	// 保存数据
	err = magic.Orm.Create(&menu).Error
	return menu, err
}

// 删除菜单
func DeleteMenu(id model.GetById) (err error) {
	// 查询是否存在
	var menu system.Menu
	err = magic.Orm.Where("id = ?", id.ID).Find(&menu).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("您要删除的菜单不存在")
		}
		return err
	}

	// todo这里的权限系统要重新设计
	// 设计权限相关的操作

	// 删除菜单
	err = magic.Orm.Delete(&menu).Error
	return err
}
