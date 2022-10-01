package system

import (
	"github.com/jianyuezhexue/MagicAdmin/model"
)

// Authority 授权
type Authority struct {
	model.BaseOrm
	AuthorityId   uint64       `json:"authorityId" form:"authorityId"`     // 角色ID
	AuthorityName string       `json:"authorityName" form:"authorityName"` // 角色名
	Desc          string       `json:"desc" form:"desc"`                   // 简介
	Pid           uint64       `json:"pid" form:"pid"`                     // 父角色ID
	DefaultRouter string       `json:"defaultRouter" form:"DefaultRouter"` // 默认菜单(默认dashboard)
	MenuIds       string       `json:"menuIds" form:"menuIds"`             // 权限下的ID
	Children      []*Authority `json:"children" gorm:"-"`
}
