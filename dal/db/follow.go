/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-01-19 11:42:43
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-02 18:25:09
 * @FilePath: /ByteCamp/dal/db/follow.go
 * @Description: 关注实体类及相关crud
 *
 * Copyright (c) 2023 by zy 953725892@qq.com, All Rights Reserved.
 */

package db

//TODO:为高频字段添加索引

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

//TODO:这里查询了follow的所有字段，而实际上我们只需要follow_id，这会导致无法索引覆盖，需要优化
//根据用户id查询关注列表
func GetFollowingByUserId(userId int) ([]*Follow, error) {
	var follows []*Follow
	err := DB.Where("user_id = ? and cancel = ?", userId, false).Find(&follows).Error
	return follows, err
}

//TODO:这里查询了follow的所有字段，而实际上我们只需要user_id，这会导致无法索引覆盖，需要优化
//根据用户id查询粉丝列表
func GetFansByUserId(userId int) ([]*Follow, error) {
	var follows []*Follow
	err := DB.Where("follow_id = ? and cancel = ?", userId, false).Find(&follows).Error
	return follows, err
}

func GetFriendsByUserId(userId int) ([]int64, error) {
	var friends []int64
	err := DB.Raw("select a.follow_id from follows as a INNER JOIN follows as b ON a.follow_id=b.user_id where a.user_id = b.follow_id and a.user_id = ?", userId).Scan(&friends).Error
	return friends, err
}
