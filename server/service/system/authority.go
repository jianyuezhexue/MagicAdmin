package system

import (
	"errors"
	"strconv"
	"strings"

	"github.com/gomodule/redigo/redis"
	"github.com/jianyuezhexue/MagicAdmin/config"
	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model/system"
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
		err = magic.Orm.Where("authorityId = ?", authorityId).First(&authorites).Error
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
func MakeTrees(listData []*system.Authority, Pid uint64) []*system.Authority {
	trees := make([]*system.Authority, 0)
	for _, item := range listData {
		if item.Pid == Pid {
			item.Children = MakeTrees(listData, item.AuthorityId)
			trees = append(trees, item)
		}
	}
	return trees
}

// 查询树形列表
func (d *AuthorityServer) TreeList() (res magic.PageResult, err error) {
	// 初始化DB
	var list []*system.Authority
	err = magic.Orm.Order("authorityId").Find(&list).Error
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
