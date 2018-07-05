package nsq

import (
	"log"

	"github.com/groupchat/chat"
	nats "github.com/nats-io/go-nats"
)

type Subscriber struct {
	addr string
}

func NewSubscriber(addr string) chat.Subscriber {
	return &Subscriber{
		addr: addr,
	}
}

func (s *Subscriber) Subscribe(usr *chat.User) error {
	conn, err := nats.Connect(s.addr)
	if err != nil {
		return err
	}

	_, err = conn.Subscribe(string(usr.RoomID), s.handleMessage(usr))
	if err != nil {
		return err
	}

	return nil
}

func (s *Subscriber) handleMessage(consumer chat.Consumer) func(m *nats.Msg) {
	return func(msg *nats.Msg) {
		log.Println(msg.Data)
		if err := consumer.Consume(string(msg.Data)); err != nil {
			log.Println(err)
		}
	}
}
