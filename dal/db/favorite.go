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

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	UserId  int
	VideoId int
	Version int
	Cancel  bool
}
