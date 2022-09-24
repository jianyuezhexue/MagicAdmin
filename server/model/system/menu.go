package system

import (
	"github.com/jianyuezhexue/MagicAdmin/model"
)

// Meta 属性
type Meta struct {
	KeepAlive bool   `json:"keepAlive" form:"keepAlive"`                   // 是否缓存
	Title     string `json:"title" form:"title" binding:"required,max=40"` // 菜单名
	Icon      string `json:"icon" form:"icon" binding:"required,max=20"`   // 菜单图标
	CloseTab  bool   `json:"closeTab" form:"closeTab"`                     // 自动关闭tab
}

// Menu 后台菜单
type Menu struct {
	model.BaseOrm
	Meta      `json:"meta"`
	ParentId  uint   `json:"parentId"`  // 父菜单ID
	Path      string `json:"path"`      // 路由path
	Name      string `json:"name"`      // 路由name
	Hidden    bool   `json:"hidden"`    // 是否在列表隐藏
	Component string `json:"component"` // 对应前端文件路径
	Sort      int    `json:"sort"`      // 排序标记
	Children  []Menu `json:"children" gorm:"-"`
}

type FormMenu struct {
	ParentId  uint   `json:"parentId" form:"parentId" binding:"numeric"`    // 父菜单ID
	Path      string `json:"path" form:"path" binding:"required,max=40"`    // 路由path
	Name      string `json:"name" form:"name" binding:"required,max=40"`    // 路由name
	Hidden    bool   `json:"hidden" form:"hidden"`                          // 是否在列表隐藏
	Component string `json:"component" form:"component" binding:"required"` // 对应前端文件路径
	Sort      int    `json:"sort" form:"sort" binding:"numeric"`            // 排序标记
	Meta      Meta   `json:"meta" gorm:"embedded"`                          // 附加属性
}
