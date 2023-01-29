package main

import (
	video "douyin/kitex_gen/video/videosrv"
	"log"
)

func main() {
	svr := video.NewServer(new(VideoSrvImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
