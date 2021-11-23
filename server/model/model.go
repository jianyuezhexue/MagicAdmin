package model

import (
	"time"

	"gorm.io/gorm"
)

// BaseOrm ORM默认字段
type BaseOrm struct {
	ID        uint           `gorm:"primarykey"`     // 主键ID
	CreatedAt time.Time      `json:"createdAt"`      // 创建时间
	UpdatedAt time.Time      `json:"updatedAt"`      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}

// PageInfo common input parameter structure
type PageInfo struct {
	Page     int `json:"page" form:"page"`         // 页码
	PageSize int `json:"pageSize" form:"pageSize"` // 每页大小
}
