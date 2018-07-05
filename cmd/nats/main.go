package main

import (
	"log"

	"github.com/groupchat/chat"
	"github.com/groupchat/chat/handler"
	nats "github.com/groupchat/mq/nats"
)

func main() {
	log.SetFlags(log.Llongfile | log.Ldate)

	//Init Message Queueing
	publisher, err := nats.NewPublisher("nats://172.31.5.45:4222")
	if err != nil {
		log.Println(err)
		return
	}

	subscriber := nats.NewSubscriber("nats://172.31.5.45:4222")

	//Init Chat Server
	server := chat.NewServer(publisher, subscriber)

	//Init Chat Server Handler
	handler := handler.New(server, ":8080")
	handler.Run()
}
