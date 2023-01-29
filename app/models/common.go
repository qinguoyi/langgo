package models

import (
	"time"
)

// ID 自增ID主键
type ID struct {
	ID uint `json:"id" gorm:"column:id;AUTO_INCREMENT"`
}

// Timestamps 创建、更新时间
type Timestamps struct {
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}

// SoftDeletes 软删除
type SoftDeletes struct {
	Status int `json:"status" gorm:"column:status;index"`
}
