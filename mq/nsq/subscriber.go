package nsq

import (
	"log"

	"github.com/groupchat/chat"
	"github.com/nsqio/go-nsq"
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
	config := nsq.NewConfig()
	config.Set("max_attempts", 200)
	consumer, err := nsq.NewConsumer(string(usr.RoomID), string(usr.UserID), config)
	if err != nil {
		log.Println(err)
		return err
	}

	consumer.ChangeMaxInFlight(5)
	consumer.AddHandler(nsq.HandlerFunc(s.handleMessage(usr)))

	if err := consumer.ConnectToNSQLookupd(s.addr); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *Subscriber) handleMessage(consumer chat.Consumer) func(m *nsq.Message) error {
	return func(msg *nsq.Message) error {
		log.Println(msg.Body)
		if err := consumer.Consume(string(msg.Body)); err != nil {
			log.Println(err)
			return err
		}
		return nil
	}
}
