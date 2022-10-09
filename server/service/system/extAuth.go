package system

import (
	"errors"

	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model"
	"github.com/jianyuezhexue/MagicAdmin/model/system"
	"gorm.io/gorm"
)

type ExtAuthServer struct{}

var ExtAuthApp = new(ExtAuthServer)

// 删除菜单
func (a *ExtAuthServer) Delete(id model.GetById) (res system.ExtAuth, err error) {
	// 查询是否存在
	var find system.ExtAuth
	err = magic.Orm.Where("id = ?", id.ID).Find(&find).Error
	if err != nil {
		return res, err
	}

	// 不存在提醒
	if err == gorm.ErrRecordNotFound {
		return res, errors.New("您删除的角色不存在")
	}

	// 删除执行
	err = magic.Orm.Delete(&find).Error
	return res, err
}
