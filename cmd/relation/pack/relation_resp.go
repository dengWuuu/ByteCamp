/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-01-31 00:49:20
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-01 02:08:39
 * @FilePath: /ByteCamp/cmd/relation/pack/relation_resp.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */

package pack

import (
	relation "douyin/kitex_gen/relation"
	"douyin/pkg/errno"
	"errors"
)

//根据error，判断其属于那种errno.ErrNo,调用GetRelationActionResp返回对应的relation.DouyinRelationActionResponse
func BuildRelationActionResponse(err error) *relation.DouyinRelationActionResponse {
	if err == nil {
		return getRelationActionResp(errno.Success)
	} else {
		e := errno.ErrNo{}
		if errors.As(err, &e) {
			return getRelationActionResp(e)
		}

		s := errno.ErrUnknown.WithMessage(err.Error())
		return getRelationActionResp(s)
	}

}

//根据传入的errno.ErrNo，返回对应的relation.DouyinRelationActionResponse
func getRelationActionResp(err errno.ErrNo) *relation.DouyinRelationActionResponse {
	return &relation.DouyinRelationActionResponse{
		StatusCode: int32(err.ErrCode),
		StatusMsg:  &err.ErrMsg,
	}
}
