/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-01-19 11:42:43
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-01-19 14:22:34
 * @FilePath: /ByteCamp/dal/db/follow.go
 * @Description: 关注实体类及相关crud
 *
 * Copyright (c) 2023 by zy 953725892@qq.com, All Rights Reserved.
 */

package db

import (
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
