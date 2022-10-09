package system

import (
	"github.com/gin-gonic/gin"
	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model/system"
	"github.com/mojocn/base64Captcha"
)

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
