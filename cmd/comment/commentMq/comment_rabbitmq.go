package commentMq

import (
	"douyin/pkg/rabbitmq"
	"time"
)

var commentMq *rabbitmq.RabbitMQ

type CommentRmqMessage struct {
	VideoId    int
	UserId     int
	Content    string
	CreateTime time.Time
	CommentId  int
	ActionType int
}

func InitCommentMq() {
	rabbitmq.InitRabbitMQ()
	commentMq = rabbitmq.NewRabbitMq("comment_queue", "comment_exchange", "comment")
}
