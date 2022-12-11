package models

// MultiPart 分片信息
type MultiPart struct {
	ID           int    `gorm:"column:id;primaryKey;not null;autoIncrement;comment:自增ID"`
	StorageUid   int    `gorm:"column:storage_uid;not null;comment:存储UID"`
	FileMd5      string `gorm:"column:file_md5;not null;comment:文件md5"`
	PartFileName string `gorm:"column:part_file_name;not null;comment:分片文件名称"`
	PartMd5      string `gorm:"column:part_md5;not null;comment:分片md5"`
	Status       int    `gorm:"column:status;not null;comment:状态信息"`
}
