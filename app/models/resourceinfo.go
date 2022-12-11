package models

import "time"

/*
ResourceInfo 表结构定义及增删改查接口
*/

// ResourceInfo 资源总表
type ResourceInfo struct {
	UID                 string     `gorm:"column:UID;primaryKey;not null;comment:唯一ID"`
	Type                string     `gorm:"column:Type;not null;comment:资源类型"`
	Bucket              string     `gorm:"column:Bucket;not null;comment:桶"`
	Name                string     `gorm:"column:Name;not null;comment:原始名称"`
	StorageName         string     `gorm:"column:StorageName;not null;comment:存储名称"`
	Address             string     `gorm:"column:Address;not null;comment:存储地址"`
	CompressName        *string    `gorm:"column:CompressName;comment:压缩存储名称"`
	CompressStorageName *string    `gorm:"column:CompressStorageName;comment:压缩存储名称"`
	CompressAddress     *string    `gorm:"column:CompressAddress;comment:压缩存储地址"`
	CreatedAt           *time.Time `gorm:"column:CreatedAt;not null;comment:创建时间"`
	UpdatedAt           *time.Time `gorm:"column:UpdatedAt;not null;comment:更新时间"`
}

// NewResourceInfo 获取新对象
func NewResourceInfo() *ResourceInfo {
	return &ResourceInfo{}
}
