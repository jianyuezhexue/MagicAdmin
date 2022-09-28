package system

import (
	"github.com/jianyuezhexue/MagicAdmin/model"
)

// 数据字典目录表
type Dictionary struct {
	model.BaseOrm
	Pid   int    `json:"pid"`
	Name  string `json:"name"`
	Value string `json:"value"`
	Super int    `json:"super"`
	Sort  uint   `json:"sort"`
	Desc  string `json:"desc"`
}

// 分页查询
type SearchInfo struct {
	model.PageInfo
	Dictionary
}
