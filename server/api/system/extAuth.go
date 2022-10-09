package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model"
	serviceSystem "github.com/jianyuezhexue/MagicAdmin/service/system"
)

type ExtAuthCtr struct{}

var ExtAuthAPI = new(ExtAuthCtr)

// 删除拓展
func (e *ExtAuthCtr) Delete(c *gin.Context) {

	// 接收参数
	var id model.GetById
	err := c.ShouldBindUri(&id)
	if err != nil {
		magic.Fail(c, http.StatusBadGateway, err.Error(), id)
		return
	}

	// 逻辑处理
	res, err := serviceSystem.ExtAuthApp.Delete(id)
	if err != nil {
		magic.Fail(c, http.StatusBadGateway, err.Error(), res)
		return
	}

	// 返回成功
	magic.Success(c, "删除成功", res)
}
