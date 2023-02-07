package FavoriteMq

import (
	"context"
	"douyin/dal/db"
	"douyin/kitex_gen/favorite"
	"encoding/json"

	"github.com/u2takey/go-utils/klog"
)

func FavoriteConsumer() {
	_, err := favoriteMq.Channel.QueueDeclare(favoriteMq.QueueName, true, false, false, false, nil)
	if err != nil {
		klog.Fatalf("favorite add consumer declare error")
		panic(err)
	}

	//2、接收消息
	msgChanel, err := favoriteMq.Channel.Consume(
		favoriteMq.QueueName,
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
		FavoriteAction(string(msg.Body))
		// 主动应答
		// TODO 主动应答会出现问题
		// err := msg.Ack(true)
		// if err != nil {
		// 	klog.Info("ack失败")
		// 	return
		// }

	}
}

func FavoriteAction(msg string) {
	var req *favorite.DouyinFavoriteActionRequest
	err := json.Unmarshal([]byte(msg), req)
	if err != nil {
		klog.Fatalf("favoriteMq序列化消费信息失败")
		return
	}
	// 点赞
	if req.ActionType == 1 {
		err = db.AddFavorite(context.Background(), req.UserId, req.VideoId)
		if err != nil {
			klog.Fatalf("favoriteMq添加点赞关系失败")
			return
		}
	}
	if req.ActionType == 2 {
		err = db.DeleteFavorite(context.Background(), req.UserId, req.VideoId)
		if err != nil {
			klog.Fatalf("favoriteMq取消点赞关系失败")
			return
		}
	}
}
