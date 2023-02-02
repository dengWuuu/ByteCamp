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
