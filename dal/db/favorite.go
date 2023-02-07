/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-01-19 11:40:55
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-01-19 14:22:30
 * @FilePath: /ByteCamp/dal/db/favorite.go
 * @Description: 点赞实体类及相关crud
 *
 * Copyright (c) 2023 by zy 953725892@qq.com, All Rights Reserved.
 */

package db

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"gorm.io/plugin/optimisticlock"
)

type Favorite struct {
	gorm.Model
	UserId  int
	VideoId int
	Version optimisticlock.Version
	Cancel  bool
}

// 获取用户ID获取所有的点赞视频
func GetFavoritesByUserId(user_id int64) (resp []*Favorite, err error) {
	err = DB.Where("user_id = ? and cancel = ", user_id, false).Find(&resp).Error
	return resp, err
}

// 点赞操作，同时要将点赞的视频点赞数量加一
func AddFavorite(ctx context.Context, user_id, video_id int64) error {
	// 需要在事务里面进行操作
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 检查用户
		if res := tx.WithContext(ctx).First(&User{}, user_id).Error; res != nil {
			return res
		}
		// 检查视频
		if res := tx.WithContext(ctx).First(&Video{}, video_id).Error; res != nil {
			return res
		}
		// 检查是否存在
		var temp Favorite
		duplicate := tx.Where("user_id = ? and video_id = ?", user_id, video_id).First(&temp)
		if duplicate.RowsAffected > 0 {
			return errors.New("点赞关系重复出现")
		}
		favorite_relation := &Favorite{
			UserId:  int(user_id),
			VideoId: int(video_id),
			Cancel:  false,
		}
		// 创建新的联系
		if res := tx.WithContext(ctx).Create(favorite_relation).Error; res != nil {
			return res
		}
		// 同时增加视频的点赞数量
		result := tx.Model(&Video{}).Where("ID = ?", video_id).Update("favorite_count", gorm.Expr("favorite_count + ?", 1))
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected > 1 {
			// 数据库存在相同的两个视频
			return errors.New("数据库出错")
		}
		return nil
	})
	return err
}

// 取消点赞操作，同时需要将取消点赞的视频点赞数量减一
func DeleteFavorite(ctx context.Context, user_id, video_id int64) error {
	// 需要在事务里面进行操作
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 检查用户
		if res := tx.WithContext(ctx).First(&User{}, user_id).Error; res != nil {
			return res
		}
		// 检查视频
		if res := tx.WithContext(ctx).First(&Video{}, video_id).Error; res != nil {
			return res
		}
		// 检查是否存在
		var temp Favorite
		duplicate := tx.Where("user_id = ? and video_id = ?", user_id, video_id).First(&temp)
		if duplicate.RowsAffected == 0 {
			return errors.New("不存在点赞关系")
		}
		// 删除联系
		if res := tx.Model(&temp).Update("cancel", true).Error; res != nil {
			return res
		}
		// 同时减少视频的点赞数量
		result := tx.Model(&Video{}).Where("ID = ?", video_id).Update("favorite_count", gorm.Expr("favorite_count - ?", 1))
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected > 1 {
			// 数据库存在相同的两个视频
			return errors.New("数据库出错")
		}
		return nil
	})
	return err
}
