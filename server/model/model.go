package model

import (
	"database/sql/driver"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// BaseOrm ORM默认字段
type BaseOrm struct {
	Id        uint           `json:"id" gorm:"primarykey"` // 主键ID
	CreatedAt LocalTime      `json:"-"`                    // 创建时间
	UpdatedAt LocalTime      `json:"updatedAt"`            // 更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;"`      // 删除时间
}

// LocalTime 格式化时间
type LocalTime struct {
	time.Time
}

// MarshalJSON 格式时间格式
func (t LocalTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

// Value Value
func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan Scan
func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// 分页查询参数结构
type PageInfo struct {
	Page     int `json:"page" form:"page"`         // 页码
	PageSize int `json:"pageSize" form:"pageSize"` // 每页大小
}

// 分页查询返回结构
type ResPageData struct {
	List  interface{} `json:"list"`
	Total int         `json:"total"`
	PageInfo
}

// 根据id查询接参结构
type GetById struct {
	ID int `json:"id" uri:"id" form:"id" binding:"required"` // 主键ID
}
