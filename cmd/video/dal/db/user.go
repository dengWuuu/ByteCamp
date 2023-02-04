package db

import (
	"gorm.io/gorm"
	"gorm.io/plugin/optimisticlock"
)

type User struct {
	gorm.Model
	Name           string
	Password       string
	FollowingCount int `gorm:"default:0" json:"following_count"`
	FollowerCount  int `gorm:"default:0" json:"follower_count"`
	Version        optimisticlock.Version
}

func GetUserById(userId int64) (*User, error) {
	user := new(User)
	err := DB.First(&user, userId).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
