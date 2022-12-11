package models

import "time"

// StorageFileModel 存储文件
type StorageFileModel struct {
	ID             int        `gorm:"column:id;primaryKey;not null;comment:自增ID;autoIncrement"`
	UID            int        `gorm:"column:uid;not null;comment:唯一标识"`
	SourcePath     string     `gorm:"column:source_path;not null;comment:文件路径"`
	StorageUid     string     `gorm:"column:storage_path;not null;comment:存储UID"`
	Md5            string     `gorm:"column:md5;not null;comment:md5"`
	Path           string     `gorm:"column:path;not null;comment:路径"`
	Height         int        `gorm:"column:height;comment:高度"`
	Width          int        `gorm:"column:width;comment:宽度"`
	Address        string     `gorm:"column:address;not null;comment:地址"`
	StorageSize    int        `gorm:"column:storage_size;not null;comment:文件大小"`
	UserID         string     `gorm:"column:user_id;not null;comment:用户ID"`
	OrganizationID string     `gorm:"column:organization_id;not null;comment:组织ID"`
	CreatedAt      *time.Time `gorm:"column:created_at;not null;comment:创建时间"`
	UpdatedAt      *time.Time `gorm:"column:UpdatedAt;not null;comment:更新时间"`
	TaskStatus     int        `gorm:"column:task_status;not null;comment:上传确认信息"`
	Status         int        `gorm:"column:status;not null;comment:状态信息"`
}
