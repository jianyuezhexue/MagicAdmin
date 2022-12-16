package community

import (
	"github.com/jianyuezhexue/MagicAdmin/model/community"
	"github.com/jianyuezhexue/MagicAdmin/service"
)

type ArticleService struct{}

var ArticleApp = new(ArticleService)

// 文章列表
func (a *ArticleService) List() {

}

// 新建文章
func (a *ArticleService) Create(data community.Article) service.BackData {

	return service.Back(0, "新建文章成功", data)
}

// 编辑文章
func (a *ArticleService) Update() {

}

// 删除文章
func (a *ArticleService) Delete() {

}
