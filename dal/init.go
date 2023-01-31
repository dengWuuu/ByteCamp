/*
 * @Author: zy 953725892@qq.com
 * @Date: 2023-01-19 10:01:15
 * @LastEditors: zy 953725892@qq.com
 * @LastEditTime: 2023-01-19 10:01:30
 * @FilePath: /ByteCamp/dal/init.go
 * @Description: 用于初始化数据库连接
 *
 * Copyright (c) 2023 by zy 953725892@qq.com, All Rights Reserved.
 */

package dal

import "douyin/dal/db"

func Init() {
	db.Init("../../config")
}
