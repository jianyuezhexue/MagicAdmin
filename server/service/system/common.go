package system

import (
	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model"
	"github.com/jianyuezhexue/MagicAdmin/model/system"
)

type CommonServer struct{}

var CommonApp = new(CommonServer)

// 分页查询媒体列表
func (c *CommonServer) MediaList(data model.PageInfo) (res model.ResPageData, err error) {

	// 初始化化表
	db := magic.Orm.Model(&system.MediaFiles{})

	// 查总数
	var total int64
	err = db.Count(&total).Error
	if err != nil {
		return res, err
	}

	// 查数据
	limit := data.PageSize
	offset := (data.Page - 1) * data.PageSize
	var list []system.MediaFiles
	err = db.Limit(limit).Offset(offset).Find(&list).Error
	if err != nil {
		return res, err
	}

	// 组合返回数据
	res.List = list
	res.Total = total
	res.Page = data.Page
	res.PageSize = data.PageSize

	return res, err
}
