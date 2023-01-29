/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-01-19 11:27:35
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-01-19 14:22:42
 * @FilePath: /ByteCamp/dal/db/video.go
 * @Description: 视频实体类及相关crud
 *
 * Copyright (c) 2023 by zy 953725892@qq.com, All Rights Reserved.
 */

package db

import (
	"time"

	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	AuthorId   int
	PlayUrl    string
	CoverUrl   string
	UploadTime time.Time
	Title      string
	Version    int
}
