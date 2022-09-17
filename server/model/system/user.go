package system

import (
	"github.com/jianyuezhexue/MagicAdmin/model"
	uuid "github.com/satori/go.uuid"
)

// User 后台系统用户
type User struct {
	model.BaseOrm
	UUID        uuid.UUID `json:"uuid"`        // 用户UUID
	UserName    string    `json:"userName"`    // 用户登录名
	Password    string    `json:"-" `          // 用户登录密码
	NickName    string    `json:"nickName"`    // 用户昵称
	SideMode    string    `json:"sideMode" `   // 用户侧边主题
	HeadImg     string    `json:"headImg" `    // 用户头像
	BaseColor   string    `json:"baseColor"`   // 基础颜色
	ActiveColor string    `json:"activeColor"` // 活跃颜色
	AuthorityId string    `json:"authorityId"` // 用户角色ID
	Token       string    `json:"token" gorm:"-"`
	Authority   Authority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	// Authorities []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`
}

// FormRegister 注册表单提交数据结构
type FormRegister struct {
	UserName    string `json:"userName" form:"userName" binding:"required,min=3,max=10"`
	Password    string `json:"password" form:"password" binding:"required"`
	NickName    string `json:"nickName" form:"nickName" binding:"required,min=3,max=10"`
	AuthorityId string `json:"authorityId" form:"authorityId" binding:"required,numeric"`
}

// FormLogin 登录表单提交数据结构
type FormLogin struct {
	UserName  string `json:"userName" form:"userName" binding:"required,min=3,max=10"` // 用户名
	Password  string `json:"password" form:"password" binding:"required"`              // 密码
	CaptchaId string `json:"captchaId" form:"captchaId" binding:"required"`            // 验证码ID
	Captcha   string `json:"captcha" form:"captcha" binding:"required"`                // 验证码字符
}

// 验证码
type CaptchaResponse struct {
	CaptchaId     string `json:"captchaId"`
	PicPath       string `json:"picPath"`
	CaptchaLength int    `json:"captchaLength"`
}
