package nsq

import (
	"log"

	"github.com/garyburd/redigo/redis"

	"github.com/groupchat/chat"
)

var (
	gRedisConn = func() (redis.Conn, error) {
		return redis.Dial("tcp", ":6379")
	}
)

type Publisher struct {
	addr string
}

func NewPublisher(addr string) (chat.Publisher, error) {
	publisher := &Publisher{
		addr: addr,
	}

	return publisher, nil
}

func (p *Publisher) Publish(roomID chat.RoomID, message string) error {
	if c, err := gRedisConn(); err != nil {
		log.Printf("error on redis conn. %s\n", err)
	} else {
		c.Do("PUBLISH", string(roomID), message)
	}
	return nil
}
