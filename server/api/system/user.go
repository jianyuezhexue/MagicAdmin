package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model/system"
	serviceSystem "github.com/jianyuezhexue/MagicAdmin/service/system"
)

type UserCtr struct{}

var UserApi = new(UserCtr)

// Register 用户注册账号
func (u *UserCtr) Register(c *gin.Context) {
	// 表单验证
	var param system.FormRegister
	err := c.ShouldBind(&param)
	if err != nil {
		magic.Fail(c, http.StatusBadRequest, err.Error(), param)
		return
	}

	// 逻辑处理
	res, err := serviceSystem.UserApp.Register(param)
	if err != nil {
		magic.Fail(c, http.StatusBadGateway, err.Error(), res)
		return
	}

	// 结果返回
	magic.Success(c, "注册成功", res)
}

// Login 用户登录
func (u *UserCtr) Login(c *gin.Context) {
	// 参数校验
	var param system.FormLogin
	err := c.ShouldBind(&param)
	if err != nil {
		magic.Fail(c, http.StatusBadRequest, err.Error(), param)
		return
	}

	// 验证码校验
	if !store.Verify(param.CaptchaId, param.Captcha, true) {
		magic.Fail(c, 1202, "验证码错误", param)
		return
	}

	// 逻辑实现
	res, err := serviceSystem.UserApp.Login(param)
	if err != nil {
		magic.Fail(c, http.StatusBadGateway, err.Error(), res)
		return
	}

	// 结果返回
	magic.Success(c, "登录成功", res)
}

// 用户信息
func (u *UserCtr) UserInfo(c *gin.Context) {
	// 解密uuid
	uuid := magic.TokenInfo(c).UUID

	// 逻辑处理
	res, err := serviceSystem.UserApp.UserInfo(uuid)
	if err != nil {
		magic.Fail(c, http.StatusBadGateway, err.Error(), res)
		return
	}

	// 结果返回
	magic.Success(c, "查询用户信息成功", res)
}

// 用户列表
func (u *UserCtr) List(c *gin.Context) {
	// 参数校验
	var param system.SearchUser
	err := c.ShouldBind(&param)
	if err != nil {
		magic.Fail(c, http.StatusBadRequest, err.Error(), param)
		return
	}

	// 逻辑处理
	res, err := serviceSystem.UserApp.List(param)
	if err != nil {
		magic.Fail(c, http.StatusBadGateway, err.Error(), res)
		return
	}
	magic.Success(c, "分页查询用户列表成功", res)
}
