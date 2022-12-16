package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model/system"
	serviceSystem "github.com/jianyuezhexue/MagicAdmin/service/system"
)

type RecodeCtr struct{}

var RecodeApi = new(RecodeCtr)

// 查询记录
func (r *RecodeCtr) List(c *gin.Context) {
	// 参数校验
	var param system.SearchRecord
	err := c.ShouldBind(&param)
	if err != nil {
		magic.HttpFail(c, http.StatusBadRequest, err.Error(), param)
		return
	}

	// 逻辑处理
	res := serviceSystem.RecodeApp.List(param)
	if err != nil {
		magic.HttpFail(c, http.StatusBadGateway, res.Msg, res.Data)
		return
	}
	magic.Success(c, res.Msg, res.Data)
}
