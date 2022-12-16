package model

import (
	"database/sql/driver"
	"strings"
	"time"

	"gorm.io/gorm"
)

// BaseOrm ORM默认字段
type BaseOrm struct {
	Id        uint64         `json:"id" gorm:"primarykey"` // 主键ID
	CreatedAt time.Time      `json:"createdAt"`            // 创建时间
	UpdatedAt time.Time      `json:"-"`                    // 更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;"`      // 删除时间
}

// 逗号分割的字符串类型
type Strs []string

func (m *Strs) Scan(val interface{}) error {
	s := val.([]uint8)
	ss := strings.Split(string(s), ",")
	*m = ss
	return nil
}
func (m Strs) Value() (driver.Value, error) {
	str := strings.Join(m, ",")
	return str, nil
}

// 分页查询参数结构
type PageInfo struct {
	Page     int `json:"page" form:"page"`         // 页码
	PageSize int `json:"pageSize" form:"pageSize"` // 每页大小
}

// 分页查询返回结构
type ResPageData struct {
	List     any `json:"list"`
	Total    any `json:"total"`
	Page     int `json:"page" form:"page"`         // 页码
	PageSize int `json:"pageSize" form:"pageSize"` // 每页大小
}

// 根据id查询接参结构
type GetById struct {
	Id int `json:"id" uri:"id" form:"id" binding:"required"` // 主键ID
}

// 传ID数组接收餐结构
type IdArr struct {
	Data []string `josn:"data" uri:"data" form:"data" binding:"required"`
}
