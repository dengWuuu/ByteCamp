/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-01-19 11:23:37
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-01 16:37:33
 * @FilePath: /ByteCamp/dal/db/user.go
 * @Description: 用户实体类及相关crud
 *
 * Copyright (c) 2023 by zy 953725892@qq.com, All Rights Reserved.
 */

package db

import (
	"douyin/pkg/bcrypt"
	"douyin/pkg/errno"

	"github.com/cloudwego/hertz/pkg/common/hlog"
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

func GetUserById(userId int64) (*User, error) {
	user := new(User)
	err := DB.First(&user, userId).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUser(user *User) error {
	return DB.Create(user).Error
}

func GetUsersByIds(userIds []int64) ([]*User, error) {
	users := make([]*User, 0)
	err := DB.Where("id in ?", userIds).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// CheckUser 检验用户登录信息是否正确
func CheckUser(username string, password string) ([]*User, error) {
	//首先加密密码然后进行比对
	p, err := bcrypt.PasswordHash(password)
	if err != nil {
		hlog.Fatalf("checkUser时加密失败")
		return nil, err
	}
	users, err := GetUsersByUserName(username)
	if err != nil {
		hlog.Fatalf("根据用户名查找用户信息失败")
		return nil, err
	}
	if len(users) == 0 {
		return nil, errno.ErrUserNotFound
	}
	user := users[0]

	passwordMatch := bcrypt.PasswordVerify(p, user.Password)
	if !passwordMatch {
		return nil, errno.ErrPasswordIncorrect
	}
	return users, nil
}
