package nsq

import (
	"log"

	"github.com/garyburd/redigo/redis"
	"github.com/groupchat/chat"
	nats "github.com/nats-io/go-nats"
)

type Subscriber struct {
	gPubSubConn *redis.PubSubConn
	gRedisConn  *redis.Conn
}

func NewSubscriber(addr string) chat.Subscriber {
	conn, err := redis.Dial("tcp", addr)
	if err != nil {
		return nil
	}

	return &Subscriber{
		gPubSubConn: &redis.PubSubConn{Conn: conn},
		gRedisConn:  &conn,
	}
}

func (s *Subscriber) Subscribe(usr *chat.User) error {
	err := s.gPubSubConn.Subscribe(usr.RoomID)
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
