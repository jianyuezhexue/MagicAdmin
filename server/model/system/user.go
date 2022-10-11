package system

import (
	"github.com/jianyuezhexue/MagicAdmin/model"
	uuid "github.com/satori/go.uuid"
)

// 后台系统用户
type User struct {
	model.BaseOrm
	UUID         uuid.UUID   `json:"uuid"` // 用户UUID
	UserName     string      `json:"userName" form:"userName" binding:"required,min=1,max=10"`
	Password     string      `json:"-" form:"password" binding:""`
	NickName     string      `json:"nickName" form:"nickName" binding:"required,min=1,max=10"`
	Phone        string      `json:"phone" form:"phone" binding:"required"`                     // 用户手机号
	Email        string      `json:"email" form:"email" binding:"required"`                     // 用户邮箱
	HeadImg      string      `json:"headImg" form:"headImg" binding:"required"`                 // 用户头像
	AuthorityId  string      `json:"authorityId" form:"authorityId" binding:"required,numeric"` // 当前角色ID
	AuthorityIds model.Strs  `json:"authorityIds" form:"authorityIds" binding:"required"`       // 所有角色
	Enable       int         `json:"enable" form:"enable" binding:"required"`                   // 用户是否被冻结 1正常 2冻结
	SideMode     string      `json:"sideMode" `                                                 // 用户侧边主题
	BaseColor    string      `json:"baseColor"`                                                 // 基础颜色
	ActiveColor  string      `json:"activeColor"`                                               // 活跃颜色
	Token        string      `json:"token" gorm:"-"`
	Authority    Authority   `json:"authority" gorm:"-"`
	Authorities  []Authority `json:"authorities" gorm:"-"`
}

// 登录表单提交数据结构
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

// 分页查询字典详情
type SearchUser struct {
	model.PageInfo
	UserName string `json:"userName" form:"userName"`
}

// 设置角色
type SetUserAuth struct {
	Id  int        `json:"id" uri:"id" form:"id" binding:"required"`    // 用户ID
	Ids model.Strs `josn:"ids" uri:"ids" form:"ids" binding:"required"` // 角色ID数组
}
