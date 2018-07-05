package nsq

import (
	"log"

	"github.com/groupchat/chat"
	nats "github.com/nats-io/go-nats"
)

type Subscriber struct {
	addr string
	conn *nats.Conn
}

func NewSubscriber(addr string) chat.Subscriber {
	conn, _ := nats.Connect(addr)

	return &Subscriber{
		addr: addr,
		conn: conn,
	}
}

func (s *Subscriber) Subscribe(usr *chat.User) error {
	_, err := s.conn.Subscribe(string(usr.RoomID), s.handleMessage(usr))
	if err != nil {
		return err
	}

	return nil
}

func (s *Subscriber) handleMessage(consumer chat.Consumer) func(m *nats.Msg) {
	return func(msg *nats.Msg) {
		if err := consumer.Consume(string(msg.Data)); err != nil {
			log.Println(err)
		}
	}
}
