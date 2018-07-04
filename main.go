package main

import (
	"log"

	"github.com/groupchat/chat/chatserver"
	"github.com/groupchat/mq/nsq"
)

func main() {
	//Init Message Queueing
	publisher, err := nsq.NewPublisher("172.31.0.58:4150")
	if err != nil {
		log.Println(err)
		return
	}

	subscriber := nsq.NewSubscriber("172.31.0.58:4161")

	//Init Chat Server
	server := chatserver.New(":8080", publisher, subscriber)
	server.Run()
}
