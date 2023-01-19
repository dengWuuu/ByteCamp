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
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	VideoId   int
	UserId    int
	Content   string
	CreatTime time.Time
	Version   int
	Cancel    bool
}
