package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model/system"
	serviceSystem "github.com/jianyuezhexue/MagicAdmin/service/system"
)

// todo:自定义验证报错信息

// ConfigInfo
func ConfigInfo(c *gin.Context) {
	res := magic.Config
	magic.Success(c, "打印所有的配置文件", res)
}

// Register 用户注册账号
func Register(c *gin.Context) {
	// 表单验证
	var form system.FormRegister
	err := c.ShouldBind(&form)
	if err != nil {
		magic.Fail(c, http.StatusBadRequest, err.Error(), form)
		return
	}

	// 逻辑处理
	res, err := serviceSystem.Register(form)
	if err != nil {
		magic.Fail(c, http.StatusBadGateway, err.Error(), res)
		return
	}

	// 结果返回
	magic.Success(c, "注册成功", res)
}

// Login 用户登录
func Login(c *gin.Context) {
	// 参数校验
	var form system.FormLogin
	err := c.ShouldBind(&form)
	if err != nil {
		magic.Fail(c, http.StatusBadRequest, err.Error(), form)
		return
	}

	// 逻辑实现
	res, err := serviceSystem.Login(form)
	if err != nil {
		magic.Fail(c, http.StatusBadGateway, err.Error(), res)
		return
	}

	// 结果返回
	magic.Success(c, "登录成功", res)
}
