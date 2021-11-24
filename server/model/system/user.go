package system

import (
	"github.com/jianyuezhexue/MagicAdmin/model"
	uuid "github.com/satori/go.uuid"
)

// User 后台系统用户
type User struct {
	model.BaseOrm
	UUID        uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`                                                         // 用户UUID
	Username    string    `json:"userName" gorm:"comment:用户登录名"`                                                      // 用户登录名
	Password    string    `json:"-"  gorm:"comment:用户登录密码"`                                                           // 用户登录密码
	NickName    string    `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                          // 用户昵称
	SideMode    string    `json:"sideMode" gorm:"default:dark;comment:用户侧边主题"`                                        // 用户侧边主题
	HeadImg     string    `json:"headImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	BaseColor   string    `json:"baseColor" gorm:"default:#fff;comment:基础颜色"`                                         // 基础颜色
	ActiveColor string    `json:"activeColor" gorm:"default:#1890ff;comment:活跃颜色"`                                    // 活跃颜色
	AuthorityID string    `json:"authorityId" gorm:"default:888;comment:用户角色ID"`                                      // 用户角色ID
	// Authority   SysAuthority   `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	// Authorities []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`
}

// FormRegister User register structure
type FormRegister struct {
	Username    string `json:"userName" form:"userName" binding:"required,min=3,max=10"`
	Password    string `json:"password" form:"password" binding:"required"`
	NickName    string `json:"nickName" form:"nickName" binding:"required,min=3,max=10"`
	AuthorityID string `json:"authorityId" form:"authorityId" binding:"required,numeric"`
}
