package models

import (
	"time"
)

// ID 自增ID主键
type ID struct {
	ID uint `json:"id" gorm:"primaryKey"`
}

// Timestamps 创建、更新时间
type Timestamps struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// SoftDeletes 软删除
type SoftDeletes struct {
	Status int `json:"status" gorm:"index"`
}
