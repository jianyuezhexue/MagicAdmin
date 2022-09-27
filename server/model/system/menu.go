package system

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/duke-git/lancet/v2/convertor"
	"github.com/jianyuezhexue/MagicAdmin/model"
)

// Meta 属性
type Meta struct {
	Title string `json:"title" form:"title" binding:"required,max=40"` // 菜单名
	Icon  string `json:"icon" form:"icon" binding:"required,max=20"`   // 菜单图标
}

// Api属性
type Api struct {
	Type  string `json:"type"`  // 请求类型
	Name  string `json:"name"`  // 中文名
	Route string `json:"route"` // 路由地址
}

// 拓展权限结构
type ExtAuth struct {
	Type string `json:"type"` // 权限类型:按钮权限，数据权限
	Name string `json:"name"` // 权限中文名
	Val  string `json:"val"`  // 字段值
}

type Array[T Api | ExtAuth] []T

// 字符串转成数组返回
func (a *Array[T]) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to scan Array value:", value))
	}
	if len(bytes) > 0 {
		return json.Unmarshal(bytes, a)
	}
	*a = make([]T, 0)
	return nil
}

// 数组转成字符串存储
func (a Array[T]) Value() (driver.Value, error) {
	if a == nil {
		return "[]", nil
	}
	return convertor.ToString(a), nil
}

// Menu 后台菜单
type Menu struct {
	model.BaseOrm
	Meta      `json:"meta"`
	ParentId  uint           `json:"parentId"`  // 父菜单ID
	Path      string         `json:"path"`      // 路由path
	Name      string         `json:"name"`      // 路由name
	Hidden    bool           `json:"hidden"`    // 是否在列表隐藏
	Component string         `json:"component"` // 对应前端文件路径
	Sort      int            `json:"sort"`      // 排序标记
	Apis      Array[Api]     `json:"apis"`      // 菜单下的路由
	ExtAuth   Array[ExtAuth] `json:"extAuth"`   // 菜单下的拓展权限
	Children  []Menu         `json:"children" gorm:"-"`
}

type FormMenu struct {
	ParentId  uint           `json:"parentId" form:"parentId" binding:"numeric"`    // 父菜单ID
	Path      string         `json:"path" form:"path" binding:"required,max=40"`    // 路由path
	Name      string         `json:"name" form:"name" binding:"required,max=40"`    // 路由name
	Hidden    bool           `json:"hidden" form:"hidden"`                          // 是否在列表隐藏
	Component string         `json:"component" form:"component" binding:"required"` // 对应前端文件路径
	Sort      int            `json:"sort" form:"sort" binding:"numeric"`            // 排序标记
	Apis      Array[Api]     `json:"apis" gorm:"apis"`                              // 菜单下的路由
	ExtAuth   Array[ExtAuth] `json:"extAuth" gorm:"extAuth"`                        // 菜单下的拓展权限
	Meta      Meta           `json:"meta" gorm:"embedded"`                          // 附加属性
}
