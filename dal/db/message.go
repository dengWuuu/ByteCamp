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

func GetUserMessageChat(fromUserId uint, toUserId uint) (messages []*Message, err error) {
	err = DB.Where("from_user_id = ? and to_user_id = ? ", fromUserId, toUserId).Order("created_at").Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func CreateMessage(message *Message) (err error) {
	err = DB.Create(message).Error
	if err != nil {
		return err
	}
	return nil
}
