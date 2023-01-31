/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-01-30 23:34:04
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-01-30 23:35:16
 * @FilePath: /ByteCamp/cmd/relation/service/relation_service.go
 * @Description:relation微服务service
 *
 * Copyright (c) 2023 by ${git_name_email}, All Rights Reserved.
 */
package service

import "context"

type RelationService struct {
	ctx context.Context
}

func NewRelationService(ctx context.Context) *RelationService {
	return &RelationService{
		ctx: ctx,
	}
}
