package model

import (
	"time"
)

type User struct {
	ID           int64     `gorm:"primary_key"`          // 自增主键
	UserName     string    `gorm:"column:user_name"`     // 用户名
	UserID       string    `gorm:"column:user_id"`       // 用户身份证
	Password     string    `gorm:"column:password"`      // 登录密码
	RegisterTime time.Time `gorm:"column:register_time"` // 注册时间
	LastLogin    time.Time `gorm:"column:last_login"`    // 上次登录时间
}

func (*User) TableName() string {
	return "User"
}

func (user *User) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":            user.ID,
		"user_name":     user.UserName,
		"user_id":       user.UserID,
		"password":      user.Password,
		"register_time": user.RegisterTime,
		"last_login":    user.LastLogin,
	}
}
