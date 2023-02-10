package pack

import (
	"douyin/dal/db"
	"douyin/kitex_gen/message"
)

func Messages(messages []*db.Message) []*message.Message {
	rpcMessages := make([]*message.Message, len(messages))
	for i := 0; i < len(messages); i++ {
		var m = messages[i]
		creatAt := m.CreatedAt.String()
		rpcMessages[i] = &message.Message{
			Id:         int64(m.ID),
			ToUserId:   int64(m.ToUserId),
			FromUserId: int64(m.FromUserId),
			Content:    m.Content,
			CreateTime: &creatAt,
		}
	}
	return rpcMessages
}
