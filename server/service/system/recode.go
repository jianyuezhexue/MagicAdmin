package system

import (
	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model"
	"github.com/jianyuezhexue/MagicAdmin/model/system"
)

type RecodeServer struct{}

var RecodeApp = new(RecodeServer)

// 保存记录
func (r *RecodeServer) Create(data system.Record) (err error) {
	// 存储数据
	err = magic.Orm.Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

// 分页查询
func (r *RecodeServer) List(data system.SearchRecord) magic.BackData {
	// 初始化DB
	db := magic.Orm.Model(&system.Record{})

	// 组合搜索条件

	// 查询总数
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return magic.Back(2900, err.Error(), err.Error())
	}

	// 查询列表数据
	limit := data.PageSize
	offset := (data.Page - 1) * data.PageSize
	var list []system.Record
	err = db.Limit(limit).Offset(offset).Order("id DESC").Find(&list).Error
	if err != nil {
		return magic.Back(2902, err.Error(), err.Error())
	}

	// 组合返回数据
	res := model.ResPageData{
		List:     list,
		Total:    total,
		Page:     data.Page,
		PageSize: data.PageSize,
	}

	// 返回数据
	return magic.Back(0, "分页查询操作记录成功", res)
}
