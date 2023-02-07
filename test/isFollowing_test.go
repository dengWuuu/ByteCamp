/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-02-06 14:12:54
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-06 14:25:03
 * @FilePath: \ByteCamp\test\isFollowing_test.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package test

import (
	"context"
	"testing"

	"douyin/dal/db"
	"douyin/pkg/redis"
)

func TestIsFollowing(t *testing.T) {
	db.Init("../config")
	ctx := context.Background()
	res, _ := redis.IsFollowing(ctx, 1, 2)
	t.Log(res)
	res, _ = redis.IsFollowing(ctx, 1, 5)
	t.Log(res)
	res, _ = redis.IsFollowing(ctx, 1, 10)
	t.Log(res)
	res, _ = redis.IsFollowing(ctx, 23, 24)
	t.Log(res)
}
