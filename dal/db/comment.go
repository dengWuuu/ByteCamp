/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-01-19 11:44:09
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-01-19 14:22:06
 * @FilePath: /ByteCamp/dal/db/comment.go
 * @Description: 评论实体类及相关crud
 *
 * Copyright (c) 2023 by zy 953725892@qq.com, All Rights Reserved.
 */

package db

import (
	"context"
	"errors"
	"time"

	"gorm.io/plugin/optimisticlock"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	VideoId   int
	UserId    int
	Content   string
	CreatTime time.Time
	Version   optimisticlock.Version
	Cancel    bool
}

// 根据视频的ID来查询到相对应的评论，返回评论列表和信息
func GetCommentByVideoId(ctx context.Context, videoId int64) ([]*Comment, error) {
	var comments []*Comment
	err := DB.WithContext(ctx).Model(&Comment{}).Where(&Comment{VideoId: int(videoId)}).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// 根据对应的评论ID获取对应的评论列表
func GetCommentByCommentId(ctx context.Context, commentIds []int) ([]*Comment, error) {
	var comments []*Comment
	if len(commentIds) == 0 {
		return comments, nil
	}
	if err := DB.WithContext(ctx).Where("id in ?", commentIds).Find(&comments).Error; err != nil {
		return comments, err
	}
	return comments, nil
}

// 根据评论ID修改评论内容
func UpdateComment(ctx context.Context, commentId int, content string) error {
	return DB.WithContext(ctx).Model(&Comment{}).Where("id = ?", commentId).Update("content", content).Error
}

// 模糊查询（返回对应的数据，数量，状态信息）
func QueryComment(ctx context.Context, videoId int, searchKey *string, limit, offset int) ([]*Comment, int64, error) {
	var total int64
	var comments []*Comment
	res := DB.WithContext(ctx).Model(&Comment{}).Where("video_id = ?", videoId)
	if searchKey != nil {
		res = res.Where("content like ?", "%"+*searchKey+"%")
	}
	if err := res.Count(&total).Error; err != nil {
		return comments, total, err
	}
	if err := res.Limit(limit).Offset(offset).Find(&comments).Error; err != nil {
		return comments, total, err
	}
	return comments, total, nil
}

// 创建一个新的评论，同样的需要在视频增加评论的数目
func CreateComment(ctx context.Context, comment *Comment) error {
	// 需要在事务中进行处理
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		// 增加评论
		if res := tx.Create(comment).Error; res != nil {
			return res
		}
		// 修改对应的video的评论数
		// comment_count = comment_count + 1
		result := tx.Model(&Video{}).Where("ID = ?", comment.VideoId).Update("comment_count", gorm.Expr("comment_count + ?", 1))
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected > 1 {
			// 存在两个相同ID的视频
			return errors.New("数据库出错")
		}
		return nil
	})
	return err
}

// 删除评论，因此还要在对应的视频将评论数减一（根据ID删除）
func DeleteCommentById(ctx context.Context, videoId, commentId int) error {
	// 需要在事务中处理
	err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		// 减少评论
		if res := tx.Delete(&Comment{}, commentId).Error; res != nil {
			return res
		}
		// 修改对应的video的评论数
		// comment_count = comment_count - 1
		result := tx.Model(&Video{}).Where("ID = ?", videoId).Update("comment_count", gorm.Expr("comment_count - ?", 1))
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected > 1 {
			// 存在两个相同ID的视频
			return errors.New("数据库出错")
		}
		return nil
	})
	return err
}
