package system

import (
	"github.com/jianyuezhexue/MagicAdmin/model"
)

type Record struct {
	model.BaseOrm
	Ip       string  `json:"ip" gorm:"-"` // 请求ip
	Path     string  `json:"path"`        // 请求路径
	Method   string  `json:"method"`      // 请求方法
	Status   int     `json:"status"`      // 请求状态
	CostTime float64 `json:"costTime"`    // 耗时
	Agent    string  `json:"agent"`       // 代理
	Msg      string  `json:"msg"`         // 错误信息
	Params   string  `json:"params"`      // 请求参数
	Resp     string  `json:"resp"`        // 响应Body
	UserId   int     `json:"userId"`      // 用户id
	UserName string  `json:"userName"`    // 用户名称
	// User         User          `json:"user"`
}
