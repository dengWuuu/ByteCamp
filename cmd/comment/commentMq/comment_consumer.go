package commentMq

import (
	"context"
	"douyin/dal/db"
	"encoding/json"

	"github.com/cloudwego/kitex/pkg/klog"
)

func CommentConsumer() {
	_, err := commentMq.Channel.QueueDeclare(commentMq.QueueName, true, false, false, false, nil)
	if err != nil {
		klog.Fatalf("comment add consumer declare error")
		panic(err)
	}

	//2、接收消息
	msgChanel, err := commentMq.Channel.Consume(
		commentMq.QueueName,
		//用来区分多个消费者
		"",
		//是否自动应答
		true,
		//是否具有排他性
		false,
		//如果设置为true，表示不能将同一个connection中发送的消息传递给这个connection中的消费者
		false,
		//消息队列是否阻塞
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	for msg := range msgChanel {
		// 这里写你的处理逻辑
		// 获取到的消息是amqp.Delivery对象，从中可以获取消息信息
		commentAction(string(msg.Body))
		// 主动应答
		// TODO 主动应答会出现问题
		// err := msg.Ack(true)
		// if err != nil {
		// 	klog.Info("ack失败")
		// 	return
		// }

	}
}

func commentAction(msg string) {
	var req *CommentRmqMessage
	err := json.Unmarshal([]byte(msg), &req)
	if err != nil {
		klog.Fatalf("rabbitMq commentAdd消费时序列化失败")
		return
	}
	// 根据请求创建新的评论
	if req.ActionType == 1 {
		commentModel := &db.Comment{
			UserId:  req.UserId,
			VideoId: req.VideoId,
			Content: req.Content,
		}
		// * 一定要记得加上ID和时间
		commentModel.CreatedAt = req.CreateTime
		commentModel.ID = uint(req.CommentId)
		err := db.CreateComment(context.Background(), commentModel)
		if err != nil {
			klog.Fatalf("rabbitmq 消费者在数据库中创建评论失败")
			panic(err)
		}
	}
	// 根据请求删除评论
	if req.ActionType == 2 {
		err := db.DeleteCommentById(context.Background(), req.VideoId, req.CommentId)
		if err != nil {
			klog.Fatalf("rabbitmq 消费者在数据库中删除评论失败")
			return
		}
	}
}
