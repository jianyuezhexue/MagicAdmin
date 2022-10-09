package system

import "github.com/jianyuezhexue/MagicAdmin/model"

// 拓展权限结构
type ExtAuth struct {
	model.BaseOrm
	MenuId   uint64 `json:"menuId"`   // 菜单ID
	MenuName string `json:"menuName"` // 菜单名称
	Type     string `json:"type"`     // 权限类型:按钮权限，数据权限
	Name     string `json:"name"`     // 权限中文名
	Val      string `json:"val"`      // 字段值
}
