package system

import (
	"github.com/jianyuezhexue/MagicAdmin/model"
)

// Authority 授权
type Authority struct {
	model.BaseOrm
	Id            uint64       `json:"id" form:"id"`                       // 主键ID
	Name          string       `json:"name" form:"name"`                   // 角色名
	Desc          string       `json:"desc" form:"desc"`                   // 简介
	Pid           uint64       `json:"pid" form:"pid"`                     // 父角色ID
	DefaultRouter string       `json:"defaultRouter" form:"DefaultRouter"` // 默认菜单(默认dashboard)
	MenuIds       string       `json:"menuIds" form:"menuIds"`             // 权限下的ID
	Children      []*Authority `json:"children" gorm:"-"`
}

// 设置角色菜单
type SetMenuAuth struct {
	model.IdArr
	Id uint64 `json:"id" form:"id" binding:"required"` // 主键ID
}
