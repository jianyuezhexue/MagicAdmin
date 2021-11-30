package service

import (
	"fmt"
	"strconv"

	"github.com/gomodule/redigo/redis"
	"github.com/jianyuezhexue/MagicAdmin/config"
	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model/system"
)

// MenuIds 通过权限ID查出所有菜单ID
func MenuIds(authorityId int) (res []int, err error) {
	// 先查缓存
	menuIds := ""
	key := config.MenueIds + strconv.Itoa(authorityId)
	menuIds, err = magic.Redis.Get(key)
	if err != nil && err != redis.ErrNil { // 区分空值和报错
		return nil, err
	}
	fmt.Print(menuIds)

	// 没缓存查DB
	if err == redis.ErrNil {
		var authorites system.Authority
		// 防击穿查询
		_, err, _ = magic.SingleFlight.Do(key, func() (interface{}, error) {
			res := magic.Orm.Where("authorityId = ?", authorityId).First(&authorites)
			return authorites, res.Error
		})
	}

	// 字符串解析
	return nil, nil
}
