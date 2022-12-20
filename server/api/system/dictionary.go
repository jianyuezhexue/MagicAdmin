package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model"
	"github.com/jianyuezhexue/MagicAdmin/model/system"
	serviceSystem "github.com/jianyuezhexue/MagicAdmin/service/system"
)

type DictionaryCtr struct{}

var DictionaryApi = new(DictionaryCtr)

// 新建目录
func (d *DictionaryCtr) Create(c *gin.Context) {
	// 参数校验
	var param system.Dictionary
	err := c.ShouldBind(&param)
	if err != nil {
		magic.HttpFail(c, http.StatusBadRequest, err.Error(), param)
		return
	}

	// 逻辑处理
	res, err := serviceSystem.DictionaryApp.Create(param)
	if err != nil {
		magic.HttpFail(c, http.StatusBadGateway, err.Error(), res)
		return
	}

	magic.Success(c, "创建字典目录成功", res)
}

// 分页目录
func (d *DictionaryCtr) List(c *gin.Context) {
	// 参数校验
	var param system.SearchDictionary
	err := c.ShouldBind(&param)
	if err != nil {
		magic.HttpFail(c, http.StatusBadRequest, err.Error(), param)
		return
	}

	// 逻辑处理
	res, err := serviceSystem.DictionaryApp.List(param)
	if err != nil {
		magic.HttpFail(c, http.StatusBadGateway, err.Error(), res)
		return
	}
	magic.Success(c, "分页查询字典目录成功", res)
}

// 单条目录
func (d *DictionaryCtr) Item(c *gin.Context) {
	// 参数校验
	var id model.GetById
	err := c.ShouldBindUri(&id)
	if err != nil {
		magic.HttpFail(c, http.StatusBadRequest, err.Error(), id)
		return
	}

	// 逻辑处理
	res, err := serviceSystem.DictionaryApp.Item(id)
	if err != nil {
		magic.HttpFail(c, 1202, err.Error(), res)
		return
	}

	// 返回结果
	magic.Success(c, "查询单条字典目录成功", res)
}

// 更新目录
func (d *DictionaryCtr) Update(c *gin.Context) {
	// 参数校验
	var param system.Dictionary
	err := c.ShouldBind(&param)
	if err != nil {
		magic.HttpFail(c, http.StatusBadRequest, err.Error(), param)
		return
	}

	// 逻辑处理
	res, err := serviceSystem.DictionaryApp.Update(param)
	if err != nil {
		magic.HttpFail(c, http.StatusBadRequest, err.Error(), res)
		return
	}

	// 返回结果
	magic.Success(c, "更新字典目录成功", res)
}

// 删除目录
func (d *DictionaryCtr) Delete(c *gin.Context) {
	// 参数校验
	var id model.GetById
	err := c.ShouldBindUri(&id)
	if err != nil {
		magic.HttpFail(c, http.StatusBadRequest, err.Error(), id)
		return
	}

	// 逻辑处理
	res, err := serviceSystem.DictionaryApp.Delete(id)
	if err != nil {
		magic.HttpFail(c, 1202, err.Error(), res)
		return
	}

	magic.Success(c, "删除字典目录成功", "")
}

// 根据key查找子目录
func (d *DictionaryCtr) DictionaryByKey(c *gin.Context) {
	// 参数校验
	var key model.GetByKey
	err := c.ShouldBindUri(&key)
	if err != nil {
		magic.HttpFail(c, http.StatusBadRequest, err.Error(), key)
		return
	}

	// 逻辑处理
	res := serviceSystem.DictionaryApp.DictionaryByKey(key.Key)
	magic.HttpSuccess(c, res)
}

// 根据key查找子目录
func (d *DictionaryCtr) DictionaryByKeys(c *gin.Context) {
	// 参数校验
	var keys system.Keys
	err := c.ShouldBind(&keys)
	if err != nil {
		magic.HttpFail(c, http.StatusBadRequest, err.Error(), keys.Keys)
		return
	}

	// 逻辑处理
	res := serviceSystem.DictionaryApp.DictionaryByKeys(keys.Keys)
	magic.HttpSuccess(c, res)
}
