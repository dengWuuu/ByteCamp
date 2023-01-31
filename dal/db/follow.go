/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-01-19 11:42:43
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-01-31 12:17:18
 * @FilePath: /ByteCamp/dal/db/follow.go
 * @Description: 关注实体类及相关crud
 *
 * Copyright (c) 2023 by zy 953725892@qq.com, All Rights Reserved.
 */

package db

import (
	"errors"

	"gorm.io/gorm"
	"gorm.io/plugin/optimisticlock"
)

type Follow struct {
	gorm.Model
	UserId   int
	FollowId int
	Version  optimisticlock.Version
	Cancel   bool
}

//添加关注关系
func AddRelation(userId, followId int) error {
	follow := Follow{
		UserId:   userId,
		FollowId: followId,
		Cancel:   false,
	}
	err := DB.Create(&follow).Error
	return err
}

//取消关注关系
func DeleteRelation(userId, followId int) error {
	//查询是否存在关注关系
	var follow Follow
	result := DB.Where("user_id = ? and follow_id = ?", userId, followId).First(&follow)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("关注关系不存在")
	}
	//更新关注关系
	err := DB.Model(&follow).Update("cancel", true).Error
	return err
}

//根据用户id查询关注列表
func GetFollowedByUserId(userId int) ([]Follow, error) {
	var follows []Follow
	err := DB.Where("user_id = ? and cancel = ?", userId, false).Find(&follows).Error
	return follows, err
}

//根据用户id查询粉丝列表
func GetFansByUserId(userId int) ([]Follow, error) {
	var follows []Follow
	err := DB.Where("follow_id = ? and cancel = ?", userId, false).Find(&follows).Error
	return follows, err
}
