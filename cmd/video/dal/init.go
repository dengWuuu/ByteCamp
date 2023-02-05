package dal

import (
	"douyin/cmd/video/dal/db"
	"douyin/cmd/video/dal/oss"
)

func Init() {
	db.Init()
	oss.Init()
}
