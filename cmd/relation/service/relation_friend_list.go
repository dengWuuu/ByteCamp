/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-02-02 18:43:44
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-03 23:34:28
 * @FilePath: /ByteCamp/cmd/relation/service/relation_friend_list.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-02-02 18:43:44
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-03 23:09:26
 * @FilePath: /ByteCamp/cmd/relation/service/relation_friend_list.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package service

import (
	"context"
	userpack "douyin/cmd/user/pack"
	"douyin/dal/db"
	"douyin/kitex_gen/relation"
	"douyin/kitex_gen/user"
)

//根据req获取RPC所需的朋友userId列表
func (service RelationService) FriendList(req *relation.DouyinRelationFriendListRequest) ([]*user.User, error) {
	//1、根据userId获取该user的所有follow列表
	ids, err := db.GetFriendsByUserId(int(req.UserId))
	if err != nil {
		return nil, err
	}
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
