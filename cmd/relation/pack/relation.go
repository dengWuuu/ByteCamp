/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-01-31 14:46:35
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-03 23:33:14
 * @FilePath: /ByteCamp/cmd/relation/pack/relation.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package pack

import (
	"context"
	userpack "douyin/cmd/user/pack"
	"douyin/dal/db"
	"douyin/kitex_gen/user"
)

func GetUsersByIds(ids []int64) ([]*user.User, error) {
	dbusers, err := db.GetUsersByIds(ids)
	if err != nil {
		return nil, err
	}
	users, err := userpack.Users(context.Background(), dbusers, 0)
	if err != nil {
		return nil, err
	}
	return users, nil
}
