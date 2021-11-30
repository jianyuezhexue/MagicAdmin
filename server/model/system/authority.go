package system

import (
	"time"
)

// Authority 授权
type Authority struct {
	CreatedAt     time.Time  // 创建时间
	UpdatedAt     time.Time  // 更新时间
	DeletedAt     *time.Time `sql:"index"`
	AuthorityId   string     `json:"authorityId"`   // 角色ID
	AuthorityName string     `json:"authorityName"` // 角色名
	ParentId      string     `json:"parentId"`      // 父角色ID
	// DataAuthorityId []Authority `json:"dataAuthorityId"  gorm:"-"`
	Children      []Authority `json:"children" gorm:"-"`
	DefaultRouter string      `json:"defaultRouter"` // 默认菜单(默认dashboard)
	MenuIds       string      `json:"menuIds"`       // 权限下的ID
}
