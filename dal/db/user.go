/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-01-19 11:23:37
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-01-19 14:22:39
 * @FilePath: /ByteCamp/dal/db/userHandler.go
 * @Description: 用户实体类及相关crud
 *
 * Copyright (c) 2023 by zy 953725892@qq.com, All Rights Reserved.
 */

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

func GetUsersByUserName(userName string) ([]*User, error) {
	userList := make([]*User, 0)
	err := DB.Where("name = ?", userName).Find(&userList).Error
	if err != nil {
		return nil, err
	}
	return userList, nil
}

func CreateUser(user *User) error {
	return DB.Create(user).Error
}
