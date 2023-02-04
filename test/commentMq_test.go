package test

import (
	"douyin/cmd/comment/commentMq"
	"douyin/dal/db"
	"douyin/kitex_gen/comment"
	"encoding/json"
	"fmt"
	"testing"
)

func TestCommentMq_Publish(t *testing.T) {
	var id int64 = 12
	for i := 1; i < 5; i++ {
		commentMq.InitCommentMq()
		var text = "kadhfkadshf"
		request := comment.DouyinCommentActionRequest{
			UserId:      0,
			Token:       "",
			VideoId:     1,
			ActionType:  2,
			CommentText: &text,
			CommentId:   &id,
		}
		marshal, err := json.Marshal(request)
		if err != nil {
			fmt.Println("序列化失败")
		}
		commentMq.CommentActionMqSend(marshal)
	}
}

func TestCommentMq_Consumer(t *testing.T) {

	db.Init("D:\\GolandProjects\\Douyin\\config")
	commentMq.InitCommentMq()
	commentMq.CommentConsumer()
}
