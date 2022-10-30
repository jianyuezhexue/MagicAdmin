package system

import (
	"github.com/jianyuezhexue/MagicAdmin/model"
)

type Record struct {
	model.BaseOrm
	UserId   int     `json:"userId"`   // 用户id
	UserName string  `json:"userName"` // 用户名称
	Ip       string  `json:"ip"`       // 请求ip
	Method   string  `json:"method"`   // 请求方法
	Path     string  `json:"path"`     // 请求路径
	Status   int     `json:"status"`   // 请求状态
	CostTime float64 `json:"costTime"` // 耗时
	Agent    string  `json:"agent"`    // 代理
	Msg      string  `json:"msg"`      // 错误信息
	Params   string  `json:"params"`   // 请求参数
	Resp     string  `json:"resp"`     // 响应Body
}

// 分页查询字典目录
type SearchRecord struct {
	model.PageInfo
	Record
}
