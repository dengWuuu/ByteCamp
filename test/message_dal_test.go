package test

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/plugin/optimisticlock"
	"testing"
	"time"

	"douyin/dal/db"
)

func TestMessageChat(t *testing.T) {
	db.Init("D:\\GolandProjects\\Douyin\\config")
	chat, err := db.GetUserMessageChat(2)
	if err != nil {
		return
	}
	fmt.Println(chat)
}

func TestMessageAction(t *testing.T) {
	db.Init("D:\\GolandProjects\\Douyin\\config")
	err := db.CreateMessage(&db.Message{
		Model: gorm.Model{
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		FromUserId: 1,
		ToUserId:   2,
		Content:    "tesst",
		Version:    optimisticlock.Version{},
	})
	if err != nil {
		return
	}
}
