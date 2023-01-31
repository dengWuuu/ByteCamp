/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-01-31 14:46:35
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-01-31 14:47:06
 * @FilePath: /ByteCamp/pack/relation.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package pack

import (
	"douyin/dal/db"
	"douyin/kitex_gen/user"
)

//根据follows列表，获取所有关注用户的rpc格式信息
func GetFollowingByFollows(follows []db.Follow) ([]*user.User, error) {
	return nil, nil
}

//根据follows列表，获取所有粉丝用户的rpc格式信息
func GetFansByFollows(follows []db.Follow) ([]*user.User, error) {
	return nil, nil
}
