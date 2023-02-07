package FavoriteMq

import (
	"douyin/pkg/rabbitmq"
)

var favoriteMq *rabbitmq.RabbitMQ

func InitFavoriteMq() {
	rabbitmq.InitRabbitMQ()
	favoriteMq = rabbitmq.NewRabbitMq("favorite_queue", "favorite_exchange", "favorite")
}
