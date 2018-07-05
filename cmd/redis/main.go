package main

import (
	"log"

	"github.com/groupchat/chat"
	"github.com/groupchat/chat/handler"
	redis "github.com/groupchat/mq/redis"
)

func main() {
	log.SetFlags(log.Llongfile | log.Ldate)

	//Init Message Queueing
	publisher, err := redis.NewPublisher("localhost:6379")
	if err != nil {
		log.Println(err)
		return
	}

	subscriber := redis.NewSubscriber("localhost:6379")

	//Init Chat Server
	server := chat.NewServer(publisher, subscriber)

	//Init Chat Server Handler
	handler := handler.New(server, ":6000")
	handler.Run()
}
