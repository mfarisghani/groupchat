package main

import (
	"log"

	"github.com/groupchat/chat/chatserver"
	nats "github.com/groupchat/mq/nats"
	gonats "github.com/nats-io/go-nats"
)

func main() {
	//Init Message Queueing
	natsPublisher, err := nats.NewPublisher(gonats.DefaultURL)
	if err != nil {
		log.Println(err)
		return
	}

	natsSubscriber := nats.NewSubscriber(gonats.DefaultURL)

	//Init Chat Server
	server := chatserver.New(":3000", natsPublisher, natsSubscriber)
	server.Run()
}
