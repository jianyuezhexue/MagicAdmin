package system

import "github.com/jianyuezhexue/MagicAdmin/model"

// JwtBlacklist JWT黑名单
type JwtBlacklist struct {
	model.BaseOrm
	Jwt string `gorm:"type:text;comment:jwt"`
}
