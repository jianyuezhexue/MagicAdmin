package system

import (
	"github.com/jianyuezhexue/MagicAdmin/model"
)

// Meta 属性
type Meta struct {
	Title string `json:"title" form:"title" binding:"required,max=40"` // 菜单名
	Icon  string `json:"icon" form:"icon" binding:"required,max=20"`   // 菜单图标
}

// Menu 后台菜单
type Menu struct {
	model.BaseOrm
	Meta      `json:"meta"`
	ParentId  uint64    `json:"parentId" form:"parentId" binding:"numeric"`      // 父菜单ID
	Path      string    `json:"path" form:"path" binding:"required,max=40"`      // 菜单path
	Name      string    `json:"name" form:"name" binding:"required,max=40"`      // 菜单name
	Hidden    bool      `json:"hidden" form:"hidden"`                            // 是否在列表隐藏
	Component string    `json:"component" form:"component" binding:"required"`   // 对应前端文件路径
	Sort      int       `json:"sort" form:"sort" binding:"numeric"`              // 排序标记
	Api       []Api     `json:"api" form:"api" gorm:"foreignKey:MenuId"`         // 菜单下的路由
	ExtAuth   []ExtAuth `json:"extAuth" form:"extAuth" gorm:"foreignKey:MenuId"` // 菜单下的拓展权限
	Children  []Menu    `json:"children" gorm:"-"`
}

// 菜单的ID和name
type MenuOption struct {
	Id    uint64 `json:"id"`    // 主键ID
	Title string `json:"title"` // 菜单名
}
