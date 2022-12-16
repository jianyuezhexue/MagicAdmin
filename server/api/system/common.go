package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model"
	"github.com/jianyuezhexue/MagicAdmin/model/system"
	serviceSystem "github.com/jianyuezhexue/MagicAdmin/service/system"
	"github.com/mojocn/base64Captcha"
)

type CommonCtr struct{}

var CommonApi = new(CommonCtr)

// 获取验证码
var store = base64Captcha.DefaultMemStore

func (co *CommonCtr) Captcha(c *gin.Context) {
	// 生成验证码
	var keyLong int = 6
	driver := base64Captcha.NewDriverDigit(80, 240, keyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	if err != nil {
		magic.HttpFail(c, 1201, err.Error(), err)
		return
	}

	// 组合返回数据
	var res magic.BackData
	code := system.CaptchaResponse{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: keyLong,
	}
	res = magic.Back(0, "获取验证码成功", code)

	// 结果返回
	magic.HttpSuccess(c, res)
}

// 分页查询媒体库列表
func (co *CommonCtr) MediaList(c *gin.Context) {
	var param model.PageInfo
	err := c.ShouldBind(&param)
	if err != nil {
		magic.HttpFail(c, http.StatusBadRequest, err.Error(), param)
		return
	}

	// 逻辑处理
	res, err := serviceSystem.CommonApp.MediaList(param)
	if err != nil {
		magic.HttpFail(c, http.StatusBadGateway, err.Error(), res)
		return
	}
	magic.Success(c, "分页查询媒体库列表成功", res)
}
