package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model/system"
	serviceSystem "github.com/jianyuezhexue/MagicAdmin/service/system"
)

type DictionaryCtr struct{}

var DictionaryApi = new(DictionaryCtr)

// 新建目录
func (d *DictionaryCtr) Create(c *gin.Context) {
	// 参数校验
	var form system.Dictionary
	err := c.ShouldBind(&form)
	if err != nil {
		magic.Fail(c, http.StatusBadRequest, err.Error(), form)
		return
	}

	// 逻辑处理
	res, err := serviceSystem.DictionaryApp.Create(form)
	if err != nil {
		magic.Fail(c, http.StatusBadGateway, err.Error(), res)
		return
	}

	magic.Success(c, "创建字典目录成功", res)
}

// 更新目录
func (d *DictionaryCtr) Update(c *gin.Context) {
	magic.Success(c, "更新字典目录成功", "")
}

// 删除目录
func (d *DictionaryCtr) Delete(c *gin.Context) {
	magic.Success(c, "删除字典目录成功", "")
}

// 分页目录
func (d *DictionaryCtr) List(c *gin.Context) {
	magic.Success(c, "分页查询字典目录成功", "")
}

// 单条目录
func (d *DictionaryCtr) Item(c *gin.Context) {
	magic.Success(c, "查询单条字典目录成功", "")
}
