package system

import (
	"github.com/jianyuezhexue/MagicAdmin/model"
)

// 数据字典目录表
type Dictionary struct {
	model.BaseOrm
	Pid   int    `json:"pid" form:"pid"`
	Name  string `json:"name" form:"name"`
	Value string `json:"value" form:"value"`
	Super int    `json:"super" form:"super"`
	Sort  uint   `json:"sort" form:"sort"`
	Desc  string `json:"desc" form:"desc"`
}

// 分页查询字典目录
type SearchDictionary struct {
	model.PageInfo
	Dictionary
}

// 数据字典详情表
type DictionaryDetail struct {
	model.BaseOrm
	Pid   int    `json:"pid" form:"pid"`
	Name  string `json:"name" form:"name"`
	Value string `json:"value" form:"value"`
	Super int    `json:"super" form:"super"`
	Sort  uint   `json:"sort" form:"sort"`
	Desc  string `json:"desc" form:"desc"`
}

// 分页查询字典详情
type SearchDictionaryDetail struct {
	model.PageInfo
	DictionaryDetail
}

// 根绝keys查找字典
type Keys struct {
	Keys string `json:"keys" form:"keys"`
}
