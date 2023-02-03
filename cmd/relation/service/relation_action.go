/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-01-31 12:15:33
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-03 21:13:52
 * @FilePath: /ByteCamp/cmd/relation/service/relation_action.go
 * @Description: relationAction接口对应service
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package service

import (
	"douyin/dal/db"
	"douyin/kitex_gen/relation"
	"errors"
)

//TODO:采用redis维护用户的关注、粉丝、朋友列表
func (service *RelationService) RelationAction(req *relation.DouyinRelationActionRequest) error {
	if req.ActionType == 1 {
		err := db.AddRelation(int(req.UserId), int(req.ToUserId))
		if err != nil {
			return err
		}
	} else if req.ActionType == 2 {
		err := db.DeleteRelation(int(req.UserId), int(req.ToUserId))
		if err != nil {
			return err
		}
	} else {
		return errors.New("action_type error")
	}
	return nil
}
