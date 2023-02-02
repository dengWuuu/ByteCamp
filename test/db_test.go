/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-01-19 14:13:42
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-02-02 18:15:43
 * @FilePath: /ByteCamp/test/db_test.go
 * @Description:
 *
 * Copyright (c) 2023 by zy 953725892@qq.com, All Rights Reserved.
 */
package test

import (
	"douyin/dal/db"
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	db.Init("../config")
	//测试redis
	ids, err := db.GetFriendsByUserId(23)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v", ids)
	//迁移数据库
	// db.DB.AutoMigrate(&db.Video{})
}
