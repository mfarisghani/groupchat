package nsq

import (
	"log"

	"github.com/nats-io/go-nats"

	"github.com/groupchat/chat"
)

type Publisher struct {
	Connection *nats.Conn
}

func NewPublisher(addr string) (chat.Publisher, error) {
	conn, err := nats.Connect(addr)
	if err != nil {
		return nil, err
	}

	publisher := &Publisher{
		Connection: conn,
	}

	return publisher, nil
}

func (p *Publisher) Publish(roomID chat.RoomID, message string) error {
	log.Println(roomID, message)
	err := p.Connection.Publish(string(roomID), []byte(message))
	if err != nil {
		return err
	}
	return nil
}
