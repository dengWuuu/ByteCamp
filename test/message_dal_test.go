package test

import (
	"fmt"
	"testing"

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
