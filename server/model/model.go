package model

import (
	"time"

	"gorm.io/gorm"
)

// BaseOrm ORM默认字段
type BaseOrm struct {
	ID        uint           `gorm:"primarykey"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}
