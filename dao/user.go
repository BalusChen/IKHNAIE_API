package dao

import (
	"context"
	"log"
	"time"

	"github.com/BalusChen/IKHNAIE_API/model"
	"github.com/jinzhu/gorm"
)

func RegisterUser(ctx context.Context, user *model.User) error {
	err := ikhnaieDB.Where("user_id = ?", user.UserID).Assign(user.ToMap()).FirstOrCreate(user).Error
	if err != nil {
		log.Printf("[RegisterUser] insert to db failed, err: %v\n", err)
		return err
	}
	return nil
}

func GetUserByUserID(ctx context.Context, userId string) *model.User {
	var user model.User
	err := ikhnaieDB.Where("user_id = ?", userId).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("[GetUserByUserID] user of userId=%q not found\n", userId)
		} else {
			log.Printf("[GetUserByUserID] select user of userId=%q failed, err: %v\n", userId, err)
		}

		return nil
	}
	return &user
}

func GetUserByUserName(ctx context.Context, userName string) *model.User {
	var user model.User
	err := ikhnaieDB.Where("user_name = ?", userName).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("[GetUserByUserName] user of userName=%q not found\n", userName)
		} else {
			log.Printf("[GetUserByUserName] select user of userName=%q failed, err: %v\n", userName, err)
		}

		return nil
	}
	return &user
}

func GetUser(ctx context.Context, userName, userId string) (bool, error) {
	var user model.User
	err := ikhnaieDB.Where("user_name = (?) OR user_id = (?)", userName, userId).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil
}

func GetUsersByStatus(ctx context.Context, status int32) ([]*model.User, error) {
	var users []*model.User
	err := ikhnaieDB.Where("status = (?)", status).Find(&users).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return users, nil
}

func UpdateUserStatus(ctx context.Context, userId string, status int32) error {
	var user model.User
	err := ikhnaieDB.Where("user_id = (?)", userId).Assign(map[string]interface{}{
		"status": status,
	}).First(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserLastLoginTime(ctx context.Context, userId string, t time.Time) error {
	var user model.User
	err := ikhnaieDB.Where("user_id = ?", userId).Assign(map[string]interface{}{
		"last_login": t,
	}).FirstOrCreate(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("[UpdateUserLastLogin] user of userId=%q not found", userId)
		} else {
			log.Printf("[UpdateUserLastLogin] failed, err: %v", err)
		}

		return err
	}
	return nil
}
