package system

import (
	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model"
	"github.com/jianyuezhexue/MagicAdmin/model/system"
)

type CommonServer struct{}

var CommonApp = new(CommonServer)

// 分页查询媒体列表
func (c *CommonServer) MediaList(data model.PageInfo) magic.BackData {

	// 初始化化表
	db := magic.Orm.Model(&system.MediaFiles{})

	// 查总数
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return magic.Back(21400, err.Error(), total)
	}

	// 查数据
	limit := data.PageSize
	offset := (data.Page - 1) * data.PageSize
	var list []system.MediaFiles
	err = db.Limit(limit).Offset(offset).Find(&list).Error
	if err != nil {
		return magic.Back(21401, err.Error(), total)
	}

	// 组合返回数据
	res := model.ResPageData{List: list, Total: total, Page: data.Page, PageSize: data.PageSize}
	return magic.Back(0, "分页查询媒体列表成功", res)
}
