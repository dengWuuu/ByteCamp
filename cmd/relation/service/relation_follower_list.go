package service

import (
	"douyin/cmd/relation/pack"
	"douyin/dal/db"
	"douyin/kitex_gen/relation"
	"douyin/kitex_gen/user"
)

//根据req获取RPC所需的粉丝user列表
func (service RelationService) FollowerList(req *relation.DouyinRelationFollowerListRequest) ([]*user.User, error) {
	//1、根据userId获取该user的所有follow列表
	fans, err := db.GetFansByUserId(int(req.UserId))
	if err != nil {
		return nil, err
	}
	fansUsers, err := pack.GetFansByFollows(fans)
	if err != nil {
		return nil, err
	}
	return fansUsers, nil
}
