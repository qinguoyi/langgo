package models

import "langgo/app/pkg/sqls"

// User .
type User struct {
	ID
	UUID     string `json:"uuid" gorm:"uuid;column:uuid;comment:唯一标识"`
	Name     string `json:"name" gorm:"not null;column:name;comment:用户名"`
	Mobile   string `json:"mobile" gorm:"not null;column:mobile;comment:用户手机号"`
	Password string `json:"password" gorm:"not null;default:'';column:password;comment:用户密码"`
	Timestamps
	SoftDeletes
}

// CreateUser .
type CreateUser struct {
	Name     string `json:"name" binding:"required"`
	Mobile   string `json:"mobile" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// QueryUser .
type QueryUser struct {
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

// QueryUserPage .
type QueryUserPage struct {
	Data []QueryUser    `json:"data"`
	Page *sqls.PageInfo `json:"page"`
}

type UpdateUser struct {
	Mobile   string `json:"mobile" binding:"required"`
	Password string `json:"password" binding:"required"`
}
