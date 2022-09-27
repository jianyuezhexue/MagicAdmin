package system

import (
	"github.com/jianyuezhexue/MagicAdmin/model"
)

// 数据字典目录表
type Dictionary struct {
	model.BaseOrm
	Pid   uint   `json:"pid"`
	Name  string `json:"name"`
	Value string `json:"value"`
	Super uint   `json:"super"`
	Sort  uint   `json:"sort"`
	Desc  string `json:"desc"`
}
