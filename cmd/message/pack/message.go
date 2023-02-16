package pack

import (
	"douyin/dal/db"
	"douyin/kitex_gen/message"
)

func Messages(messages []*db.Message) []*message.Message {
	rpcMessages := make([]*message.Message, len(messages))
	for i := 0; i < len(messages); i++ {
		var m = messages[i]
		creatAt := m.CreatedAt.Format("2006-01-02 15:04:05")
		rpcMessages[i] = &message.Message{
			Id:         int64(m.ID),
			Content:    m.Content,
			CreateTime: &creatAt,
		}
	}
	return rpcMessages
}
