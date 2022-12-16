package community

import "github.com/jianyuezhexue/MagicAdmin/model"

type Article struct {
	model.BaseOrm
	Uuid         string `json:"uuid"`                                                  // 用户ID
	AuthName     string `json:"authName"`                                              // 文章作者
	Title        string `json:"title" form:"title" binding:"required,max=50"`          // 文章标题
	Summary      string `json:"summary" form:"summary" binding:"required,max=100"`     // 文章摘要
	Content      string `json:"content" form:"content" binding:"required,max=5000"`    // 文章内容
	Type         int64  `json:"type" form:"type" binding:"required,numeric,max=2"`     // 文章类型：1-图文；2-视频文章
	PicList      string `json:"picList" form:"picList" binding:"required,max=150"`     // 图片列表
	VideoList    string `json:"videoList" form:"videoList" binding:"required,max=150"` // 视频列表
	CategoryList string `json:"categoryList" form:"categoryList"`                      // 栏目IDs
	MarkList     string `json:"markList" form:"markList"`                              // 标志IDs
	TopicList    string `json:"topicList" form:"topicList"`                            // 话题IDs
	TagList      string `json:"tagList" form:"tagList"`                                // 标签IDs
	Status       string `json:"status"`                                                // 文章状态
	LikeNum      int64  `json:"likeNum"`                                               // 点赞数量
	CommontNum   int64  `json:"commontNum"`                                            // 评论数量
	HotNum       int64  `json:"hotNum"`                                                // 热力值
	IpAddress    string `json:"ipAddress"`                                             // IP所属省份
}
