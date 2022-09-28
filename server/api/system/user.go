package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model/system"
	serviceSystem "github.com/jianyuezhexue/MagicAdmin/service/system"
	"github.com/mojocn/base64Captcha"
)

// ConfigInfo
func ConfigInfo(c *gin.Context) {
	res := magic.Config
	magic.Success(c, "打印所有的配置文件", res)
}

// 获取验证码
var store = base64Captcha.DefaultMemStore

func Captcha(c *gin.Context) {
	// 生成验证码
	var keyLong int = 6
	driver := base64Captcha.NewDriverDigit(80, 240, keyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	if err != nil {
		magic.Fail(c, 1201, err.Error(), err)
		return
	}

	// 组合返回数据
	var res = system.CaptchaResponse{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: keyLong,
	}

	// 结果返回
	magic.Success(c, "验证码获取成功", res)
}

// Register 用户注册账号
func Register(c *gin.Context) {
	// 表单验证
	var param system.FormRegister
	err := c.ShouldBind(&param)
	if err != nil {
		magic.Fail(c, http.StatusBadRequest, err.Error(), param)
		return
	}

	// 逻辑处理
	res, err := serviceSystem.Register(param)
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
	res, err := serviceSystem.Login(param)
	if err != nil {
		magic.Fail(c, http.StatusBadGateway, err.Error(), res)
		return
	}

	// 结果返回
	magic.Success(c, "登录成功", res)
}

// 用户信息
func UserInfo(c *gin.Context) {
	// 解密uuid
	uuid := magic.TokenInfo(c).UUID

	// 逻辑处理
	res, err := serviceSystem.UserInfo(uuid)
	if err != nil {
		magic.Fail(c, http.StatusBadGateway, err.Error(), res)
		return
	}

	// 结果返回
	magic.Success(c, "查询用户信息成功", res)
}
