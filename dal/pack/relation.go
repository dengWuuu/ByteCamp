package pack

import (
	"douyin/dal/db"
	"douyin/kitex_gen/user"
)

//根据follows列表，获取所有关注用户的rpc格式信息
func GetFollowingByFollows(follows []db.Follow) ([]*user.User, error) {

}

//根据follows列表，获取所有粉丝用户的rpc格式信息
func GetFansByFollows(follows []db.Follow) ([]*user.User, error) {

}
