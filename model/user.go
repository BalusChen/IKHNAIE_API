package model

import (
	"time"
)

const (
	UserType_Admin  = 1
	UserType_Normal = 2

	UserStatus_Inactive = 1
	UserStatus_Active   = 2
)

type User struct {
	ID           int64     `gorm:"primary_key"`          // 自增主键
	Type         int32     `gorm:"column:type"`          // 用户类型：1-管理员，2-普通用户
	Status       int32     `gorm:"column:status"`        // 普通用户状态：1-未准入，2-已准入
	UserName     string    `gorm:"column:user_name"`     // 用户名
	UserID       string    `gorm:"column:user_id"`       // 用户身份证
	Password     string    `gorm:"column:password"`      // 登录密码
	Organization string    `gorm:"column:organization"`  // 普通用户所属组织
	RegisterTime time.Time `gorm:"column:register_time"` // 注册时间
	LastLogin    time.Time `gorm:"column:last_login"`    // 上次登录时间
}

func (*User) TableName() string {
	return "User"
}

func (user *User) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":            user.ID,
		"type":          user.Type,
		"status":        user.Status,
		"user_name":     user.UserName,
		"user_id":       user.UserID,
		"password":      user.Password,
		"organization":  user.Organization,
		"register_time": user.RegisterTime,
		"last_login":    user.LastLogin,
	}
}
