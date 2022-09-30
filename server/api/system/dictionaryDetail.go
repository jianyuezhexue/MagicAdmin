package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model"
	"github.com/jianyuezhexue/MagicAdmin/model/system"
	serviceSystem "github.com/jianyuezhexue/MagicAdmin/service/system"
)

type DictionaryDetailCtr struct{}

var DictionaryDetailApi = new(DictionaryDetailCtr)

// 新建目录
func (d *DictionaryDetailCtr) Create(c *gin.Context) {
	// 参数校验
	var param system.DictionaryDetail
	err := c.ShouldBind(&param)
	if err != nil {
		magic.Fail(c, http.StatusBadRequest, err.Error(), param)
		return
	}

	// 逻辑处理
	res, err := serviceSystem.DictionaryDetailApp.Create(param)
	if err != nil {
		magic.Fail(c, http.StatusBadGateway, err.Error(), res)
		return
	}

	magic.Success(c, "创建字典目录成功", res)
}

// 分页目录
func (d *DictionaryDetailCtr) List(c *gin.Context) {
	// 参数校验
	var param system.SearchDictionaryDetail
	err := c.ShouldBind(&param)
	if err != nil {
		magic.Fail(c, http.StatusBadRequest, err.Error(), param)
		return
	}

	// 逻辑处理
	res, err := serviceSystem.DictionaryDetailApp.List(param)
	if err != nil {
		magic.Fail(c, http.StatusBadGateway, err.Error(), res)
		return
	}
	magic.Success(c, "分页查询字典目录成功", res)
}

// 单条目录
func (d *DictionaryDetailCtr) Item(c *gin.Context) {
	// 参数校验
	var id model.GetById
	err := c.ShouldBindUri(&id)
	if err != nil {
		magic.Fail(c, http.StatusBadRequest, err.Error(), id)
		return
	}

	// 逻辑处理
	res, err := serviceSystem.DictionaryDetailApp.Item(id)
	if err != nil {
		magic.Fail(c, 1202, err.Error(), res)
		return
	}

	// 返回结果
	magic.Success(c, "查询单条字典目录成功", res)
}

// 更新目录
func (d *DictionaryDetailCtr) Update(c *gin.Context) {
	// 参数校验
	var param system.DictionaryDetail
	err := c.ShouldBind(&param)
	if err != nil {
		magic.Fail(c, http.StatusBadRequest, err.Error(), param)
		return
	}

	// 逻辑处理
	res, err := serviceSystem.DictionaryDetailApp.Update(param)
	if err != nil {
		magic.Fail(c, http.StatusBadRequest, err.Error(), res)
		return
	}

	// 返回结果
	magic.Success(c, "更新字典目录成功", res)
}

// 删除目录
func (d *DictionaryDetailCtr) Delete(c *gin.Context) {
	// 参数校验
	var id model.GetById
	err := c.ShouldBindUri(&id)
	if err != nil {
		magic.Fail(c, http.StatusBadRequest, err.Error(), id)
		return
	}

	// 逻辑处理
	res, err := serviceSystem.DictionaryDetailApp.Delete(id)
	if err != nil {
		magic.Fail(c, 1202, err.Error(), res)
		return
	}

	magic.Success(c, "删除字典目录成功", "")
}
