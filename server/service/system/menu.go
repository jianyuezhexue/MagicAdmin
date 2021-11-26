package service

import (
	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model/system"
)

func getMenuTreeMap(authorityId string) (err error, treeMap map[string][]system.Menu) {
	var allMenus []system.Menu
	treeMap = make(map[string][]system.Menu)
	err = magic.Orm.Where("authority_id = ?", authorityId).Order("sort").Preload("Parameters").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return err, treeMap
}

// 获取动态菜单树
func GetMenuTree() (menus []system.Menu, err error) {
	err, menuTree := getMenuTreeMap(authorityId)
	menus = menuTree["0"]
	for i := 0; i < len(menus); i++ {
		err = getChildrenList(&menus[i], menuTree)
	}
	return menus, err
}

// 获取子菜单
func getChildrenList(menu *system.Menu, treeMap map[string][]system.Menu) (err error) {
	menu.Children = treeMap[menu.MenuId]
	for i := 0; i < len(menu.Children); i++ {
		err = getChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

// // 获取路由分页
// func GetInfoList() (err error, list interface{}, total int64) {
// 	var menuList []system.SysBaseMenu
// 	err, treeMap := menuService.getBaseMenuTreeMap()
// 	menuList = treeMap["0"]
// 	for i := 0; i < len(menuList); i++ {
// 		err = menuService.getBaseChildrenList(&menuList[i], treeMap)
// 	}
// 	return err, menuList, total
// }

// // 获取菜单的子菜单
// func getBaseChildrenList(menu *system.SysBaseMenu, treeMap map[string][]system.SysBaseMenu) (err error) {
// 	menu.Children = treeMap[strconv.Itoa(int(menu.ID))]
// 	for i := 0; i < len(menu.Children); i++ {
// 		err = menuService.getBaseChildrenList(&menu.Children[i], treeMap)
// 	}
// 	return err
// }

// // 添加基础路由
// func AddBaseMenu(menu system.SysBaseMenu) error {
// 	if !errors.Is(magic.Orm.Where("name = ?", menu.Name).First(&system.SysBaseMenu{}).Error, gorm.ErrRecordNotFound) {
// 		return errors.New("存在重复name，请修改name")
// 	}
// 	return magic.Orm.Create(&menu).Error
// }

// // 获取路由总树map
// func getBaseMenuTreeMap() (err error, treeMap map[string][]system.SysBaseMenu) {
// 	var allMenus []system.SysBaseMenu
// 	treeMap = make(map[string][]system.SysBaseMenu)
// 	err = magic.Orm.Order("sort").Preload("Parameters").Find(&allMenus).Error
// 	for _, v := range allMenus {
// 		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
// 	}
// 	return err, treeMap
// }

// // 获取基础路由树
// func GetBaseMenuTree() (err error, menus []system.SysBaseMenu) {
// 	err, treeMap := menuService.getBaseMenuTreeMap()
// 	menus = treeMap["0"]
// 	for i := 0; i < len(menus); i++ {
// 		err = menuService.getBaseChildrenList(&menus[i], treeMap)
// 	}
// 	return err, menus
// }

// // 为角色增加menu树
// func AddMenuAuthority(menus []system.SysBaseMenu, authorityId string) (err error) {
// 	var auth system.SysAuthority
// 	auth.AuthorityId = authorityId
// 	auth.SysBaseMenus = menus
// 	err = AuthorityServiceApp.SetMenuAuthority(&auth)
// 	return err
// }

// // 查看当前角色树
// func GetMenuAuthority(info *request.GetAuthorityId) (err error, menus []system.Menu) {
// 	err = magic.Orm.Where("authority_id = ? ", info.AuthorityId).Order("sort").Find(&menus).Error
// 	//sql := "SELECT authority_menu.keep_alive,authority_menu.default_menu,authority_menu.created_at,authority_menu.updated_at,authority_menu.deleted_at,authority_menu.menu_level,authority_menu.parent_id,authority_menu.path,authority_menu.`name`,authority_menu.hidden,authority_menu.component,authority_menu.title,authority_menu.icon,authority_menu.sort,authority_menu.menu_id,authority_menu.authority_id FROM authority_menu WHERE authority_menu.authority_id = ? ORDER BY authority_menu.sort ASC"
// 	//err = magic.Orm.Raw(sql, authorityId).Scan(&menus).Error
// 	return err, menus
// }
