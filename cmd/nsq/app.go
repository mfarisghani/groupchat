package main

import (
	"log"

	"github.com/groupchat/chat"
	"github.com/groupchat/chat/handler"
	"github.com/groupchat/mq/nsq"
)

func main() {
	log.SetFlags(log.Llongfile | log.Ldate)

	//Init Message Queueing
	publisher, err := nsq.NewPublisher("devel-go.tkpd:4150")
	if err != nil {
		log.Println(err)
		return
	}

	subscriber := nsq.NewSubscriber("devel-go.tkpd:4161")

	//Init Chat Server
	server := chat.NewServer(publisher, subscriber)

	//Init Chat Server Handler
	handler := handler.New(server, ":8080")
	handler.Run()
}
