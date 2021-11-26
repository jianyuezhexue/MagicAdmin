package system

import (
	"github.com/jianyuezhexue/MagicAdmin/model"
)

// Meta 属性
type Meta struct {
	KeepAlive   bool   `json:"keepAlive"`   // 是否缓存
	DefaultMenu bool   `json:"defaultMenu"` // 是否是基础路由（开发中）
	Title       string `json:"title"`       // 菜单名
	Icon        string `json:"icon"`        // 菜单图标
	CloseTab    bool   `json:"closeTab"`    // 自动关闭tab
}

// Menu 后台菜单
type Menu struct {
	model.BaseOrm
	MenuLevel     uint                              `json:"-"`
	ParentId      string                            `json:"parentId"`  // 父菜单ID
	Path          string                            `json:"path"`      // 路由path
	Name          string                            `json:"name"`      // 路由name
	Hidden        bool                              `json:"hidden"`    // 是否在列表隐藏
	Component     string                            `json:"component"` // 对应前端文件路径
	Sort          int                               `json:"sort"`      // 排序标记
	Meta          `json:"meta" gorm:"comment:附加属性"` // 附加属性
	SysAuthoritys []SysAuthority                    `json:"authoritys" gorm:"many2many:sys_authority_menus;"`
	Children      []Menu                            `json:"children" gorm:"-"`
}
