package main

import (
	favorite "douyin/kitex_gen/favorite/favoritesrv"
	"log"
)

func main() {
	svr := favorite.NewServer(new(FavoriteSrvImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
