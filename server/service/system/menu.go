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

type MenuServer struct{}

var MenuApp = new(MenuServer)

// 转树形结构
func (m *MenuServer) MakeTree(datas []system.Menu, ParentId uint64) []system.Menu {
	trees := make([]system.Menu, 0)

	for _, item := range datas {
		if item.ParentId == ParentId {
			item.Children = MenuApp.MakeTree(datas, item.Id)
			trees = append(trees, item)
		}
	}
	return trees
}

// 菜单树
func (m *MenuServer) MenuTree() (menus []system.Menu, err error) {

	// 查询所有菜单
	err = magic.Orm.Find(&menus).Error
	if err != nil {
		return nil, err
	}

	// 树形转换
	treeMenuus := MenuApp.MakeTree(menus, 0)

	// 返回数据
	return treeMenuus, err
}

// 获取动态菜单树
func (m *MenuServer) MyMenu(authorityId int) (menus []system.Menu, err error) {
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
	treeMenuus := MenuApp.MakeTree(myMenus, 0)
	return treeMenuus, err
}

// 分页获取菜单列表
func (m *MenuServer) Menus(pageInfo model.PageInfo) (list model.ResPageData, err error) {
	// 查菜单
	var menus []system.Menu
	err = magic.Orm.Order("sort").Find(&menus).Error
	if err != nil {
		return list, err
	}
	// 树形转换
	treeMenuus := MenuApp.MakeTree(menus, 0)

	// 数据封装
	list.List = treeMenuus
	list.Page = 1
	list.PageSize = 999
	list.Total = len(menus)
	return list, err
}

// 根据id查询菜单
func (m *MenuServer) FindMenu(id model.GetById) (res system.Menu, err error) {
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
func (m *MenuServer) UpdateMenu(menu system.Menu) (res system.Menu, err error) {
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
func (m *MenuServer) CreateMenu(menu system.Menu) (res system.Menu, err error) {
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
func (m *MenuServer) DeleteMenu(id model.GetById) (err error) {
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
