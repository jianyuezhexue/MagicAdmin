package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model"
	"github.com/jianyuezhexue/MagicAdmin/model/system"
	serviceSystem "github.com/jianyuezhexue/MagicAdmin/service/system"
)

type UserCtr struct{}

var UserApi = new(UserCtr)

// 新建后台用户
func (u *UserCtr) Register(c *gin.Context) {
	// 表单验证
	var param system.User
	err := c.ShouldBind(&param)
	if err != nil {
		magic.HttpFail(c, http.StatusBadRequest, err.Error(), param)
		return
	}

	// 逻辑处理
	res, err := serviceSystem.UserApp.Register(param)
	if err != nil {
		magic.HttpFail(c, http.StatusBadGateway, err.Error(), res)
		return
	}

	// 结果返回
	magic.Success(c, "注册成功", res)
}

// 用户登录
func (u *UserCtr) Login(c *gin.Context) {
	// 参数校验
	var param system.FormLogin
	err := c.ShouldBind(&param)
	if err != nil {
		magic.HttpFail(c, http.StatusBadRequest, err.Error(), param)
		return
	}

	// 验证码校验
	if !store.Verify(param.CaptchaId, param.Captcha, true) {
		magic.HttpFail(c, 1202, "验证码错误", param)
		return
	}

	// 逻辑实现
	res, err := serviceSystem.UserApp.Login(param)
	if err != nil {
		magic.HttpFail(c, http.StatusBadGateway, err.Error(), res)
		return
	}

	// 结果返回
	magic.Success(c, "登录成功", res)
}

// 用户信息|登录之后
func (u *UserCtr) UserInfo(c *gin.Context) {
	// 解密uuid
	uuid := magic.TokenInfo(c).UUID

	// 逻辑处理
	res, err := serviceSystem.UserApp.UserInfo(uuid)
	if err != nil {
		magic.HttpFail(c, http.StatusBadGateway, err.Error(), res)
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
		magic.HttpFail(c, http.StatusBadRequest, err.Error(), param)
		return
	}

	// 逻辑处理
	res, err := serviceSystem.UserApp.List(param)
	if err != nil {
		magic.HttpFail(c, http.StatusBadGateway, err.Error(), res)
		return
	}
	magic.Success(c, "分页查询用户列表成功", res)
}

// 删除用户
func (u UserCtr) Delete(c *gin.Context) {
	// 接受参数
	var id model.GetById
	err := c.ShouldBindUri(&id)
	if err != nil {
		magic.HttpFail(c, http.StatusBadRequest, err.Error(), id)
		return
	}

	// 逻辑处理
	res, err := serviceSystem.UserApp.Delete(id)
	if err != nil {
		magic.HttpFail(c, http.StatusBadGateway, err.Error(), res)
		return
	}
	magic.Success(c, "删除用户成功", res)
}

// 更新用户信息
func (u *UserCtr) Update(c *gin.Context) {
	var param system.User
	err := c.ShouldBind(&param)
	if err != nil {
		magic.HttpFail(c, http.StatusBadRequest, err.Error(), param)
		return
	}

	// 逻辑处理
	res, err := serviceSystem.UserApp.Update(param)
	if err != nil {
		magic.HttpFail(c, http.StatusBadGateway, err.Error(), res)
		return
	}
	magic.Success(c, "更新用户成功", res)
}

// 设置用户角色
func (u *UserCtr) SetUserAuth(c *gin.Context) {
	// 参数校验
	var param system.SetUserAuth
	err := c.ShouldBind(&param)
	if err != nil {
		magic.HttpFail(c, http.StatusBadRequest, err.Error(), param)
		return
	}

	// 逻辑处理
	res, err := serviceSystem.UserApp.SetUserAuth(param)
	if err != nil {
		magic.HttpFail(c, http.StatusBadGateway, err.Error(), res)
		return
	}
	magic.Success(c, "设置用户角色成功", res)
}

// 设置可用状态
func (u *UserCtr) SetUserStatus(c *gin.Context) {
	var id model.GetById
	err := c.ShouldBind(&id)
	if err != nil {
		magic.HttpFail(c, http.StatusBadRequest, err.Error(), id)
		return
	}

	// 逻辑处理
	res, err := serviceSystem.UserApp.SetUserStatus(id)
	if err != nil {
		magic.HttpFail(c, http.StatusBadGateway, err.Error(), res)
		return
	}
	magic.Success(c, "设置用户角色成功", res)
}

// 重置密码|123456
func (u UserCtr) ReSetPwd(c *gin.Context) {
	var id model.GetById
	err := c.ShouldBindUri(&id)
	if err != nil {
		magic.HttpFail(c, http.StatusBadRequest, err.Error(), id)
		return
	}

	// 逻辑处理
	res, err := serviceSystem.UserApp.ReSetPwd(id)
	if err != nil {
		magic.HttpFail(c, http.StatusBadGateway, err.Error(), res)
		return
	}
	magic.Success(c, "重置密码成功", res)
}

// 切换角色
func (u UserCtr) SwitchAuth(c *gin.Context) {

	// 接收参数
	var param system.SwitchAuth
	err := c.ShouldBind(&param)
	if err != nil {
		magic.HttpFail(c, http.StatusBadRequest, err.Error(), param)
		return
	}

	// 获取当前用户UUID & 合并参数
	userInfo := magic.TokenInfo(c)
	param.UUID = userInfo.UUID

	// 逻辑处理
	res, err := serviceSystem.UserApp.SwitchAuth(param)
	if err != nil {
		magic.HttpFail(c, http.StatusBadGateway, err.Error(), res)
		return
	}

	// 返回头带最新的token
	c.Header("new-token", res.Token)
	magic.Success(c, "重置密码成功", res)
}
