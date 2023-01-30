/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-01-19 11:23:37
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-01-19 14:22:39
 * @FilePath: /ByteCamp/dal/db/user.go
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
	Name     string
	Password string
	Version  optimisticlock.Version
}
