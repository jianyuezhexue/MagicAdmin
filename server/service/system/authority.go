package system

import (
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
