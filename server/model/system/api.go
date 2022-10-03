package system

import "github.com/jianyuezhexue/MagicAdmin/model"

// Api属性
type Api struct {
	model.BaseOrm
	Type  string `json:"type"`  // 请求类型
	Name  string `json:"name"`  // 中文名
	Route string `json:"route"` // 路由地址
}
