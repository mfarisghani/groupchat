package nsq

import (
	"log"

	"github.com/garyburd/redigo/redis"
	"github.com/groupchat/chat"
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

	go func() {
		for {
			switch v := s.gPubSubConn.Receive().(type) {
			case redis.Message:
				if err := usr.Consume(string(v.Data)); err != nil {
					log.Println(err)
					break
				}
			case redis.Subscription:
				log.Printf("subscription message: %s: %s %d\n", v.Channel, v.Kind, v.Count)
			case error:
				log.Println("error pub/sub, delivery has stopped")
				break
			}
		}
	}()

	return nil
}
