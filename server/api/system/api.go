package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model"
	"github.com/jianyuezhexue/MagicAdmin/model/system"
	serviceSystem "github.com/jianyuezhexue/MagicAdmin/service/system"
)

type ApiCtr struct{}

var ApiAPI = new(ApiCtr)

// 分页角色列表
func (a *ApiCtr) List(c *gin.Context) {
	// 参数校验
	var param system.SearchApi
	err := c.ShouldBind(&param)
	if err != nil {
		magic.Fail(c, http.StatusBadRequest, err.Error(), param)
		return
	}

	// 逻辑处理
	res, err := serviceSystem.ApiApp.List(param)
	if err != nil {
		magic.Fail(c, http.StatusBadGateway, err.Error(), res)
		return
	}
	magic.Success(c, "分页查询字典目录成功", res)
}

// 删除API
func (a *ApiCtr) Delete(c *gin.Context) {

	// 接收参数
	var id model.GetById
	err := c.ShouldBindUri(&id)
	if err != nil {
		magic.Fail(c, http.StatusBadGateway, err.Error(), id)
		return
	}

	// 逻辑处理
	res, err := serviceSystem.ApiApp.Delete(id)
	if err != nil {
		magic.Fail(c, http.StatusBadGateway, err.Error(), res)
		return
	}

	// 返回成功
	magic.Success(c, "删除成功", res)
}
