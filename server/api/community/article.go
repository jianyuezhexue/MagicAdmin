package community

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model"
	"github.com/jianyuezhexue/MagicAdmin/model/community"
	communityService "github.com/jianyuezhexue/MagicAdmin/service/community"
)

type ArticleCtr struct{}

var Article = new(ArticleCtr)

// 文章列表
func (a *ArticleCtr) List(c *gin.Context) {
	// 接收参数
	var param community.SearchArticle
	err := c.ShouldBind(&param)
	if err != nil {
		magic.HttpFail(c, http.StatusBadRequest, err.Error(), param)
		return
	}

	// 逻辑处理
	res := communityService.ArticleApp.List(param)
	magic.HttpSuccess(c, res)
}

// 新建文章
func (a *ArticleCtr) Create(c *gin.Context) {
	// 接收参数
	var param community.Article
	err := c.ShouldBind(&param)
	if err != nil {
		magic.HttpFail(c, http.StatusBadRequest, err.Error(), param)
		return
	}

	// 补全|uuid,作者，IP地址
	tokenInfo := magic.TokenInfo(c)
	param.Uuid = tokenInfo.UUID.String()
	param.AuthName = tokenInfo.NickName
	param.IpAddress = c.ClientIP()

	// 逻辑处理
	res := communityService.ArticleApp.Create(param)
	magic.HttpSuccess(c, res)
}

// 查找文章
func (a *ArticleCtr) Find(c *gin.Context) {
	// 接收参数
	var id model.GetById
	err := c.ShouldBindUri(&id)
	if err != nil {
		magic.HttpFail(c, http.StatusBadRequest, err.Error(), id)
		return
	}

	// 逻辑处理
	res := communityService.ArticleApp.Find(id)
	magic.HttpSuccess(c, res)
}

// 编辑文章
func (a *ArticleCtr) Update(c *gin.Context) {

}

// 删除文章
func (a *ArticleCtr) Delete(c *gin.Context) {

}
