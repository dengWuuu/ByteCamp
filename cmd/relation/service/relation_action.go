package service

import (
	"douyin/dal/db"
	"douyin/kitex_gen/relation"
	"errors"
)

//执行用户对其他用户的关注或取消关注
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
