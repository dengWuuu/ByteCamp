/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-01-19 14:13:42
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-02 22:41:18
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
	ids, err := db.GetFriendsByUserId(23)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v", ids)
	//迁移数据库
	// db.DB.AutoMigrate(&db.Video{})
}

func TestRedis(t *testing.T) {
	db.Init("../config")
	ctx := context.Background()
	db.FollowersRedis.Set(ctx, "test1", "1", time.Minute*10)
	db.FollowingRedis.Set(ctx, "test2", "2", time.Minute*10)
	db.FriendsRedis.Set(ctx, "test3", "3", time.Minute*10)
	var res string
	var err error

	res, err = db.FollowingRedis.Get(ctx, "test1").Result()

	res, err = db.FollowingRedis.Get(ctx, "test2").Result()

	res, err = db.FollowingRedis.Get(ctx, "test3").Result()

	res, err = db.FollowersRedis.Get(ctx, "test1").Result()

	res, err = db.FollowersRedis.Get(ctx, "test2").Result()

	res, err = db.FollowersRedis.Get(ctx, "test3").Result()

	res, err = db.FriendsRedis.Get(ctx, "test1").Result()

	res, err = db.FriendsRedis.Get(ctx, "test2").Result()

	res, err = db.FriendsRedis.Get(ctx, "test3").Result()

	fmt.Printf("%v,%v", res, err)
}
