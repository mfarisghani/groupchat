package main

import (
	"log"

	"github.com/groupchat/chat"
	"github.com/groupchat/chat/handler"
	nats "github.com/groupchat/mq/nats"
	gonats "github.com/nats-io/go-nats"
)

func main() {
	log.SetFlags(log.Llongfile | log.Ldate)

	//Init Message Queueing
	publisher, err := nats.NewPublisher(gonats.DefaultURL)
	if err != nil {
		log.Println(err)
		return
	}

	subscriber := nats.NewSubscriber(gonats.DefaultURL)

	//Init Chat Server
	server := chat.NewServer(publisher, subscriber)

	//Init Chat Server Handler
	handler := handler.New(server, ":3000")
	handler.Run()
}
