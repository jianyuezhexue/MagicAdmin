package system

import (
	"errors"
	"fmt"

	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model/system"
)

type DictionaryServer struct{}

var DictionaryApp = new(DictionaryServer)

// 新建目录
func (d DictionaryServer) Create(data system.Dictionary) (system.Dictionary, error) {
	// 查重Value值
	find := system.Dictionary{}
	err := magic.Orm.Where("value = ?", data.Value).Find(&find).Error
	fmt.Println("查出来的错误是:", err)
	if err != nil {
		return find, errors.New("数据库跪了")
	}

	// val值重复提醒
	if data.Value == find.Value {
		return find, errors.New("英文字典名不能重复，请修改")
	}

	// 保存数据
	err = magic.Orm.Create(&data).Error
	return data, err
}
