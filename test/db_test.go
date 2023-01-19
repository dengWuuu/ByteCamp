/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-01-19 14:13:42
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-01-19 14:18:30
 * @FilePath: /ByteCamp/test/db_test.go
 * @Description:
 *
 * Copyright (c) 2023 by zy 953725892@qq.com, All Rights Reserved.
 */
package test

import (
	"context"
	"douyin/dal/db"
	"fmt"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	db.Init()
	//测试redis
	db.Redis.Set(context.Background(), "test", "test", time.Hour)
	fmt.Printf("%v", db.Redis.Get(context.Background(), "test"))

	//迁移数据库
	db.DB.AutoMigrate(&db.User{}, &db.Video{}, &db.Favorite{}, &db.Follow{}, &db.Comment{})
}
