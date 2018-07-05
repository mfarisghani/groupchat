package main

import (
	"log"

	"github.com/groupchat/chat/chatserver"
	"github.com/groupchat/mq/nsq"
)

func main() {
	//Init Message Queueing
	nsqPublisher, err := nsq.NewPublisher("devel-go.tkpd:4150")
	if err != nil {
		log.Println(err)
		return
	}

	nsqSubscriber := nsq.NewSubscriber("devel-go.tkpd:4161")

	//Init Chat Server
	server := chatserver.New(":8080", nsqPublisher, nsqSubscriber)
	server.Run()
}
