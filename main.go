package main

import (
	"log"

	"github.com/groupchat/chat/chatserver"
	"github.com/groupchat/mq/nsq"
)

func main() {
	log.SetFlags(log.Llongfile | log.Ldate)

	//Init Message Queueing
	publisher, err := nsq.NewPublisher("10.255.13.17:4150")
	if err != nil {
		log.Println(err)
		return
	}

	subscriber := nsq.NewSubscriber("10.255.13.17:4161")

	//Init Chat Server
	server := chatserver.New(":8080", publisher, subscriber)
	server.Run()
}
