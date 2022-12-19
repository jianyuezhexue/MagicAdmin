package community

import (
	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model/community"
)

type ArticleService struct{}

var ArticleApp = new(ArticleService)

// 文章列表
func (a *ArticleService) List() {

}

// 新建文章
func (a *ArticleService) Create(data community.Article) magic.BackData {

	// 保存数据
	err := magic.Orm.Create(&data).Error
	if err != nil {
		return magic.Back(2100, err.Error(), data)
	}

	// 返回成功
	return magic.Back(0, "新建文章成功", data)
}

// 编辑文章
func (a *ArticleService) Update() {

}

// 删除文章
func (a *ArticleService) Delete() {

}
