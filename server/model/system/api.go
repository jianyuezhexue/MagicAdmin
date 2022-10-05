package system

import (
	"github.com/jianyuezhexue/MagicAdmin/model"
)

// Api属性
type Api struct {
	model.BaseOrm
	MenuId   uint64 `json:"menuId"`   // 菜单ID
	MenuName string `json:"menuName"` // 菜单名称
	Method   string `json:"method"`   // 请求类型
	Name     string `json:"name"`     // 中文名
	Route    string `json:"route"`    // 路由地址
}

// 分页查询字典目录
type SearchApi struct {
	model.PageInfo
	Api
}

// 分页返回
type ApiPageResult struct {
	List       any          `json:"list"`
	Total      int64        `json:"total"`
	Page       int          `json:"page" form:"page"`         // 页码
	PageSize   int          `json:"pageSize" form:"pageSize"` // 每页大小
	MenuOption []MenuOption `json:"menuOption"`
}
