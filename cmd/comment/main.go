package main

import (
	comment "douyin/kitex_gen/comment/commentsrv"
	"log"
)

func main() {
	svr := comment.NewServer(new(CommentSrvImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
