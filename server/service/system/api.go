package system

import (
	"errors"

	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model"
	"github.com/jianyuezhexue/MagicAdmin/model/system"
	"gorm.io/gorm"
)

type ApiServer struct{}

var ApiApp = new(ApiServer)

// 分页查询
func (d *ApiServer) List(data system.SearchApi) (res system.ApiPageResult, err error) {
	// 初始化DB
	db := magic.Orm.Model(&system.Api{})

	// 组合搜索条件 | 可以用循环简化代码
	if data.Name != "" {
		db = db.Where("`name` LIKE ?", "%"+data.Name+"%")
	}

	// 查询总数
	var total int64
	err = db.Count(&total).Error
	if err != nil {
		return res, err
	}

	// 查询列表数据
	limit := data.PageSize
	offset := (data.Page - 1) * data.PageSize
	var list []system.Api
	err = db.Limit(limit).Offset(offset).Order("menuId").Find(&list).Error
	if err != nil {
		return res, errors.New("DB跪了")
	}

	// 查询菜单选项
	menuOption, err := MenuApp.MenuOption()
	if err != nil {
		return res, errors.New("系统繁忙，请稍后再试")
	}

	// 组合返回数据
	res = system.ApiPageResult{
		List:       list,
		MenuOption: menuOption,
		Total:      total,
		Page:       data.Page,
		PageSize:   data.PageSize,
	}

	// 返回数据
	return res, err
}

// 删除菜单
func (a *ApiServer) Delete(id model.GetById) (res system.Api, err error) {
	// 查询是否存在
	var find system.Api
	err = magic.Orm.Where("id = ?", id.ID).Find(&find).Error
	if err != nil {
		return res, err
	}

	// 不存在提醒
	if err == gorm.ErrRecordNotFound {
		return res, errors.New("您删除的角色不存在")
	}

	// 删除执行
	err = magic.Orm.Unscoped().Delete(&find).Error
	return res, err
}
