package db

import (
	"gorm.io/gorm"
	"gorm.io/plugin/optimisticlock"
)

type Message struct {
	gorm.Model
	FromUserId uint
	ToUserId   uint
	Content    string
	Version    optimisticlock.Version
}

func GetUserMessageChat(toUserId uint) (messages []*Message, err error) {
	err = DB.Where("to_user_id = ? ", toUserId).Order("created_at").Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func MessageAction(message Message) {

}
