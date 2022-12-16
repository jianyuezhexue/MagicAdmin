package system

import (
	"errors"

	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model"
	"github.com/jianyuezhexue/MagicAdmin/model/system"
	"gorm.io/gorm"
)

type MenuServer struct{}

var MenuApp = new(MenuServer)

// 转树形结构|私有
func (m *MenuServer) makeTree(datas []system.Menu, ParentId uint64) []system.Menu {
	trees := make([]system.Menu, 0)

	for _, item := range datas {
		if item.ParentId == ParentId {
			item.Children = MenuApp.makeTree(datas, item.Id)
			trees = append(trees, item)
		}
	}
	return trees
}

// 全部菜单树|设置权限用
func (m *MenuServer) MenuTree(id model.GetById) (res any, err error) {
	// 查询所有菜单
	var menus []system.Menu
	err = magic.Orm.Preload("Api").Preload("ExtAuth").Order("sort ASC").Find(&menus).Error
	if err != nil {
		return nil, err
	}

	// 树形转换
	treeMenus := MenuApp.makeTree(menus, 0)

	// 查询当前权限
	var auth system.Authority
	err = magic.Orm.Where("id = ?", id.Id).Find(&auth).Error
	if err != nil {
		return nil, err
	}

	// 返回数据
	result := make(map[string]any)
	result["treeMenus"] = treeMenus
	result["auth"] = auth
	return result, err
}

// 我的权限菜单树
func (m *MenuServer) MyMenu(authorityId int) (menus any, err error) {
	// 查数据库[防击穿]
	treeMenus, err, _ := magic.SingleFlight.Do("ServiceMenu", func() (any, error) {
		return AuthorityApp.MyMenuTree(authorityId)
	})
	if err != nil {
		return nil, err
	}

	// 返回结果
	return treeMenus, err
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
	treeMenus := MenuApp.makeTree(menus, 0)

	// 数据封装
	list.List = treeMenus
	list.Page = 1
	list.PageSize = 999
	list.Total = len(menus)
	return list, err
}

// 查询所有菜单的ID和name
func (m *MenuServer) MenuOption() (res []system.MenuOption, err error) {
	// todo 这里用redis缓存

	var list []system.MenuOption
	err = magic.Orm.Model(&system.Menu{}).Find(&list).Error
	if err != nil {
		return res, errors.New("DB跪了")
	}
	return list, err
}

// 根据id查询菜单
func (m *MenuServer) FindMenu(id model.GetById) (res system.Menu, err error) {
	// 查询数据
	var menu system.Menu
	err = magic.Orm.Where("id = ?", id.Id).Preload("Api").Preload("ExtAuth").Find(&menu).Error
	if err != nil {
		return res, err
	}
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

	// 校验[idx_type_route]重复
	if len(menu.Api) > 0 {
		routes := make([][]any, 0)
		for _, api := range menu.Api {
			item := make([]any, 0)
			item = append(item, api.Method, api.Route)
			routes = append(routes, item)
		}
		apis := []system.Api{}
		err = magic.Orm.Where("(method,route) IN ?", routes).Find(&apis).Error
		if err != nil {
			return res, errors.New("数据库跪了")
		}

		// 非当前菜单包含算重复
		for _, item := range apis {
			if item.MenuId != menu.Id {
				return res, errors.New("路由有重复，请检查")
			}
		}
	}

	// 关联更新数据
	err = magic.Orm.Session(&gorm.Session{FullSaveAssociations: true}).Updates(menu).Error
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
	if err != nil {
		return res, errors.New("数据库跪了")
	}

	// 重名异常
	if menu.Name == find.Name {
		return res, errors.New("菜单名有重复，请修改")
	}

	// 校验[idx_type_route]重复
	if len(menu.Api) > 0 {
		routes := make([][]any, 0)
		for _, api := range menu.Api {
			item := make([]any, 0)
			item = append(item, api.Method, api.Route)
			routes = append(routes, item)
		}
		apis := []system.Api{}
		err = magic.Orm.Where("(method,route) IN ?", routes).Find(&apis).Error
		if err != nil {
			return res, errors.New("数据库跪了")
		}

		// 重复提醒
		if len(apis) > 0 {
			return res, errors.New("路由有重复，请检查")
		}
	}

	// 保存数据
	err = magic.Orm.Create(&menu).Error
	return menu, err
}

// 删除菜单
func (m *MenuServer) DeleteMenu(id model.GetById) (err error) {
	// 查询是否存在
	var menu []system.Menu
	err = magic.Orm.Where("id = ?", id.Id).Or("parentId = ?", id.Id).Find(&menu).Error
	if err != nil {
		return err
	}

	// 异常提醒
	if err == gorm.ErrRecordNotFound {
		return errors.New("您删除的菜单不存在")
	}

	// 包含子集提醒
	if len(menu) > 1 {
		return errors.New("当前菜单包含子菜单，请先删除子菜单")
	}

	tx := magic.Orm.Begin()
	// 删除菜单
	err = tx.Delete(&menu).Error
	if err != nil {
		tx.Rollback()
		return errors.New("DB跪了")
	}
	// 删除API ｜ Unscoped 硬删除
	var api []system.Api
	err = tx.Where("menuId = ?", menu[0].Id).Unscoped().Delete(&api).Error
	if err != nil {
		tx.Rollback()
		return errors.New("DB跪了")
	}
	// 删除extAuth |Unscoped 硬删除
	var extAuth []system.ExtAuth
	err = tx.Where("menuId = ?", menu[0].Id).Unscoped().Delete(&extAuth).Error
	if err != nil {
		tx.Rollback()
		return errors.New("DB跪了")
	}
	tx.Commit()

	// 返回结果
	return err
}
