package system

import "github.com/jianyuezhexue/MagicAdmin/model"

// 媒体库
type MediaFiles struct {
	model.BaseOrm
	Name string `json:"name" gorm:"comment:文件名"` // 文件名
	Url  string `json:"url" gorm:"comment:文件地址"` // 文件地址
	Tag  string `json:"tag" gorm:"comment:文件标签"` // 文件标签
	Key  string `json:"key" gorm:"comment:编号"`   // 文件编号
}
