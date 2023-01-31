/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-01-31 00:49:20
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-01-31 00:54:09
 * @FilePath: /ByteCamp/dal/pack/relation_resp.go
 * @Description:
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */

package pack

import (
	relation "douyin/kitex_gen/relation"
)

func BuildRelationActionResponse(err error) *relation.DouyinRelationActionResponse {
	var msg string
	if err != nil {
		msg = err.Error()
		return &relation.DouyinRelationActionResponse{
			StatusCode: -1,
			StatusMsg:  &msg,
		}
	} else {
		msg = "success"
		return &relation.DouyinRelationActionResponse{
			StatusCode: 0,
			StatusMsg:  &msg,
		}
	}
}
