package main

import (
	user "douyin/kitex_gen/user/usersrv"
	"log"
)

func main() {
	svr := user.NewServer(new(UserSrvImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
