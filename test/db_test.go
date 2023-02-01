/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-01-19 14:13:42
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-01 22:16:17
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
	db.Init("../config")
	//测试redis
	db.Redis.Set(context.Background(), "test", "test", time.Hour)
	fmt.Printf("%v", db.Redis.Get(context.Background(), "test"))
	users := make([]*db.User, 0)
	ids := make([]int64, 0)
	ids = append(ids, 1)
	ids = append(ids, 2)
	ids = append(ids, 3)
	db.DB.Where("id in ?", ids).Find(&users)
	fmt.Printf("%v", users)
	//迁移数据库
	db.DB.AutoMigrate(&db.Video{})
}
