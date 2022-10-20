package system

import (
	"errors"
	"strconv"
	"strings"

	"github.com/gomodule/redigo/redis"
	"github.com/jianyuezhexue/MagicAdmin/config"
	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model"
	"github.com/jianyuezhexue/MagicAdmin/model/system"
	"gorm.io/gorm"
)

// MenuIds 通过权限ID查出所有菜单ID
func MenuIds(authorityId int) (res []string, err error) {
	// 先查缓存
	menuIds := ""
	key := config.MenueIds + strconv.Itoa(authorityId)
	menuIds, err = magic.Redis.Get(key)
	if err != nil && err != redis.ErrNil { // 区分空值和报错
		return nil, err
	}

	// 没缓存查DB
	if err == redis.ErrNil {
		var authorites system.Authority
		err = magic.Orm.Where("id = ?", authorityId).First(&authorites).Error
		if err != nil {
			return nil, err
		}
		menuIds = authorites.MenuIds

		// 设置缓存
		magic.Redis.Set(key, menuIds, -1)
	}

	// 字符串解析
	slice := strings.Split(menuIds, ",")
	return slice, nil
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

	// 验证是否有子角色
	magic.Print(find)

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
	// 删除当前角色菜单缓存
	key := config.MenueIds + strconv.Itoa(int(data.Id))
	_, err = magic.Redis.Del(key)
	if err != nil {
		return res, err
	}

	// 修改角色菜单权限
	menuIdStr := strings.Join(data.Data, ",")
	err = magic.Orm.Model(&system.Authority{}).Where("id = ?", data.Id).Update("menuIds", menuIdStr).Error
	if err != nil {
		return res, errors.New("系统繁忙，请稍后再试")
	}
	return data, err
}

// 设置角色API权限
func (a *AuthorityServer) SetApiAuth(data system.SetAuth) (res system.SetAuth, err error) {

	// 查询API信息
	var apis []system.Api
	err = magic.Orm.Where("id", data.Data).Find(&apis).Error
	if err != nil {
		return res, errors.New("系统繁忙，请稍后再试")
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
	remove, err := enforcer.RemoveFilteredPolicy(0, strconv.Itoa(data.Id))
	if err != nil || !remove {
		return res, errors.New("系统繁忙，请稍后再试")
	}

	// 开启事务
	tx := magic.Orm.Begin()

	// 设置API权限
	apiIdStr := strings.Join(data.Data, ",")
	err = tx.Model(&system.Authority{}).Where("id = ?", data.Id).Update("apiIds", apiIdStr).Error
	if err != nil {
		tx.Rollback()
		return res, errors.New("系统繁忙，请稍后再试")
	}

	// 设置casbin规则
	_, err = enforcer.AddPolicies(casbinRules)
	if err != nil {
		tx.Rollback()
		return res, errors.New("系统繁忙，请稍后再试")
	}

	// 提交事务
	tx.Commit()

	// 重载casbin规则
	err = enforcer.LoadPolicy()
	if err != nil {
		// todo 这里要记录系统日志
		return res, errors.New("设置失败[重载规则失败]")
	}

	return data, err
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
