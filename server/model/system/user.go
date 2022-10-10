package system

import (
	"github.com/jianyuezhexue/MagicAdmin/model"
	uuid "github.com/satori/go.uuid"
)

// 后台系统用户
type User struct {
	model.BaseOrm
	UUID         uuid.UUID   `json:"uuid"`         // 用户UUID
	UserName     string      `json:"userName"`     // 用户登录名
	Password     string      `json:"-" `           // 用户登录密码
	NickName     string      `json:"nickName"`     // 用户昵称
	SideMode     string      `json:"sideMode" `    // 用户侧边主题
	HeadImg      string      `json:"headImg" `     // 用户头像
	BaseColor    string      `json:"baseColor"`    // 基础颜色
	ActiveColor  string      `json:"activeColor"`  // 活跃颜色
	AuthorityId  string      `json:"authorityId"`  // 用户角色ID
	AuthorityIds model.Strs  `json:"authorityIds"` // 用户角色ID
	Phone        string      `json:"phone"`        // 用户手机号
	Email        string      `json:"email"`        // 用户邮箱
	Enable       int         `json:"enable"`       //用户是否被冻结 1正常 2冻结
	Token        string      `json:"token" gorm:"-"`
	Authority    Authority   `json:"authority" gorm:"-"`
	Authorities  []Authority `json:"authorities" gorm:"-"`
}

// 注册表单提交数据结构
type FormRegister struct {
	UserName    string `json:"userName" form:"userName" binding:"required,min=3,max=10"`
	Password    string `json:"password" form:"password" binding:"required"`
	NickName    string `json:"nickName" form:"nickName" binding:"required,min=3,max=10"`
	AuthorityId string `json:"authorityId" form:"authorityId" binding:"required,numeric"`
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
	User
}

// 设置角色
type SetUserAuth struct {
	Id  int        `json:"id" uri:"id" form:"id" binding:"required"`    // 用户ID
	Ids model.Strs `josn:"ids" uri:"ids" form:"ids" binding:"required"` // 角色ID数组
}
