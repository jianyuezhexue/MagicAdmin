package system

import (
	"errors"
	"strconv"
	"strings"

	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model"
	"github.com/jianyuezhexue/MagicAdmin/model/system"
	"gorm.io/gorm"
)

// MenuIds 通过权限ID查出所有菜单ID
func (a *AuthorityServer) MyMenuTree(authorityId int) (treeMenus []system.Menu, err error) {
	// 查询菜单IDs
	var authorites system.Authority
	err1 := magic.Orm.Where("id = ?", authorityId).First(&authorites).Error
	if err1 != nil {
		return nil, err1
	}
	menuIds := authorites.MenuIds

	// 查菜单
	var myMenus []system.Menu
	idSlice := strings.Split(menuIds, ",")
	err = magic.Orm.Where("id IN ?", idSlice).Order("sort").Preload("Api").Find(&myMenus).Error
	if err != nil {
		return nil, err
	}

	// 树形转换
	treeMenus = MenuApp.makeTree(myMenus, 0)
	return treeMenus, err
}

type AuthorityServer struct{}

var AuthorityApp = new(AuthorityServer)

// 列表转树形
func MakeTrees(listData []*system.Authority, Pid int) []*system.Authority {
	trees := make([]*system.Authority, 0)
	for _, item := range listData {
		if item.Pid == Pid {
			item.Children = MakeTrees(listData, item.Id)
			trees = append(trees, item)
		}
	}
	return trees
}

// 查询树形列表
func (a *AuthorityServer) TreeList() (res magic.PageResult, err error) {
	// 初始化DB
	var list []*system.Authority
	err = magic.Orm.Order("id").Find(&list).Error
	if err != nil {
		return res, errors.New("数据库请求失败")
	}

	// 树形转化
	tree := MakeTrees(list, 0)

	// 组合返回数据
	res = magic.PageResult{
		List:     tree,
		Total:    999,
		Page:     1,
		PageSize: 10,
	}

	// 返回数据
	return res, err
}

// 创建角色
func (a *AuthorityServer) Create(data system.Authority) (res system.Authority, err error) {
	// 角色名查重
	var find system.Authority
	err = magic.Orm.Where("name = ?", data.Name).Find(&find).Error
	if err != nil {
		return find, err
	}

	// 重名提醒
	if data.Name == find.Name {
		return find, errors.New("角色名有重复，请检查")
	}

	// 保存数据
	err = magic.Orm.Create(&data).Error
	return data, err
}

// 修改角色
func (a *AuthorityServer) Update(data system.Authority) (res system.Authority, err error) {
	// 查询是否存在
	var find system.Authority
	err = magic.Orm.Where("id = ?", data.Id).Find(&find).Error
	if err != nil {
		return find, err
	}

	// 异常提醒
	if err == gorm.ErrRecordNotFound {
		return find, errors.New("您编辑的角色不存在")
	}

	// 更新数据
	err = magic.Orm.Updates(&data).Error
	return data, err
}

// 删除角色
func (a *AuthorityServer) Delete(id model.GetById) (res []system.Authority, err error) {
	// 查询是否存在
	var find []system.Authority
	err = magic.Orm.Where("id = ?", id.Id).Or("pid = ?", id.Id).Find(&find).Error
	if err != nil {
		return find, err
	}

	// 异常提醒
	if err == gorm.ErrRecordNotFound {
		return find, errors.New("您删除的角色不存在")
	}

	// 包含子集提醒
	if len(find) > 1 {
		return find, errors.New("当前角色包含子角色")
	}

	// 更新数据
	err = magic.Orm.Delete(&find).Error
	return find, err
}

// 设置角色菜单权限
func (a *AuthorityServer) SetMenuAuth(data system.SetAuth) (res system.SetAuth, err error) {
	// // 删除当前角色菜单缓存
	// cacheKey := "myMenuTree"
	// err = magic.LocalCache.Delete(cacheKey)
	// if err != nil {
	// 	return res, err
	// }

	// 修改角色菜单权限
	menuIdStr := strings.Join(data.Data, ",")
	err = magic.Orm.Model(&system.Authority{}).Where("id = ?", data.Id).Update("menuIds", menuIdStr).Error
	if err != nil {
		return res, errors.New("系统繁忙，请稍后再试")
	}
	return data, err
}

// 设置角色API权限
func (a *AuthorityServer) SetApiAuth(data system.SetAuth) magic.BackData {

	// 查询API信息
	var apis []system.Api
	err := magic.Orm.Where("id", data.Data).Find(&apis).Error
	if err != nil {
		return magic.Back(-2200, "系统繁忙，请稍后再试", "DB跪了")
	}

	// 组合casbin的数据
	casbinRules := make([][]string, 0)
	for _, item := range apis {
		itemRule := make([]string, 0)
		itemRule = append(itemRule, strconv.Itoa(data.Id))
		itemRule = append(itemRule, item.Method)
		itemRule = append(itemRule, item.Route)
		casbinRules = append(casbinRules, itemRule)
	}

	// 删除当前角色下的规则
	enforcer := CasbinApp.Casbin()
	_, err = enforcer.RemoveFilteredPolicy(0, strconv.Itoa(data.Id))
	if err != nil {
		return magic.Back(2201, "系统繁忙，请稍后再试", err.Error())
	}

	// 开启事务
	tx := magic.Orm.Begin()

	// 设置API权限
	apiIdStr := strings.Join(data.Data, ",")
	err = tx.Model(&system.Authority{}).Where("id = ?", data.Id).Update("apiIds", apiIdStr).Error
	if err != nil {
		tx.Rollback()
		return magic.Back(2202, "系统繁忙，请稍后再试", "设置API权限错误")
	}

	// 设置casbin规则
	_, err = enforcer.AddPolicies(casbinRules)
	if err != nil {
		tx.Rollback()
		return magic.Back(2203, "系统繁忙，请稍后再试", "设置casbin规则错误")
	}

	// 提交事务
	tx.Commit()

	// 重载casbin规则
	err = enforcer.LoadPolicy()
	if err != nil {
		// todo 这里要记录系统日志
		return magic.Back(2203, "重载规则失败", "重载规则失败")
	}

	// 返回结果
	return magic.Back(0, "设置角色成功", data)
}

// 设置角色数据/按钮权限
func (a *AuthorityServer) SetExtAuth(data system.SetAuth) (res system.SetAuth, err error) {
	menuIdStr := strings.Join(data.Data, ",")
	err = magic.Orm.Model(&system.Authority{}).Where("id = ?", data.Id).Update("extAuthIds", menuIdStr).Error
	if err != nil {
		return res, errors.New("系统繁忙，请稍后再试")
	}
	return data, err
}
