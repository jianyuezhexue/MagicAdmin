package community

import (
	"github.com/jianyuezhexue/MagicAdmin/magic"
	"github.com/jianyuezhexue/MagicAdmin/model"
	"github.com/jianyuezhexue/MagicAdmin/model/community"
)

type ArticleService struct{}

var ArticleApp = new(ArticleService)

// 文章列表
func (a *ArticleService) List(data community.SearchArticle) magic.BackData {
	// 初始化DB
	db := magic.Orm.Model(&community.Article{})

	// 筛选条件

	// 查询总数
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return magic.Back(2900, err.Error(), err.Error())
	}

	// 查询列表数据
	limit := data.PageSize
	offset := (data.Page - 1) * data.PageSize
	var list []community.Article
	err = db.Omit("content").Limit(limit).Offset(offset).Order("id DESC").Find(&list).Error
	if err != nil {
		return magic.Back(2902, err.Error(), err.Error())
	}

	// 组合返回数据
	res := model.ResPageData{
		List:     list,
		Total:    total,
		Page:     data.Page,
		PageSize: data.PageSize,
	}

	return magic.Back(0, "查询文章列表成功", res)
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

// 查找文章
func (a *ArticleService) Find(model.GetById) magic.BackData {

	//
	res := 0
	return magic.Back(0, "根据ID查找文章成功过", res)
}

// 编辑文章
func (a *ArticleService) Update() {

}

// 删除文章
func (a *ArticleService) Delete() {

}
