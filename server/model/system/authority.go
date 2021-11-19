package system

import (
	"time"
)

// SysAuthority 授权
type SysAuthority struct {
	CreatedAt       time.Time      // 创建时间
	UpdatedAt       time.Time      // 更新时间
	DeletedAt       *time.Time     `sql:"index"`
	AuthorityID     string         `json:"authorityId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"` // 角色ID
	AuthorityName   string         `json:"authorityName" gorm:"comment:角色名"`                                    // 角色名
	ParentID        string         `json:"parentId" gorm:"comment:父角色ID"`                                       // 父角色ID
	DataAuthorityID []SysAuthority `json:"dataAuthorityId" gorm:"many2many:sys_data_authority_id"`
	Children        []SysAuthority `json:"children" gorm:"-"`
	SysBaseMenus    []SysBaseMenu  `json:"menus" gorm:"many2many:sys_authority_menus;"`
	DefaultRouter   string         `json:"defaultRouter" gorm:"comment:默认菜单;default:dashboard"` // 默认菜单(默认dashboard)
}
