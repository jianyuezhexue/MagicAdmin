package system

import (
	"time"
)

// SysAuthority 授权
type SysAuthority struct {
	CreatedAt       time.Time      // 创建时间
	UpdatedAt       time.Time      // 更新时间
	DeletedAt       *time.Time     `sql:"index"`
	AuthorityID     string         `json:"authorityId"`   // 角色ID
	AuthorityName   string         `json:"authorityName"` // 角色名
	ParentID        string         `json:"parentId"`      // 父角色ID
	DataAuthorityID []SysAuthority `json:"dataAuthorityId"`
	Children        []SysAuthority `json:"children" gorm:"-"`
	DefaultRouter   string         `json:"defaultRouter"` // 默认菜单(默认dashboard)
}
