package reids

import (
	"context"
	"douyin/dal/db"
	"gorm.io/gorm"
	"gorm.io/plugin/optimisticlock"
	"testing"
)

func TestGetUserFromRedis(t *testing.T) {
	db.Init("D:\\GolandProjects\\Douyin\\config")
	userId := []uint{1010}
	GetUsersFromRedis(context.Background(), userId)
}

func TestPutUserToRedis(t *testing.T) {
	db.Init("D:\\GolandProjects\\Douyin\\config")
	user := db.User{
		Model:          gorm.Model{ID: 1010},
		Name:           "wdw",
		Password:       "test",
		FollowingCount: 0,
		FollowerCount:  0,
		Version:        optimisticlock.Version{},
	}
	PutUserToRedis(context.Background(), &user)
}
